package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Egor123qwe/loggy/pkg/model/log"
	"github.com/Egor123qwe/loggy/pkg/model/msg"
	"github.com/Egor123qwe/loggy/pkg/model/msg/event"
	"github.com/Egor123qwe/loggy/pkg/producer"
	"github.com/Egor123qwe/loggy/pkg/service/logger"
)

const (
	produceTimeout = 10 * time.Second
)

type sender struct {
	producer producer.Producer

	safe bool
}

func newSender(producer producer.Producer, safe bool) logger.Sender {
	s := sender{
		producer: producer,
		safe:     safe,
	}

	return s
}

func (w sender) Send(log log.Log) error {
	m := msg.MSG{
		Type: string(event.AddLogs),
		Content: []msg.Log{{
			ID:      log.ID,
			TraceID: log.TraceID,
			Time:    log.Time,
			Module:  log.Module,
			Level:   log.Level.String(),
			Message: log.Message,
		}},
	}

	result, err := json.Marshal(m)
	if err != nil {
		return err
	}

	errCh := make(chan error, 1)

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), produceTimeout)
		defer cancel()

		errCh <- w.producer.Produce(ctx, result)
	}()

	if w.safe {
		return <-errCh
	}

	return nil
}
