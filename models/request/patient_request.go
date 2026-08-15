package request

type PatientRequest struct {
	NationalID  *string `json:"national_id"`
	PassportID  *string `json:"passport_id"`
	FirstName   *string `json:"first_name"`
	MiddleName  *string `json:"middle_name"` // ใช้ omitempty หรือ *string หากต้องการให้ส่ง null ได้
	LastName    *string `json:"last_name"`
	DateOfBirth *string `json:"date_of_birth"` // Format: YYYY-MM-DD
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
}
