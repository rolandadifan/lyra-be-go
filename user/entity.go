package user

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Token struct {
	ID        int
	UserID    int
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
