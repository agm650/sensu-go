package actions

import (
	corev2 "github.com/sensu/core/v2"
	"github.com/sensu/sensu-go/backend/messaging"
	"github.com/sensu/sensu-go/backend/store"
	storev2 "github.com/sensu/sensu-go/backend/store/v2"
	"golang.org/x/net/context"
)

// TessenController exposes actions which a viewer can perform
type TessenController struct {
	store storev2.Interface
	bus   messaging.MessageBus
}

// NewTessenController returns a new TessenController
func NewTessenController(store storev2.Interface, bus messaging.MessageBus) TessenController {
	return TessenController{
		store: store,
		bus:   bus,
	}
}

// CreateOrUpdate creates or updates the tessen configuration
func (c TessenController) CreateOrUpdate(ctx context.Context, config *corev2.TessenConfig) error {
	tstore := storev2.Of[*corev2.TessenConfig](c.store)
	if err := tstore.CreateOrUpdate(ctx, config); err != nil {
		switch err := err.(type) {
		case *store.ErrNotValid:
			return NewErrorf(InvalidArgument)
		default:
			return NewError(InternalErr, err)
		}
	}

	// Publish to Tessend
	if err := c.bus.Publish(messaging.TopicTessen, config); err != nil {
		return NewError(InternalErr, err)
	}

	return nil
}

// Get gets the tessen configuration
func (c TessenController) Get(ctx context.Context) (*corev2.TessenConfig, error) {
	tstore := storev2.Of[*corev2.TessenConfig](c.store)
	// tessen resource does not have a name or namespace
	config, err := tstore.Get(ctx, storev2.ID{})
	if err != nil {
		switch err := err.(type) {
		case *store.ErrNotFound:
			return nil, NewErrorf(NotFound)
		default:
			return nil, NewError(InternalErr, err)
		}
	}

	return config, nil
}

// TessenMetricController exposes actions which a viewer can perform
type TessenMetricController struct {
	bus messaging.MessageBus
}

// NewTessenMetricController returns a new TessenMetricController
func NewTessenMetricController(bus messaging.MessageBus) TessenMetricController {
	return TessenMetricController{
		bus: bus,
	}
}

// Publish publishes the metrics to the message bus used by TessenD
func (c TessenMetricController) Publish(ctx context.Context, metrics []corev2.MetricPoint) error {
	if err := c.bus.Publish(messaging.TopicTessenMetric, metrics); err != nil {
		return NewError(InternalErr, err)
	}

	return nil
}
