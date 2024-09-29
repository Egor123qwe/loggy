package logger

import (
	"errors"
	"strings"
	"time"

	"github.com/Egor123qwe/loggy/pkg/model/level"
	"github.com/Egor123qwe/loggy/pkg/model/log"
)

func ParseLog(s string) (log.Log, error) {
	parts := strings.Split(s, ": ")

	if len(parts) != 4 {
		return log.Log{}, errors.New("invalid log format")
	}

	timeString := parts[0]

	time, err := time.Parse("2006-01-02 15:04:05.000", timeString)
	if err != nil {
		return log.Log{}, err
	}

	// Extract the level and module part
	levelString := strings.TrimLeft(parts[1], "(")
	levelString = strings.SplitN(levelString, ")", 2)[0]

	traceID := parts[2]
	message := parts[3]

	// Create a new Log struct
	log := log.Log{
		TraceID: traceID,

		Time:  time,
		Level: level.Parse(levelString),

		Message: message,
	}

	return log, nil
}
