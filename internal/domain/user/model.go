package user

import (
	"fmt"
	"time"
)

type UserDomain struct {
	Id        string
	Name      string
	Email     string
	Passport  string
	CreatedAt time.Time
}

func (u *UserDomain) GetFullData() string {
	return fmt.Sprintf("%v %v %v", u.Name, u.Email, u.Passport)
}
