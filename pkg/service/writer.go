package service

import (
	"context"
	"encoding/json"

	"github.com/Egor123qwe/loggy/pkg/model/log"
	"github.com/Egor123qwe/loggy/pkg/model/msg"
	"github.com/Egor123qwe/loggy/pkg/model/msg/event"
	"github.com/Egor123qwe/loggy/pkg/producer"
	"github.com/Egor123qwe/loggy/pkg/service/logger"
)

type sender struct {
	producer producer.Producer
}

func newSender(producer producer.Producer) logger.Sender {
	return sender{producer: producer}
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
	
	return w.producer.Produce(context.Background(), result)
}
