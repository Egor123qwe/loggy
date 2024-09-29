package service

import (
	"github.com/Egor123qwe/loggy/pkg/model"
	"github.com/Egor123qwe/loggy/pkg/service/logger"
	"github.com/op/go-logging"
)

// library data
var (
	logsrv logger.Service
)

func Log(traceID string) *logging.Logger {
	if logsrv == nil {
		panic(model.NotInitializedErr)
	}

	return logsrv.New(traceID)
}

func Close() error {
	if logsrv == nil {
		return model.NotInitializedErr
	}

	err := logsrv.Close()
	logsrv = nil

	return err
}
