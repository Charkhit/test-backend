package response

import "time"

type PeopleResponse struct {
	ID             *string   `json:"id"`
	FirstName      *string   `json:"firstName"`
	LastName       *string   `json:"lastName"`
	DateBirth      *string   `json:"birthDate"` // Format: YYYY-MM-DD
	Address        *string   `json:"address"`
	CreateDateTime time.Time `json:"createDateTime"`
}
