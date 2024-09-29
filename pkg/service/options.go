package service

import (
	"github.com/Egor123qwe/loggy/pkg/model/level"
)

type Options struct {
	Level  level.Level
	Module string

	// Server logger options
	Server []Server // used only if Server is not nil

	// file logger options
	File []File // used only if File is not nil

	// other options
	ToStderr bool
}

type Server struct {
	URL  string
	URLs []string

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
