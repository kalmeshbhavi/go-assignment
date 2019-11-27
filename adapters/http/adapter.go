package http

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kalmeshbhavi/go-assignment/engine"
)

type HTTPAdapter struct {
	engine  engine.Engine
	server  http.Server
	Router  *mux.Router
	context context.Context
}

func (adapter *HTTPAdapter) Start() {
	adapter.server = http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: adapter.Router,
	}

	// ErrServerClosed is returned by the Server's Serve
	// after a call to Shutdown or Close, we can ignore it.
	go func() {
		if err := adapter.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()
}

func (adapter *HTTPAdapter) Stop() {
	log.Print("stopping server gracefully")
	err := adapter.server.Shutdown(adapter.context)
	if err != nil {
		log.Fatal("error during graceful shutdown", err)
	}
}

func NewHTTPAdapter(e engine.Engine) *HTTPAdapter {
	// todo: init your http server and routes

	adapter := &HTTPAdapter{engine: e}
	adapter.Router = mux.NewRouter()
	adapter.route(getServerAdapters())
	return adapter
}

func (adapter *HTTPAdapter) route(adapters []ServerAdapter) {
	adapter.Router.Handle("/knight", adapter.getAll()).Methods("GET")
	adapter.Router.Handle("/knight", adapter.create()).Methods("POST")
	adapter.Router.Handle("/knight/{id}", adapter.get()).Methods("GET")
}

func getServerAdapters() []ServerAdapter {
	return []ServerAdapter{
		responseServerAdapter(),
	}
}

type ServerAdapter func(http.Handler) http.Handler

func (adapter *HTTPAdapter) ServerApply(h http.Handler, adapters ...ServerAdapter) http.Handler {

	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
