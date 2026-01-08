package user

import "fmt"

type UserDomain struct {
	Id       string
	Name     string
	Email    string
	Passport string
}

func (u *UserDomain) GetFullData() string {
	return fmt.Sprintf("%v %v %v", u.Name, u.Email, u.Passport)
}
