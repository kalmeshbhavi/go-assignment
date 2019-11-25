package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	adapter "github.com/kalmeshbhavi/go-assignment/adapters/middleware"
	"github.com/kalmeshbhavi/go-assignment/engine"
)

type Server struct {
	mux    *mux.Router
	server *http.Server
	engine engine.Engine
	logger *log.Logger
}

func NewServer(engine engine.Engine, logger *log.Logger) Server {
	server := Server{
		engine: engine,
		logger: logger,
	}

	mux := mux.NewRouter()

	server.registerHandlers(mux, getServerAdapters(nil))
	server.server = &http.Server{
		Handler: mux,
	}
	return server
}

func getServerAdapters(logger *log.Logger) []adapter.ServerAdapter {
	return []adapter.ServerAdapter{
		responseServerAdapter(),
	}
}

func (s *Server) registerHandlers(mux *mux.Router, adapters []adapter.ServerAdapter) {
	//handle(mux, "/knight", adapter.ServerApply(s.getAll(), adapters...))
	//handle(mux, "/knight/{id}", adapter.ServerApply(s.get(), adapters...))
	//handle(mux, "/knight", adapter.ServerApply(s.create(), adapters...))
}

func handle(mux *mux.Router, pattern string, handler http.Handler) {
	mux.Handle(pattern, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		resp := NewResponse(w)
		handler.ServeHTTP(resp, req)
	}))
}
