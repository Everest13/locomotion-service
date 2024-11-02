package locomotion

import (
	"github.com/go-pg/pg/v10"
)

type Service struct {
	repo *repository
}

func NewService(database *pg.DB) *Service {
	repo := newRepository(database)
	return &Service{
		repo: repo,
	}
}
