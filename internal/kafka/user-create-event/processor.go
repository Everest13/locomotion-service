package user_create_event

import (
	"context"
	"encoding/json"
	kafkaGo "github.com/segmentio/kafka-go"
	"locomotion-service/internal/service/user"
)

type Processor struct {
	userService *user.Service
}

func NewProcessor(userService *user.Service) *Processor {
	return &Processor{userService: userService}
}

func (m *Processor) Process(_ context.Context, msg *kafkaGo.Message) (err error) {
	data := &user.User{}
	if err := json.Unmarshal(msg.Value, data); err != nil {
		return nil
	}

	err = m.userService.CreateUser(data)
	if err != nil {
		return err
	}

	return nil
}
