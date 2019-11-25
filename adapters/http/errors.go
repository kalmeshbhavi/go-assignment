package http

import (
	"log"
	"net/http"

	"github.com/kalmeshbhavi/go-assignment/errors"
)

// Error for responding error
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func generateErrorResponse(code string, errMsg string) Error {
	return Error{
		Code:    code,
		Message: errMsg,
	}
}

func handleError(resp Response, r *http.Request) {
	writeErrResponse(r, resp)
}

func writeErrResponse(r *http.Request, resp Response) {
	err := resp.GetError()
	httpStatus, errCodeStr := errors.GetHttpStatusAndCode(err)
	msg := getErrorMessage(err)
	errorResponse := generateErrorResponse(errCodeStr, msg)

	if httpStatus == http.StatusInternalServerError {
		//ctxzap.Extract(r.Context()).Error("Failed to write data when unauthorized", zap.Error(err))
	}
	//ctxzap.Extract(r.Context()).Error(r.RequestURI, zap.Error(err))

	err = resp.WriteJSON(httpStatus, errorResponse)
	if err != nil {
		//logger.FromContext(r.Context()).Error("Failed to send error response", zap.Error(err))
		log.Fatalf("Failed to send error response %v", err)
	}
}

func getErrorMessage(err errors.ServiceError) string {
	switch err.GetCode() {
	case errors.NotFound:
		return "Not found"
	case errors.InvalidRequest:
		return "Invalid request"
	}
	return ""
}
