package dtos

import "time"

type UserDto struct {
	Id               int
	Login            string
	Name             string
	Surname          string
	Status           string
	Role             string
	RegistrationDate time.Time
	UpdateDate       time.Time
}
