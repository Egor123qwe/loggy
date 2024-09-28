package logger

import (
	"errors"
	"strings"
	"time"

	model "github.com/Egor123qwe/loggy/pkg/model/log"
	"github.com/op/go-logging"
)

func ParseLog(s string) (model.Log, error) {
	parts := strings.Split(s, ": ")

	if len(parts) != 4 {
		return model.Log{}, errors.New("invalid log format")
	}

	timeString := parts[0]

	time, err := time.Parse("2006-01-02 15:04:05.000", timeString)
	if err != nil {
		return model.Log{}, err
	}

	// Extract the level and module part
	levelString := strings.TrimLeft(parts[1], "(")
	levelString = strings.SplitN(levelString, ")", 2)[0]

	level := model.ConvertLevelName(levelString)

	traceID := parts[2]
	message := parts[3]

	// Create a new Log struct
	log := model.Log{
		TraceID: traceID,

		Time:    time,
		Module:  "TODO",
		Level:   level,
		Message: message,
	}

	return log, nil
}

func parseLevel(s model.Level) logging.Level {
	switch s {
	case model.DEBUG:
		return logging.DEBUG

	case model.INFO:
		return logging.INFO

	case model.NOTICE:
		return logging.NOTICE

	case model.WARNING:
		return logging.WARNING

	case model.ERROR:
		return logging.ERROR

	case model.CRITICAL:
		return logging.CRITICAL

	default:
		return logging.DEBUG
	}
}
