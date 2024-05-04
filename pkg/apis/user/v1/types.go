package v0alpha1

import (
	"time"
)

type User struct {
	Id       string
	Username string
	Email    string

	Password string
	Salt     string

	Active    bool
	Token     string
	RoleID    int
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
