package keepalived

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/sensu/sensu-go/backend/messaging"
	"github.com/sensu/sensu-go/backend/monitor"
	"github.com/sensu/sensu-go/backend/store"
	"github.com/sensu/sensu-go/types"
)

const (
	// DefaultHandlerCount is the default number of goroutines dedicated to
	// handling keepalive events.
	DefaultHandlerCount = 10

	// DefaultKeepaliveTimeout is the amount of time we consider a Keepalive
	// valid for.
	DefaultKeepaliveTimeout = 120 // seconds

	// KeepaliveCheckName is the name of the check that is created when a
	// keepalive timeout occurs.
	KeepaliveCheckName = "keepalive"

	// KeepaliveHandlerName is the name of the handler that is executed when
	// a keepalive timeout occurs.
	KeepaliveHandlerName = "keepalive"

	// RegistrationCheckName is the name of the check that is created when an
	// entity sends a keepalive and the entity does not yet exist in the store.
	RegistrationCheckName = "registration"

	// RegistrationHandlerName is the name of the handler that is executed when
	// a registration event is passed to pipelined.
	RegistrationHandlerName = "registration"
)

// MonitorFactoryFunc takes an entity and returns a Monitor. Keepalived can
// take a MonitorFactoryFunc that stubs/mocks a Deregisterer and/or an
// EventCreator to make it easier to test.
type MonitorFactoryFunc func(e *types.Entity) *monitor.Monitor

// Keepalived is responsible for monitoring keepalive events and recording
// keepalives for entities.
type Keepalived struct {
	MessageBus            messaging.MessageBus
	HandlerCount          int
	Store                 store.Store
	DeregistrationHandler string
	MonitorFactory        MonitorFactoryFunc

	mu            *sync.Mutex
	monitors      map[string]*monitor.Monitor
	wg            *sync.WaitGroup
	keepaliveChan chan interface{}
	errChan       chan error
}

// Start starts the daemon, returning an error if preconditions for startup
// fail.
func (k *Keepalived) Start() error {
	if k.MessageBus == nil {
		return errors.New("no message bus found")
	}

	if k.Store == nil {
		return errors.New("no keepalive store found")
	}

	if k.MonitorFactory == nil {
		k.MonitorFactory = func(e *types.Entity) *monitor.Monitor {
			return &monitor.Monitor{
				Entity:         e,
				Timeout:        time.Duration(e.KeepaliveTimeout),
				FailureHandler: k,
				UpdateHandler:  k,
			}
		}
	}

	k.keepaliveChan = make(chan interface{}, 10)
	err := k.MessageBus.Subscribe(messaging.TopicKeepalive, "keepalived", k.keepaliveChan)
	if err != nil {
		return err
	}

	if k.HandlerCount == 0 {
		k.HandlerCount = DefaultHandlerCount
	}

	k.mu = &sync.Mutex{}
	k.monitors = map[string]*monitor.Monitor{}

	if err := k.initFromStore(); err != nil {
		return err
	}

	k.startWorkers()

	k.startMonitorSweeper()

	k.errChan = make(chan error, 1)
	return nil
}

// Stop stops the daemon, returning an error if one was encountered during
// shutdown.
func (k *Keepalived) Stop() error {
	close(k.keepaliveChan)
	k.wg.Wait()
	for _, monitor := range k.monitors {
		go monitor.Stop()
	}
	err := k.MessageBus.Unsubscribe(messaging.TopicKeepalive, "keepalived")
	close(k.errChan)
	return err
}

// Status returns nil if the Daemon is healthy, otherwise it returns an error.
func (k *Keepalived) Status() error {
	return nil
}

// Err returns a channel that the caller can use to listen for terminal errors
// indicating a premature shutdown of the Daemon.
func (k *Keepalived) Err() <-chan error {
	return k.errChan
}

func (k *Keepalived) initFromStore() error {
	// For which clients were we previously alerting?
	keepalives, err := k.Store.GetFailingKeepalives(context.TODO())
	if err != nil {
		return err
	}

	for _, keepalive := range keepalives {
		entityCtx := context.WithValue(context.TODO(), types.OrganizationKey, keepalive.Organization)
		entityCtx = context.WithValue(entityCtx, types.EnvironmentKey, keepalive.Environment)
		event, err := k.Store.GetEventByEntityCheck(entityCtx, keepalive.EntityID, "keepalive")
		if err != nil {
			return err
		}

		// if there's no event, the entity was deregistered/deleted.
		if event == nil {
			continue
		}

		// if another backend picked it up, it will be passing.
		if event.Check.Status == 0 {
			continue
		}

		// Recreate the monitor and reset its timer to alert when it's going to
		// timeout.
		monitor := k.MonitorFactory(event.Entity)
		monitor.Reset(keepalive.Time)
		k.monitors[keepalive.EntityID] = monitor
	}

	return nil
}

func (k *Keepalived) startWorkers() {
	k.wg = &sync.WaitGroup{}
	k.wg.Add(k.HandlerCount)

	for i := 0; i < k.HandlerCount; i++ {
		go k.processKeepalives()
	}
}

func (k *Keepalived) processKeepalives() {
	defer k.wg.Done()

	var (
		monitor *monitor.Monitor
		event   *types.Event
		ok      bool
	)

	for msg := range k.keepaliveChan {
		event, ok = msg.(*types.Event)
		if !ok {
			logger.Error("keepalived received non-Event on keepalive channel")
			continue
		}

		entity := event.Entity
		if entity == nil {
			logger.Error("received keepalive with nil entity")
			continue
		}

		if err := entity.Validate(); err != nil {
			logger.WithError(err).Error("invalid keepalive event")
			continue
		}

		if err := k.handleEntityRegistration(entity); err != nil {
			logger.WithError(err).Error("error handling entity registration")
		}

		k.mu.Lock()
		monitor, ok = k.monitors[entity.ID]
		// create if it doesn't exist
		if !ok || monitor.IsStopped() {
			monitor = k.MonitorFactory(entity)
			monitor.Start()
			k.monitors[entity.ID] = monitor
		}
		k.mu.Unlock()

		if err := monitor.HandleUpdate(event); err != nil {
			logger.WithError(err).Error("error monitoring entity")
		}
	}
}

func (k *Keepalived) handleEntityRegistration(entity *types.Entity) error {
	if entity.Class != types.EntityAgentClass {
		return nil
	}

	ctx := types.SetContextFromResource(context.Background(), entity)
	fetchedEntity, err := k.Store.GetEntityByID(ctx, entity.ID)

	if err != nil {
		return err
	}

	if fetchedEntity == nil {
		event := createRegistrationEvent(entity)
		err = k.MessageBus.Publish(messaging.TopicEvent, event)
	}

	return err
}

// startMonitorSweeper spins off into oblivion if Keepalived is stopped until
// the monitors map is empty, and then the goroutine stops.
func (k *Keepalived) startMonitorSweeper() {
	go func() {
		timer := time.NewTimer(10 * time.Minute)
		for {
			<-timer.C
			for key, monitor := range k.monitors {
				if monitor.IsStopped() {
					k.mu.Lock()
					delete(k.monitors, key)
					k.mu.Unlock()
				}
			}
		}
	}()
}

func createKeepaliveEvent(entity *types.Entity) *types.Event {
	keepaliveCheck := &types.Check{
		Config: &types.CheckConfig{
			Name:         KeepaliveCheckName,
			Interval:     entity.KeepaliveTimeout,
			Handlers:     []string{KeepaliveHandlerName},
			Environment:  entity.Environment,
			Organization: entity.Organization,
		},
		Status: 1,
	}
	keepaliveEvent := &types.Event{
		Timestamp: time.Now().Unix(),
		Entity:    entity,
		Check:     keepaliveCheck,
	}

	return keepaliveEvent
}

func createRegistrationEvent(entity *types.Entity) *types.Event {
	registrationCheck := &types.Check{
		Config: &types.CheckConfig{
			Name:         RegistrationCheckName,
			Interval:     entity.KeepaliveTimeout,
			Handlers:     []string{RegistrationHandlerName},
			Environment:  entity.Environment,
			Organization: entity.Organization,
		},
		Status: 1,
	}
	registrationEvent := &types.Event{
		Timestamp: time.Now().Unix(),
		Entity:    entity,
		Check:     registrationCheck,
	}

	return registrationEvent
}

// HandleUpdate sets the entity's last seen time and publishes an OK check event
// to the message bus.
func (k *Keepalived) HandleUpdate(e *types.Event) error {
	entity := e.Entity

	ctx := types.SetContextFromResource(context.Background(), entity)
	if err := k.Store.DeleteFailingKeepalive(ctx, e.Entity); err != nil {
		return err
	}

	entity.LastSeen = e.Timestamp
	if err := k.Store.UpdateEntity(ctx, entity); err != nil {
		logger.WithError(err).Error("error updating entity in store")
		return err
	}
	event := createKeepaliveEvent(entity)
	event.Check.Status = 0
	return k.MessageBus.Publish(messaging.TopicEventRaw, event)
}

// HandleFailure checks if the entity should be deregistered, and emits a
// keepalive event if the entity is still valid.
func (k *Keepalived) HandleFailure(e *types.Entity) error {
	ctx := types.SetContextFromResource(context.Background(), e)

	deregisterer := &Deregistration{
		Store:      k.Store,
		MessageBus: k.MessageBus,
	}
	// if the entity is supposed to be deregistered, do so.
	if e.Deregister {
		return deregisterer.Deregister(e)
	}

	// this is a real keepalive event, emit it.
	event := createKeepaliveEvent(e)
	event.Check.Status = 1
	k.MessageBus.Publish(messaging.TopicEventRaw, event)

	timeout := time.Now().Unix() + int64(e.KeepaliveTimeout)
	return k.Store.UpdateFailingKeepalive(ctx, e, timeout)
}
