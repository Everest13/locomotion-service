package kafka

import (
	"context"
	"fmt"
	kafkaGo "github.com/segmentio/kafka-go"
	"log"
)

type MessageProcessor interface {
	Process(ctx context.Context, msg *kafkaGo.Message) error
}

type Consumer struct {
	reader     *kafkaGo.Reader
	mProcessor MessageProcessor
}

func NewConsumer(brokers, topic string, partition int, processor MessageProcessor) (*Consumer, error) {
	r := kafkaGo.NewReader(kafkaGo.ReaderConfig{
		Brokers:   []string{brokers},
		Topic:     topic,
		Partition: partition,
		MaxBytes:  10e6, // 10MB
	})

	return &Consumer{
		reader:     r,
		mProcessor: processor,
	}, nil
}

func (c *Consumer) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			for {
				msg, err := c.reader.ReadMessage(context.Background())
				if err != nil {
					break
				}

				//todo
				fmt.Printf("message at offset %d: %s = %s\n", msg.Offset, string(msg.Key), string(msg.Value))

				err = c.mProcessor.Process(ctx, &msg)
				if err != nil {
					//todo
					break
				}
			}

			if err := c.reader.Close(); err != nil {
				//todo
				log.Fatal("failed to close reader:", err)
			}
		}
	}
}
