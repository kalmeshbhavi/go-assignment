package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kalmeshbhavi/go-assignment/domain"
	"github.com/kalmeshbhavi/go-assignment/errors"
)

func (s *HTTPAdapter) get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := ToResponse(w)

		vars := mux.Vars(r)
		idStr := vars["id"]

		path, serr := s.engine.GetKnight(idStr)
		if serr != nil {
			resp.SetError(errors.New(errors.NotFound, serr))
			return
		}

		resp.SetResponse(path)
	})
}

func (s *HTTPAdapter) getAll() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := ToResponse(w)
		knights := s.engine.ListKnights()
		resp.SetResponse(knights)
	})
}

func (s *HTTPAdapter) create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := ToResponse(w)

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		var sd domain.Knight
		err = json.Unmarshal(body, &sd)
		if err != nil {
			resp.SetError(errors.New(errors.NotFound, err))
			return
		}

	})
}
