package service

import (
	"github.com/Egor123qwe/loggy/pkg/producer"
	"github.com/Egor123qwe/loggy/pkg/service/api"
	"github.com/Egor123qwe/loggy/pkg/service/logger"
)

func Init(opts Options) error {
	srvOpts := logger.Options{
		Level: opts.Level,
	}

	// set files
	for _, f := range opts.File {
		file := logger.File{
			Name:      f.Name,
			MaxSizeMb: f.MaxSizeMb,
			MaxFiles:  f.MaxFiles,
		}

		srvOpts.File = append(srvOpts.File, file)
	}

	for _, s := range opts.Server {
		sender, err := initServer(s, opts.Module)
		if err != nil {
			return err
		}
		
		srvOpts.Sender = append(srvOpts.Sender, sender)
	}

	srvOpts.ToStderr = opts.ToStderr

	logsrv = logger.New(srvOpts)

	return nil
}

func initServer(s Server, module string) (logger.Sender, error) {
	var (
		producerCredentials *producer.Credentials
		apiCredentials      api.Credentials
	)

	if s.Credentials != nil {
		producerCredentials = &producer.Credentials{
			Username: s.Credentials.Username,
			Password: s.Credentials.Password,
		}

		apiCredentials = api.Credentials{
			Username: s.Credentials.Username,
			Password: s.Credentials.Password,
		}
	}

	api, err := api.New(s.URL, apiCredentials)
	if err != nil {
		return nil, err
	}

	resp, err := api.Init(module)
	if err != nil {
		return nil, err
	}

	// TODO
	producerCredentials = nil

	producer, err := producer.New(producerCredentials, s.URLs)
	if err != nil {
		return nil, err
	}

	return newSender(producer, meta{module: resp.ModuleID}), nil
}
