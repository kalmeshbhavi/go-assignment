package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/kalmeshbhavi/go-assignment/domain"
)

func (adapter *HTTPAdapter) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]

		knight, err := adapter.engine.GetKnight(idStr)
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, knight)
	}
}

func (adapter *HTTPAdapter) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		knights, err := adapter.engine.ListKnights()
		if err != nil {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, knights)
	}
}

func (adapter *HTTPAdapter) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		var knight domain.Knight
		err = json.Unmarshal(body, &knight)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		err = validateRequest(knight)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		id, err := adapter.engine.CreateKnight(&knight)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		res := struct {
			ID string `json:"id"`
		}{ID: strconv.Itoa(int(id))}
		respondWithJSON(w, http.StatusCreated, res)
	}
}

func validateRequest(knight domain.Knight) error {
	if knight.WeaponPower == 0 {
		return errors.New("Invalid request")
	}

	if knight.Strength == 0 {
		return errors.New("Invalid request")
	}
	if knight.Name == "" {
		return errors.New("Invalid request")
	}
	return nil
}
