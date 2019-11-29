package logger

import (
	"go.uber.org/zap"

	"github.com/kalmeshbhavi/go-assignment/errors"
)

func SystemErr(err error) {
	_, ok := err.(*errors.Error)
	if !ok {
		zap.Error(err)
		return
	}

}
