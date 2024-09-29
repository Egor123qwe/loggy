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
	meta     meta

	closer context.CancelFunc
	logCh  chan []byte
}

type meta struct {
	module int64
}

func newSender(producer producer.Producer, meta meta) logger.Sender {
	s := &sender{
		producer: producer,
		meta:     meta,

		logCh: make(chan []byte, 1024),
	}

	ctx, cancel := context.WithCancel(context.Background())
	s.closer = cancel

	go s.serveLogs(ctx)

	return s
}

func (w *sender) Send(log log.Log) error {
	m := msg.MSG{
		Type: string(event.AddLogs),
		Content: []msg.Log{{
			ID:       log.ID,
			TraceID:  log.TraceID,
			Time:     log.Time,
			ModuleID: w.meta.module,
			Level:    log.Level.String(),
			Message:  log.Message,
		}},
	}

	result, err := json.Marshal(m)
	if err != nil {
		return err
	}

	w.logCh <- result

	return nil
}

func (w *sender) serveLogs(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case b := <-w.logCh:
			produceCtx, _ := context.WithTimeout(context.Background(), produceTimeout)

			w.producer.Produce(produceCtx, b)
		}
	}
}

func (w *sender) Close() error {
	w.closer()
	close(w.logCh)

	for log := range w.logCh {
		produceCtx, _ := context.WithTimeout(context.Background(), produceTimeout)

		w.producer.Produce(produceCtx, log)
	}

	return w.producer.Close()
}
