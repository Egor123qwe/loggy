package level

import "strings"

type Level int

const (
	CRITICAL Level = iota
	ERROR
	WARNING
	NOTICE
	INFO
	DEBUG

	Invalid = -1
)

var levelNames = []string{
	"CRITICAL",
	"ERROR",
	"WARNING",
	"NOTICE",
	"INFO",
	"DEBUG",
}

func Parse(name string) Level {
	for i, levelName := range levelNames {
		if strings.EqualFold(levelName, name) {
			return Level(i)
		}
	}

	return Invalid
}

func (l Level) String() string {
	return levelNames[l]
}
