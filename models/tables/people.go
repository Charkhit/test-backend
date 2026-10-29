package tables

import (
	"time"
)

type People struct {
	ID          *string    `json:"id"`
	FirstName   *string    `json:"first_name"`
	LastName    *string    `json:"last_name"`
	DateOfBirth *string    `json:"date_of_birth"`
	Address     *string    `json:"address"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}
