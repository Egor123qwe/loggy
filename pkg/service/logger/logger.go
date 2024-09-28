package logger

import (
	"io"
	"os"

	"github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
)

const formatString = "%{time:2006-01-02 15:04:05.000}: (%{level}) traceID: %{module}: %{message}"

type Service interface {
	Get(traceID string) *logging.Logger
}

type service struct{}

func New(opts Options) Service {
	srv := &service{}

	level := parseLevel(opts.Level)
	format := logging.MustStringFormatter(formatString)

	backends := make([]logging.Backend, 0)

	if opts.Sender != nil {
		writer := srv.newWriter(opts.Sender, Meta{Module: opts.Module})

		backends = append(backends, add(writer, level, format))
	}

	if opts.File != nil {
		file := &lumberjack.Logger{
			Filename:   opts.File.Name,
			MaxSize:    opts.File.MaxSizeMb,
			MaxBackups: opts.File.MaxFiles,
		}

		backends = append(backends, add(file, level, format))
	}

	if opts.ToStderr {
		backends = append(backends, add(os.Stderr, level, format))
	}

	logging.SetBackend(backends...)

	return srv
}

func add(out io.Writer, level logging.Level, format logging.Formatter) logging.Backend {
	backend := logging.NewLogBackend(out, "", 0)

	formatter := logging.NewBackendFormatter(backend, format)

	leveled := logging.AddModuleLevel(formatter)
	leveled.SetLevel(level, "")

	return leveled
}

func (s service) Get(traceID string) *logging.Logger {
	return logging.MustGetLogger(traceID)
}
