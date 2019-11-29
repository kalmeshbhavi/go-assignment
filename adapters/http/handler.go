package http

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/kalmeshbhavi/go-assignment/domain"
	"github.com/kalmeshbhavi/go-assignment/errors"
)

func (adapter *HTTPAdapter) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]

		knight, err := adapter.engine.GetKnight(idStr)
		if err != nil {
			handleError1(w, err)
			//respondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, knight)
	}
}

func (adapter *HTTPAdapter) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		knights, err := adapter.engine.ListKnights()
		log.Print("knights : ", knights)
		if err != nil {
			log.Fatalf("%v", err)
			handleError1(w, err)
			//respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, knights)
	}
}

func (adapter *HTTPAdapter) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op errors.Op = "handle.create"
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			handleError1(w, errors.E(op, errors.KindInvalidRequest, err, "invalid request"))
			//respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		var knight domain.Knight
		err = json.Unmarshal(body, &knight)
		if err != nil {
			handleError1(w, errors.E(op, errors.KindInvalidRequest, err, "invalid request"))
			//respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		err = validateRequest(knight)
		if err != nil {
			handleError1(w, errors.E(op, errors.KindOf(err), err, err.Error()))
			//respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		id, err := adapter.engine.CreateKnight(&knight)
		if err != nil {
			handleError1(w, errors.E(op, errors.KindOf(err), err))
			//respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		res := struct {
			ID string `json:"id"`
		}{ID: strconv.Itoa(int(id))}
		respondWithJSON(w, http.StatusCreated, res)
	}
}

func handleError1(w http.ResponseWriter, err error) {
	serr, ok := err.(*errors.Error)

	if !ok {
		respondWithJSON(w, http.StatusInternalServerError, Error{
			Code:    strconv.Itoa(http.StatusInternalServerError),
			Message: "Something wrong! please contact admin",
		})
		log.Printf("%v", err)
		return
	}

	code := statusCodeByErrorKind(errors.KindOf(err))
	logErrors(serr)
	respondWithJSON(w, code, map[string]string{"code": strconv.Itoa(code), "message": err.Error()})
}

func logErrors(err *errors.Error) {
	ops := errors.Ops(err)

	log.Println(ops)

	s := getContext(err)
	contexts := strings.Join(s, " : ")
	log.Println(contexts)
}

func getContext(err *errors.Error) []string {
	res := []string{err.Context}

	subErr, ok := err.Err.(*errors.Error)
	if !ok {
		return res
	}
	res = append(res, getContext(subErr)...)
	return res

}

func statusCodeByErrorKind(kind errors.Kind) int {
	switch kind {
	case errors.KindNotFound:
		return http.StatusNotFound
	case errors.KindInvalidRequest:
		return http.StatusBadRequest
	case errors.KindUnexpected:
		return http.StatusInternalServerError
	case errors.KindInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func validateRequest(knight domain.Knight) error {
	const op errors.Op = "handler.ValidateRequest"
	if knight.WeaponPower == 0 {
		return errors.E(op, errors.KindInvalidRequest, "invalid weapon_power")
	}

	if knight.Strength == 0 {
		return errors.E(op, errors.KindInvalidRequest, "invalid strength")
	}
	if knight.Name == "" {
		return errors.E(op, errors.KindInvalidRequest, "invalid name")
	}
	return nil
}
