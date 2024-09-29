package logger

import (
	"io"
)

type writer struct {
	sender Sender
}

func (s *service) newWriter(sender Sender) io.WriteCloser {
	w := writer{
		sender: sender,
	}

	s.writers = append(s.writers, w)

	return w
}

func (w writer) Write(p []byte) (n int, err error) {
	log, err := ParseLog(string(p))
	if err != nil {
		return 0, err
	}

	if err := w.sender.Send(log); err != nil {
		return 0, err
	}

	return len(p), nil
}

func (w writer) Close() error {
	return w.sender.Close()
}
