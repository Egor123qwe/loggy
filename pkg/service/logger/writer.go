package logger

import (
	"io"
)

type writer struct {
	sender Sender

	meta Meta
}

type Meta struct {
	Module string
}

func (s service) newWriter(sender Sender, meta Meta) io.Writer {
	w := writer{
		sender: sender,
		meta:   meta,
	}

	return w
}

func (w writer) Write(p []byte) (n int, err error) {
	log, err := ParseLog(string(p))
	if err != nil {
		return 0, err
	}

	// set some meta data
	log.Module = w.meta.Module

	if err := w.sender.Send(log); err != nil {
		return 0, err
	}

	return len(p), nil
}
