package locomotion

import (
	"github.com/go-pg/pg/v10"
)

type repository struct {
	db *pg.DB
}

func newRepository(database *pg.DB) *repository {
	return &repository{
		db: database,
	}
}
