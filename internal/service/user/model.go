package user

import "time"

type SexType string

const (
	UnknownSexType = SexType("UNKNOWN")
	FemaleSexType  = SexType("FEMALE")
	MaleSexType    = SexType("MALE")
)

type User struct {
	tableName  struct{}  `pg:","`
	ID         int64     `pg:"id"`
	FirstName  string    `pg:"first_name"`
	SecondName string    `pg:"second_name"`
	Patronymic *string   `pg:"patronymic"`
	Age        int64     `pg:"age"`
	Sex        SexType   `pg:"sex"`
	CreatedAt  time.Time `pg:"created_at"`
	DateBirth  time.Time `pg:"date_birth"`
}

func (u *User) SetCreatedAt() {
	u.CreatedAt = time.Now()
}
