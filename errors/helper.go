package errors

import (
	"net/http"
)

func withCode(code Code, err error) ServiceError {
	if err == nil {
		return nil
	}

	return &serviceError{
		code:  code,
		error: err,
	}
}

func isInvalid(code Code) bool {
	_, exists := validErrors[code]
	return !exists
}

func errCodeIs(errCode Code, codes ...Code) bool {
	for i := range codes {
		if errCode == codes[i] {
			return true
		}
	}
	return false
}

func GetHttpStatusAndCode(err error) (httpStatus int, errCodeStr string) {
	/*errCode := err.GetCode()
	errCodeStr = errCode.ErrorCode()

	errMap := errorMap()
	httpStatus, ok := errMap[errCode]
	if !ok {
		httpStatus = http.StatusInternalServerError
	}*/
	errCodeStr = err.Error()
	httpStatus = http.StatusBadRequest

	return
}

func errorMap() map[Code]int {
	return map[Code]int{
		TemporaryUnavailable: http.StatusRequestTimeout,
		Canceled:             http.StatusRequestTimeout,
		Timeout:              http.StatusRequestTimeout,
		InvalidRequest:       http.StatusBadRequest,
		NotFound:             http.StatusNotFound,
		MethodNotAllowed:     http.StatusMethodNotAllowed,
		Internal:             http.StatusInternalServerError,
	}
}
