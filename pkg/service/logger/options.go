package logger

import (
	model "github.com/Egor123qwe/loggy/pkg/model/log"
)

type Options struct {
	Level  model.Level
	Module string

	// Writer logger options
	Sender Sender // used only if Writer is not nil

	// file logger options
	File *File // used only if File is not nil

	// other options
	ToStderr bool
}

type Sender interface {
	Send(log model.Log) error
}

type File struct {
	Name      string
	MaxSizeMb int
	MaxFiles  int
}
