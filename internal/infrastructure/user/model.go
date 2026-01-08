package user

import (
	"go-test/internal/domain/user"
	"time"
)

type UserEntity struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Passport  string    `db:"passport"`
	CreatedAt time.Time `db:"created_at"`
}

func (u *UserEntity) ToDomain() user.UserDomain {
	return user.UserDomain{
		Id:        u.Id,
		Name:      u.Name,
		Email:     u.Email,
		Passport:  u.Passport,
		CreatedAt: u.CreatedAt,
	}
}
