package service

import (
	"github.com/Egor123qwe/loggy/pkg/model/level"
)

// Level implement Level here to more convenient use this module as a library
type Level = level.Level

const (
	CRITICAL = level.CRITICAL
	ERROR    = level.ERROR
	WARNING  = level.WARNING
	NOTICE   = level.NOTICE
	INFO     = level.INFO
	DEBUG    = level.DEBUG

	Invalid = -1
)

type Options struct {
	Level  Level
	Module string

	// Server logger options
	Server []Server // used only if Server is not nil

	// file logger options
	File []File // used only if File is not nil

	// other options
	ToStderr bool
}

type Server struct {
	URL              string
	LogsChannelsURLs []string

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
