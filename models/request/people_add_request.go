package request

type PeopleAddRequest struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	DateBirth *string `json:"birthDate"` // Format: YYYY-MM-DD
	Address   *string `json:"address"`
}
