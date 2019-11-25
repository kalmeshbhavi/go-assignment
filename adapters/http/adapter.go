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
	Router  *mux.Router
	context context.Context
}

func (adapter *HTTPAdapter) Start() {
	// todo: start to listen
	log.Fatal(http.ListenAndServe(":8000", adapter.Router))
}

func (adapter *HTTPAdapter) Stop() {
	// todo: shutdown server
	//adapter.Router.Shutdown(adapter.context)
}

func NewHTTPAdapter(e engine.Engine) *HTTPAdapter {
	// todo: init your http server and routes

	adapter := &HTTPAdapter{engine: e}
	router := mux.NewRouter()

	router.HandleFunc("/knight", adapter.get()).Methods("GET")
	router.HandleFunc("/knight", nil).Methods("POST")
	router.HandleFunc("/knight/{id}", nil).Methods("GET")

	adapter.Router = router

	return adapter
}
