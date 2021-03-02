package subscribe

import "time"

type Subscribe struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
