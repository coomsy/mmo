package user

import (
	"time"
)

type User struct {
	ID       string
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
