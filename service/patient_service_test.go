package service

import (
	"errors"
	"testing"
	"time"

	"github.com/yourname/backend/models/request"
	"github.com/yourname/backend/models/tables"
)

type MockPatientRepository struct {
	mockPatientFindByID func(id string) (*tables.Patient, error)
	mockPatientFindList func(request *tables.Patient) (*[]tables.Patient, error)

	patient  map[string]*tables.Patient
	patients map[string]*[]tables.Patient
}

func NewMockPatientRepository() *MockPatientRepository {
	return &MockPatientRepository{
		patient:  make(map[string]*tables.Patient),
		patients: make(map[string]*[]tables.Patient),
	}
}

func (m *MockPatientRepository) PatientFindByID(id string) (*tables.Patient, error) {
	if m.mockPatientFindByID != nil {
		return m.mockPatientFindByID(id)
	}
	data, ok := m.patient[id]
	if !ok {
		return nil, errors.New("patient not found")
	}
	return data, nil
}

func (m *MockPatientRepository) PatientFindList(request *tables.Patient) (*[]tables.Patient, error) {
	if m.mockPatientFindList != nil {
		return m.mockPatientFindList(request)
	}

	if request == nil {
		return nil, errors.New("patient not found")
	}
	for _, data := range m.patients {
		return data, nil
	}
	return nil, errors.New("patient not found")
}

func TestGetPatientByID_Success(t *testing.T) {
	mockRepo := NewMockPatientRepository()

	firstName := "Somchai"
	lastName := "Phan"
	nationalID := "1234567890123"
	nhn := "HN-001"
	dob := time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC)

	mockRepo.patient[nationalID] = &tables.Patient{
		FirstNameTH: &firstName,
		LastNameTH:  &lastName,
		NationalID:  &nationalID,
		PatientHN:   &nhn,
		DateOfBirth: &dob,
	}

	service := PatientService{repo: mockRepo}

	patient, err := service.GetPatientByID(nationalID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if patient == nil {
		t.Fatal("expected patient, got nil")
	}
	if patient.FirstNameTH == nil || *patient.FirstNameTH != firstName {
		t.Fatalf("expected first name %q, got %#v", firstName, patient.FirstNameTH)
	}
	if patient.NationalID == nil || *patient.NationalID != nationalID {
		t.Fatalf("expected national id %q, got %#v", nationalID, patient.NationalID)
	}
	if patient.DateOfBirth == nil || !patient.DateOfBirth.Equal(dob) {
		t.Fatalf("expected date of birth %v, got %#v", dob, patient.DateOfBirth)
	}
}

func TestGetPatientByID_Failure_NotFound(t *testing.T) {
	mockRepo := NewMockPatientRepository()

	service := PatientService{repo: mockRepo}

	_, err := service.GetPatientByID("xxx")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestPatientFindList_Success(t *testing.T) {
	mockRepo := NewMockPatientRepository()
	service := PatientService{repo: mockRepo}

	lastName := "Phan"
	firstName1 := "Somchai1"
	firstName2 := "Somchai2"
	nationalID1 := "1234567890123"
	nationalID2 := "1234567890124"
	nhn1 := "HN-001"
	nhn2 := "HN-002"
	dob1 := time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC)
	dob2 := time.Date(1991, time.February, 20, 0, 0, 0, 0, time.UTC)

	patients := []tables.Patient{
		{
			FirstNameTH: &firstName1,
			LastNameTH:  &lastName,
			NationalID:  &nationalID1,
			PatientHN:   &nhn1,
			DateOfBirth: &dob1,
		},
		{
			FirstNameTH: &firstName2,
			LastNameTH:  &lastName,
			NationalID:  &nationalID2,
			PatientHN:   &nhn2,
			DateOfBirth: &dob2,
		},
	}

	mockRepo.mockPatientFindList = func(req *tables.Patient) (*[]tables.Patient, error) {
		if req == nil || req.LastNameTH == nil || *req.LastNameTH != lastName {
			return nil, errors.New("unexpected search request")
		}
		return &patients, nil
	}

	res, err := service.FindPatientsList(&request.PatientRequest{LastName: &lastName})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if res == nil || len(*res) != 2 {
		t.Fatalf("expected 2 patients, got %#v", res)
	}
	if (*res)[0].FirstNameTH == nil || *(*res)[0].FirstNameTH != firstName1 {
		t.Fatalf("expected first patient name %q, got %#v", firstName1, (*res)[0].FirstNameTH)
	}
	if (*res)[1].PatientHN == nil || *(*res)[1].PatientHN != nhn2 {
		t.Fatalf("expected second patient HN %q, got %#v", nhn2, (*res)[1].PatientHN)
	}
}

func TestPatientFindList_Failure_Invalid(t *testing.T) {
	mockRepo := NewMockPatientRepository()
	service := PatientService{repo: mockRepo}

	_, err := service.FindPatientsList(nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestPatientFindList_Failure_NotFound(t *testing.T) {
	mockRepo := NewMockPatientRepository()
	lastName := "xxx"
	service := PatientService{repo: mockRepo}

	_, err := service.FindPatientsList(&request.PatientRequest{LastName: &lastName})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
