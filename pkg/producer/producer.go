package producer

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

const topic = "store"

const (
	dialerTimeout = 10 * time.Second
)

type Producer interface {
	Produce(ctx context.Context, m []byte) error
	Close() error
}

type producer struct {
	writer *kafka.Writer
}

type Credentials struct {
	Username string
	Password string
}

func New(credentials *Credentials, brokers []string) (Producer, error) {
	var dialer *kafka.Dialer

	if credentials != nil {
		mechanism, err := scram.Mechanism(scram.SHA256, credentials.Username, credentials.Password)
		if err != nil {
			return nil, fmt.Errorf("failed to create scram mechanism: %w", err)
		}

		dialer = &kafka.Dialer{
			Timeout:       dialerTimeout,
			DualStack:     true,
			SASLMechanism: mechanism,
		}
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Dialer:   dialer,
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	return producer{writer: writer}, nil
}

func (p producer) Produce(ctx context.Context, m []byte) error {
	err := p.writer.WriteMessages(ctx,
		kafka.Message{Value: m},
	)

	if err != nil {
		err = fmt.Errorf("failed to produce message: %w", err)
	}

	return err
}

func (p producer) Close() error {
	return p.writer.Close()
}
