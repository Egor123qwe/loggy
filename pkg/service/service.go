package service

import (
	"github.com/Egor123qwe/loggy/pkg/model"
	"github.com/Egor123qwe/loggy/pkg/service/logger"
	"github.com/op/go-logging"
)

var logsrv logger.Service

func Log(traceID string) *logging.Logger {
	if logsrv == nil {
		panic(model.NotInitializedErr)
	}

	return logsrv.New(traceID)
}

func Close() error {
	return nil
}
