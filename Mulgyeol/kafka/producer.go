package kafka

import (
    "fmt"
    "github.com/segmentio/kafka-go"
    "context"
)

type KafkaProducer struct {
    Writer *kafka.Writer
}

func NewKafkaProducer(brokerAddress, topic string) *KafkaProducer {
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{brokerAddress},
        Topic:   topic,
    })
    return &KafkaProducer{Writer: writer}
}

func (p *KafkaProducer) SendMessage(ctx context.Context, key, value []byte) error {
    msg := kafka.Message{
        Key:   key,
        Value: value,
    }
    if err := p.Writer.WriteMessages(ctx, msg); err != nil {
        return fmt.Errorf("failed to send message: %w", err)
    }
    return nil
}
