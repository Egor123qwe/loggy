package msg

import (
	"time"
)

type Log struct {
	ID      uint64 `json:"id,omitempty"`
	TraceID string `json:"trace_id"`

	Time     time.Time `json:"time"`
	ModuleID int64     `json:"module_id"`
	Level    string    `json:"level"`

	Message string `json:"message"`
}
