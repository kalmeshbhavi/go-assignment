package http

import (
	"encoding/json"
	"net/http"

	"github.com/kalmeshbhavi/go-assignment/errors"
)

const (
	jsonContentType = "application/json; charset=UTF-8"
)

// Response is utils for http.ResponseWriter
type Response interface {
	http.ResponseWriter

	GetStatus() int
	GetError() errors.ServiceError
	GetResponse() interface{}

	SetError(err errors.ServiceError)
	SetResponse(i interface{})

	WriteJSON(code int, i interface{}) errors.ServiceError
}

type response struct {
	http.ResponseWriter

	status    int
	committed bool

	resp interface{}
	err  errors.ServiceError
}

func ToResponse(w http.ResponseWriter) Response {
	if resp, ok := w.(Response); ok {
		return resp
	}

	return NewResponse(w)
}

func NewResponse(w http.ResponseWriter) Response {
	return &response{
		w,
		0,
		false,
		nil,
		nil,
	}
}

func (r *response) GetStatus() int {
	return r.status
}

func (r *response) GetError() errors.ServiceError {
	return r.err
}

func (r *response) GetResponse() interface{} {
	return r.resp
}

func (r *response) SetResponse(resp interface{}) {
	if r.resp != nil {
		panic("response is not nil, cannot be overwritten")
	}
	r.resp = resp
}

func (r *response) SetError(err errors.ServiceError) {
	if r.err != nil {
		panic("error is not nil, cannot be overwritten")
	}
	r.err = err
}

func (r *response) WriteJSON(code int, i interface{}) errors.ServiceError {
	r.setContentType(code, jsonContentType)
	if err := json.NewEncoder(r).Encode(i); err != nil {
		return errors.NewFromError(err, errors.Internal)
	}
	return nil
}

func (r *response) setContentType(code int, contentType string) {
	r.Header().Set("Content-Type", contentType)
	r.WriteHeader(code)
}

func (r *response) WriteHeader(code int) {
	if r.committed {
		panic("already wrote to response header")
	}
	r.ResponseWriter.WriteHeader(code)
	r.status = code
	r.committed = true
}

func (r *response) Write(b []byte) (n int, err error) {
	if !r.committed {
		r.WriteHeader(http.StatusOK)
	}
	return r.ResponseWriter.Write(b)
}
