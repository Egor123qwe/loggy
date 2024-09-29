package logger

import (
	"io"
	"os"

	"github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
)

const formatString = "%{time:2006-01-02 15:04:05.000}: (%{level}) traceID: %{module}: %{message}"

type Service interface {
	New(traceID string) *logging.Logger

	Close() error
}

type service struct {
	writers []io.WriteCloser
}

func New(opts Options) Service {
	srv := &service{}

	level := logging.Level(opts.Level)
	format := logging.MustStringFormatter(formatString)

	backends := make([]logging.Backend, 0)

	// set senders
	for _, sender := range opts.Sender {
		writer := srv.newWriter(sender)

		backends = append(backends, add(writer, level, format))
	}

	// set files
	for _, f := range opts.File {
		file := &lumberjack.Logger{
			Filename:   f.Name,
			MaxSize:    f.MaxSizeMb,
			MaxBackups: f.MaxFiles,
		}

		backends = append(backends, add(file, level, format))
	}

	// set console
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

func (s *service) New(traceID string) *logging.Logger {
	return logging.MustGetLogger(traceID)
}

func (s *service) Close() error {
	var err error

	for _, w := range s.writers {
		if err == nil {
			err = w.Close()

			continue
		}

		w.Close()
	}

	s.writers = nil

	return err
}
