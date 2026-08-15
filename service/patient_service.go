package service

import (
	"fmt"
	"time"

	"github.com/yourname/backend/models/request"
	"github.com/yourname/backend/models/response"
	"github.com/yourname/backend/models/tables"
	"github.com/yourname/backend/repository"
)

type PatientService struct {
	repo repository.PatientRepository
}

// GetPatientByID returns a patient by primary key ID.
func (s *PatientService) GetPatientByID(id string) (*response.PatientResponse, error) {
	data, err := s.repo.PatientFindByID(id)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	res := response.PatientResponse{
		FirstNameTH:  data.FirstNameTH,
		MiddleNameTH: data.MiddleNameTH,
		LastNameTH:   data.LastNameTH,
		FirstNameEN:  data.FirstNameEN,
		MiddleNameEN: data.MiddleNameEN,
		LastNameEN:   data.LastNameEN,
		DateOfBirth:  data.DateOfBirth,
		PatientHN:    data.PatientHN,
		NationalID:   data.NationalID,
		PassportID:   data.PassportID,
		PhoneNumber:  data.PhoneNumber,
		Email:        data.Email,
		Gender:       data.Gender,
	}

	return &res, nil
}

func (s *PatientService) FindPatientsList(request *request.PatientRequest) (*[]response.PatientResponse, error) {
	if request == nil {
		return nil, fmt.Errorf("patient request is required")
	}

	req := &tables.Patient{}

	if request.FirstName != nil {
		req.FirstNameTH = request.FirstName
		req.FirstNameEN = request.FirstName
	}
	if request.MiddleName != nil {
		req.MiddleNameTH = request.MiddleName
		req.MiddleNameEN = request.MiddleName
	}

	if request.LastName != nil {
		req.LastNameTH = request.LastName
		req.LastNameEN = request.LastName
	}
	if request.NationalID != nil {
		req.NationalID = request.NationalID
	}
	if request.PassportID != nil {
		req.PassportID = request.PassportID
	}
	if request.DateOfBirth != nil {
		dob, err := time.Parse("2006-01-02", *request.DateOfBirth)
		if err != nil {
			return nil, fmt.Errorf("invalid date_of_birth: %w", err)
		}
		req.DateOfBirth = &dob
	}
	if request.PhoneNumber != nil {
		req.PhoneNumber = request.PhoneNumber
	}
	if request.Email != nil {
		req.Email = request.Email
	}

	patients, err := s.repo.PatientFindList(req)
	if err != nil {
		return nil, fmt.Errorf("find patients: %w", err)
	}

	if patients == nil || len(*patients) == 0 {
		return nil, nil
	}

	var res []response.PatientResponse
	for _, data := range *patients {
		res = append(res, response.PatientResponse{
			FirstNameTH:  data.FirstNameTH,
			MiddleNameTH: data.MiddleNameTH,
			LastNameTH:   data.LastNameTH,
			FirstNameEN:  data.FirstNameEN,
			MiddleNameEN: data.MiddleNameEN,
			LastNameEN:   data.LastNameEN,
			DateOfBirth:  data.DateOfBirth,
			PatientHN:    data.PatientHN,
			NationalID:   data.NationalID,
			PassportID:   data.PassportID,
			PhoneNumber:  data.PhoneNumber,
			Email:        data.Email,
			Gender:       data.Gender,
		})
	}
	return &res, nil
}
