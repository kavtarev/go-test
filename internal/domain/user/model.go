package user

import (
	"fmt"
	"time"
)

type UserDomain struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Passport  string    `json:"passport"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *UserDomain) GetFullData() string {
	return fmt.Sprintf("%v %v %v", u.Name, u.Email, u.Passport)
}
