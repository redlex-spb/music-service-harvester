package gateway

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type Producer interface {
	Send(topic string, v any) error
	Close() error
}

type kafkaProducer struct{ w *kafka.Writer }

func NewKafkaProducer(brokers []string) Producer {
	return &kafkaProducer{w: &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Balancer: &kafka.LeastBytes{},
	}}
}

func (k *kafkaProducer) Send(topic string, v any) error {
	data, _ := json.Marshal(v)
	return k.w.WriteMessages(context.Background(), kafka.Message{Topic: topic, Value: data})
}
func (k *kafkaProducer) Close() error { return k.w.Close() }
