package logger

import (
	"github.com/Egor123qwe/loggy/pkg/model/level"
	"github.com/Egor123qwe/loggy/pkg/model/log"
)

type Options struct {
	Level level.Level

	// Writer logger options
	Sender []Sender // used only if Writer is not nil

	// file logger options
	File []File // used only if File is not nil

	// other options
	ToStderr bool
}

type Sender interface {
	Send(log log.Log) error
}

type File struct {
	Name      string
	MaxSizeMb int
	MaxFiles  int
}
