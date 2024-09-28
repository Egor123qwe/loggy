package service

import (
	"github.com/Egor123qwe/loggy/pkg/model"
	"github.com/Egor123qwe/loggy/pkg/producer"
	"github.com/Egor123qwe/loggy/pkg/service/logger"
	"github.com/op/go-logging"
)

var logsrv logger.Service

func Init(opts Options) error {
	srvOpts := logger.Options{
		Level:  opts.Level,
		Module: opts.Module,

		ToStderr: opts.ToStderr,
	}

	if opts.File != nil {
		srvOpts.File = &logger.File{
			Name:      opts.File.Name,
			MaxSizeMb: opts.File.MaxSizeMb,
			MaxFiles:  opts.File.MaxFiles,
		}
	}

	if opts.Server != nil {
		var credentials *producer.Credentials

		if opts.Server.Credentials != nil {
			credentials = &producer.Credentials{
				Username: opts.Server.Credentials.Username,
				Password: opts.Server.Credentials.Password,
			}
		}

		producer, err := producer.New(credentials, opts.Server.URLs)
		if err != nil {
			return err
		}

		srvOpts.Sender = newSender(producer)
	}

	logsrv = logger.New(srvOpts)

	return nil
}

func Log(traceID string) *logging.Logger {
	if logsrv == nil {
		panic(model.NotInitializedErr)
	}

	return logsrv.Get(traceID)
}
