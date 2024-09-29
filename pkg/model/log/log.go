package log

import (
	"time"

	"github.com/Egor123qwe/loggy/pkg/model/level"
)

type Log struct {
	ID      uint64
	TraceID string

	Time  time.Time
	Level level.Level

	Message string
}
