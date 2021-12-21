package peoplez

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/seggga/back2/internal/peoplez/model"
)

// RESTAPI represents a REST API business logic server.
type RESTAPI struct {
	server http.Server
	errors chan error
}

// New returns a new instance of the REST API server.
func New(stor model.Storage, srvAddr string) (*RESTAPI, error) {
	// define routes
	router := chi.NewRouter()
	router.Get("/", homeHandler())
	router.Post("/new-user", newUserHandler(stor))
	router.Post("/new-union", newUnionHandler(stor))

	return &RESTAPI{
		server: http.Server{
			Addr:    srvAddr,
			Handler: router,
		},
		errors: make(chan error, 1),
	}, nil
}

// Start method starts the API server.
func (rapi *RESTAPI) Start() {
	go func() {
		rapi.errors <- rapi.server.ListenAndServe()
		close(rapi.errors)
	}()
}

// Stop method stops API server.
func (rapi *RESTAPI) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return rapi.server.Shutdown(ctx)
}

// Notify returns a channel to notify the caller about errors.
// If you receive an error from the channel you should stop the application.
func (rapi *RESTAPI) Notify() <-chan error {
	return rapi.errors
}
