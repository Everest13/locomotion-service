package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	kafkaGo "github.com/segmentio/kafka-go"
	"log"
)

type Producer struct {
	writer *kafkaGo.Writer
}

func NewProducer(kafkaURL, topic string) (*Producer, error) {
	w := &kafkaGo.Writer{
		Addr:     kafkaGo.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafkaGo.LeastBytes{},
	}
	return &Producer{writer: w}, nil
}

func (p *Producer) WriteMessage(key, value string) error {
	msg := kafkaGo.Message{
		Key:   []byte(key),
		Value: []byte(value),
	}
	err := p.writer.WriteMessages(context.Background(), msg)
	if err != nil {
		j, _ := json.Marshal(msg)
		log.Fatal(fmt.Sprintf("failed to send message: %s", j), err)
		return err
	}

	return nil
}
