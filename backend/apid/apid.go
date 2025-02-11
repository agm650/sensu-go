package apid

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	corev3 "github.com/sensu/core/v3"
	"github.com/sensu/sensu-go/backend/api"
	"github.com/sensu/sensu-go/backend/apid/actions"
	"github.com/sensu/sensu-go/backend/apid/graphql"
	"github.com/sensu/sensu-go/backend/apid/handlers"
	"github.com/sensu/sensu-go/backend/apid/middlewares"
	"github.com/sensu/sensu-go/backend/apid/routers"
	"github.com/sensu/sensu-go/backend/authentication"
	"github.com/sensu/sensu-go/backend/authorization/rbac"
	"github.com/sensu/sensu-go/backend/messaging"
	storev2 "github.com/sensu/sensu-go/backend/store/v2"
	"github.com/sensu/sensu-go/types"
)

// APId is the backend HTTP API.
type APId struct {
	Authenticator              *authentication.Authenticator
	HTTPServer                 *http.Server
	CoreSubrouter              *mux.Router
	CoreV3Subrouter            *mux.Router
	EntityLimitedCoreSubrouter *mux.Router
	GraphQLSubrouter           *mux.Router
	RequestLimit               int64

	stopping    chan struct{}
	running     *atomic.Value
	wg          *sync.WaitGroup
	errChan     chan error
	bus         messaging.MessageBus
	store       storev2.Interface
	queueGetter types.QueueGetter
	tls         *types.TLSOptions
}

// Option is a functional option.
type Option func(*APId) error

// Config configures APId.
type Config struct {
	ListenAddress  string
	RequestLimit   int64
	WriteTimeout   time.Duration
	URL            string
	Bus            messaging.MessageBus
	Store          storev2.Interface
	TLS            *types.TLSOptions
	Authenticator  *authentication.Authenticator
	ClusterVersion string
	GraphQLService *graphql.Service
}

// New creates a new APId.
func New(c Config, opts ...Option) (*APId, error) {
	a := &APId{
		store:         c.Store,
		tls:           c.TLS,
		bus:           c.Bus,
		stopping:      make(chan struct{}, 1),
		running:       &atomic.Value{},
		wg:            &sync.WaitGroup{},
		errChan:       make(chan error, 1),
		Authenticator: c.Authenticator,
		RequestLimit:  c.RequestLimit,
	}

	// prepare TLS config
	var tlsServerConfig *tls.Config
	var err error
	if c.TLS != nil {
		tlsServerConfig, err = c.TLS.ToServerTLSConfig()
		if err != nil {
			return nil, err
		}
	}

	router := NewRouter()
	_ = PublicSubrouter(router, c)
	a.GraphQLSubrouter = GraphQLSubrouter(router, c)
	_ = AuthenticationSubrouter(router, c)
	a.CoreSubrouter = CoreSubrouter(router, c)
	a.CoreV3Subrouter = CoreV3Subrouter(router, c)
	a.EntityLimitedCoreSubrouter = EntityLimitedCoreSubrouter(router, c)

	a.HTTPServer = &http.Server{
		Addr:         c.ListenAddress,
		Handler:      router,
		WriteTimeout: c.WriteTimeout,
		ReadTimeout:  15 * time.Second,
		TLSConfig:    tlsServerConfig,
	}

	for _, o := range opts {
		if err := o(a); err != nil {
			return nil, err
		}
	}

	return a, nil
}

// NewRouter creates a new mux router that implements the http.Handler interface
// and serves all requests
func NewRouter() *mux.Router {
	router := mux.NewRouter().UseEncodedPath()

	// Register a default handler when no routes match
	router.NotFoundHandler = middlewares.SimpleLogger{}.Then(http.HandlerFunc(notFoundHandler))

	return router
}

// AuthenticationSubrouter initializes a subrouter that handles all
// authentication requests
func AuthenticationSubrouter(router *mux.Router, cfg Config) *mux.Router {
	subrouter := NewSubrouter(
		router.NewRoute(),
		middlewares.SimpleLogger{},
		middlewares.RefreshToken{},
		middlewares.LimitRequest{Limit: cfg.RequestLimit},
	)

	mountRouters(subrouter,
		routers.NewAuthenticationRouter(api.NewAuthenticationClient(cfg.Authenticator)),
	)

	return subrouter
}

// CoreSubrouter initializes a subrouter that handles all requests coming to
// /api/core/v2
func CoreSubrouter(router *mux.Router, cfg Config) *mux.Router {
	subrouter := NewSubrouter(
		router.PathPrefix("/api/{group:core}/{version:v2}/"),
		middlewares.Namespace{},
		middlewares.Authentication{Store: cfg.Store},
		middlewares.SimpleLogger{},
		middlewares.AuthorizationAttributes{},
		middlewares.Authorization{Authorizer: &rbac.Authorizer{Store: cfg.Store}},
		middlewares.LimitRequest{Limit: cfg.RequestLimit},
		middlewares.Pagination{},
	)
	mountRouters(
		subrouter,
		routers.NewAssetRouter(cfg.Store),
		routers.NewAPIKeysRouter(cfg.Store),
		routers.NewChecksRouter(cfg.Store, nil),
		routers.NewClusterRolesRouter(cfg.Store),
		routers.NewClusterRoleBindingsRouter(cfg.Store),
		routers.NewEventFiltersRouter(cfg.Store),
		routers.NewHandlersRouter(cfg.Store),
		routers.NewHooksRouter(cfg.Store),
		routers.NewMutatorsRouter(cfg.Store),
		routers.NewPipelinesRouter(cfg.Store),
		routers.NewRolesRouter(cfg.Store),
		routers.NewRoleBindingsRouter(cfg.Store),
		routers.NewSilencedRouter(cfg.Store),
		routers.NewTessenRouter(actions.NewTessenController(cfg.Store, cfg.Bus)),
		routers.NewUsersRouter(cfg.Store),
	)

	return subrouter
}

func CoreV3Subrouter(router *mux.Router, cfg Config) *mux.Router {
	subrouter := NewSubrouter(
		router.PathPrefix("/api/{group:core}/{version:v3}/"),
		middlewares.Namespace{},
		middlewares.Authentication{Store: cfg.Store},
		middlewares.SimpleLogger{},
		middlewares.AuthorizationAttributes{},
		middlewares.Authorization{Authorizer: &rbac.Authorizer{Store: cfg.Store}},
		middlewares.LimitRequest{Limit: cfg.RequestLimit},
		middlewares.Pagination{},
	)
	mountRouters(
		subrouter,
		routers.NewNamespacesRouter(api.NewNamespaceClient(cfg.Store, &rbac.Authorizer{Store: cfg.Store}), handlers.NewHandlers[*corev3.Namespace](cfg.Store)),
	)
	return subrouter
}

// EntityLimitedCoreSubrouter initializes a subrouter that handles all requests
// coming to /api/core/v2 that must be gated by entity limits.
func EntityLimitedCoreSubrouter(router *mux.Router, cfg Config) *mux.Router {
	subrouter := NewSubrouter(
		router.PathPrefix("/api/{group:core}/{version:v2}/"),
		middlewares.Namespace{},
		middlewares.Authentication{Store: cfg.Store},
		middlewares.SimpleLogger{},
		middlewares.AuthorizationAttributes{},
		middlewares.Authorization{Authorizer: &rbac.Authorizer{Store: cfg.Store}},
		middlewares.LimitRequest{Limit: cfg.RequestLimit},
		middlewares.Pagination{},
	)
	mountRouters(
		subrouter,
		routers.NewEntitiesRouter(cfg.Store),
		routers.NewEventsRouter(cfg.Store, cfg.Bus),
	)

	return subrouter
}

// GraphQLSubrouter initializes a subrouter that handles all requests for
// GraphQL
func GraphQLSubrouter(router *mux.Router, cfg Config) *mux.Router {
	subrouter := NewSubrouter(
		router.NewRoute(),
		middlewares.LimitRequest{Limit: cfg.RequestLimit},
		// We permit requests that do not include an access token or API key,
		// this allows unauthenticated clients to run introspecton queries or
		// query resources that do not require authorization, such as health
		// and version info.
		//
		// https://github.com/graphql/graphiql
		// https://graphql.org/learn/introspection/
		middlewares.Authentication{IgnoreUnauthorized: true, Store: cfg.Store},
		middlewares.SimpleLogger{},
	)

	// The write timeout hangs up the request making it more difficult for
	// clients to determine what occurred. As such give the service as much time
	// as possible to produce results.
	timeout := cfg.WriteTimeout - (50 * time.Millisecond)
	if timeout < 0 {
		timeout = 0
	}

	mountRouters(
		subrouter,
		&routers.GraphQLRouter{
			Service: cfg.GraphQLService,
			Timeout: timeout,
		},
	)

	return subrouter
}

// PublicSubrouter initializes a subrouter that handles all requests to public
// endpoints
func PublicSubrouter(router *mux.Router, cfg Config) *mux.Router {
	subrouter := NewSubrouter(
		router.NewRoute(),
		middlewares.SimpleLogger{},
		middlewares.LimitRequest{Limit: cfg.RequestLimit},
	)

	mountRouters(subrouter,
		routers.NewVersionRouter(actions.NewVersionController(cfg.ClusterVersion)),
		routers.NewTessenMetricRouter(actions.NewTessenMetricController(cfg.Bus)),
	)

	subrouter.Handle("/metrics", promhttp.Handler())

	return subrouter
}

func notFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	resp := map[string]interface{}{
		"message": "not found", "code": actions.NotFound,
	}
	_ = json.NewEncoder(w).Encode(resp)
}

// Start APId.
func (a *APId) Start() error {
	logger.Warn("starting apid on address: ", a.HTTPServer.Addr)
	ln, err := net.Listen("tcp", a.HTTPServer.Addr)
	if err != nil {
		return fmt.Errorf("failed to start apid: %s", err)
	}

	a.wg.Add(1)

	go func() {
		defer a.wg.Done()
		var err error
		if a.tls != nil {
			// TLS configuration comes from ToServerTLSConfig
			err = a.HTTPServer.ServeTLS(ln, "", "")
		} else {
			err = a.HTTPServer.Serve(ln)
		}
		if err != nil && err != http.ErrServerClosed {
			a.errChan <- fmt.Errorf("failure while serving api: %s", err)
		}
	}()

	return nil
}

// Stop httpApi.
func (a *APId) Stop() error {
	if err := a.HTTPServer.Shutdown(context.TODO()); err != nil {
		// failure/timeout shutting down the server gracefully
		logger.Error("failed to shutdown http server gracefully - forcing shutdown")
		if closeErr := a.HTTPServer.Close(); closeErr != nil {
			logger.Error("failed to shutdown http server forcefully")
		}
	}

	a.running.Store(false)
	close(a.stopping)
	a.wg.Wait()
	close(a.errChan)

	return nil
}

// Err returns a channel to listen for terminal errors on.
func (a *APId) Err() <-chan error {
	return a.errChan
}

// Name returns the daemon name
func (a *APId) Name() string {
	return "apid"
}

func mountRouters(parent *mux.Router, subRouters ...routers.Router) {
	for _, subRouter := range subRouters {
		subRouter.Mount(parent)
	}
}
