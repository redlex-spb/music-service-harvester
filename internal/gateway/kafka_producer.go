package gateway

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type Producer interface{ Send(string, any) error }
type kafkaProducer struct{ w *kafka.Writer }

func NewKafkaProducer(b []string) Producer {
	return &kafkaProducer{w: &kafka.Writer{Addr: kafka.TCP(b...), Balancer: &kafka.LeastBytes{}}}
}
func (k *kafkaProducer) Send(t string, v any) error {
	b, _ := json.Marshal(v)
	return k.w.WriteMessages(context.Background(), kafka.Message{Topic: t, Value: b})
}
