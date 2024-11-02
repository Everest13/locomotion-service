package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg/v10"
	"locomotion-service/internal/utility/kafka"
	"log"
)

type Service struct {
	repo                    *repository
	quickStartEventProducer *kafka.Producer
}

func NewService(database *pg.DB, quickStartEventProducer *kafka.Producer) *Service {
	repo := newRepository(database)
	return &Service{
		repo:                    repo,
		quickStartEventProducer: quickStartEventProducer,
	}
}

func (s *Service) GetUser(ctx context.Context, id int64) (*User, error) {
	user, err := s.repo.getUser(ctx, id)
	if err != nil {
		return nil, err
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Fatalln("failed to marshal created user:", err)
		return nil, err
	}

	msg := fmt.Sprintf(string(userJson))
	err = s.quickStartEventProducer.WriteMessage("get user", msg)
	if err != nil {
		log.Fatalln("failed to send kafka msg:", err)
		return nil, err
	}

	return user, nil
}

func (s *Service) CreateUser(user *User) error {
	err := s.repo.createUser(user)
	if err != nil {
		return err
	}

	return nil
}
