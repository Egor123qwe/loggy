package service

import (
	model "github.com/Egor123qwe/loggy/pkg/model/log"
)

type Options struct {
	Level  model.Level
	Module string

	// Server logger options
	Server *Server // used only if Server is not nil

	// file logger options
	File *File // used only if File is not nil

	// other options
	ToStderr bool
}

type Server struct {
	URLs     []string
	WithWait bool // log func will wait for producer send log

	// Credentials are optional
	Credentials *Credentials
}

type File struct {
	Name      string
	MaxSizeMb int
	MaxFiles  int
}

type Credentials struct {
	Username string
	Password string
}
