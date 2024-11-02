package user

import (
	"context"
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

func (r *repository) getUser(ctx context.Context, id int64) (*User, error) {
	user := &User{ID: id}
	err := r.db.Model(user).WherePK().Select()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) createUser(user *User) error {
	user.SetCreatedAt()
	_, err := r.db.Model(user).Insert()
	if err != nil {
		return err
	}

	return nil
}

//todo check and remove
// createSchema creates database schema for User and Story models.
//func createSchema(db *pg.DB) error {
//	models := []interface{}{
//		(*User)(nil),
//	}
//
//	for _, model := range models {
//		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
//			Temp: true,
//		})
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
