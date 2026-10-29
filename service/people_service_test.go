package service

// import (
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/chakhit/backend/models/request"
// 	"github.com/chakhit/backend/models/tables"
// )

// type MockPeopleRepository struct {
// 	mockPeopleFindByID func(id string) (*tables.People, error)
// 	mockPeopleFindList func(request *tables.People) (*[]tables.People, error)

// 	People  map[string]*tables.People
// 	Peoples map[string]*[]tables.People
// }

// func NewMockPeopleRepository() *MockPeopleRepository {
// 	return &MockPeopleRepository{
// 		People:  make(map[string]*tables.People),
// 		Peoples: make(map[string]*[]tables.People),
// 	}
// }

// func (m *MockPeopleRepository) PeopleFindByID(id string) (*tables.People, error) {
// 	if m.mockPeopleFindByID != nil {
// 		return m.mockPeopleFindByID(id)
// 	}
// 	data, ok := m.People[id]
// 	if !ok {
// 		return nil, errors.New("People not found")
// 	}
// 	return data, nil
// }

// func (m *MockPeopleRepository) PeopleFindList(request *tables.People) (*[]tables.People, error) {
// 	if m.mockPeopleFindList != nil {
// 		return m.mockPeopleFindList(request)
// 	}

// 	if request == nil {
// 		return nil, errors.New("People not found")
// 	}
// 	for _, data := range m.Peoples {
// 		return data, nil
// 	}
// 	return nil, errors.New("People not found")
// }

// func TestGetPeopleByID_Success(t *testing.T) {
// 	mockRepo := NewMockPeopleRepository()

// 	firstName := "Somchai"
// 	lastName := "Phan"
// 	nationalID := "1234567890123"
// 	nhn := "HN-001"
// 	dob := time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC)

// 	mockRepo.People[nationalID] = &tables.People{
// 		FirstNameTH: &firstName,
// 		LastNameTH:  &lastName,
// 		NationalID:  &nationalID,
// 		PeopleHN:   &nhn,
// 		DateOfBirth: &dob,
// 	}

// 	service := PeopleService{repo: mockRepo}

// 	People, err := service.GetPeopleByID(nationalID)
// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}
// 	if People == nil {
// 		t.Fatal("expected People, got nil")
// 	}
// 	if People.FirstNameTH == nil || *People.FirstNameTH != firstName {
// 		t.Fatalf("expected first name %q, got %#v", firstName, People.FirstNameTH)
// 	}
// 	if People.NationalID == nil || *People.NationalID != nationalID {
// 		t.Fatalf("expected national id %q, got %#v", nationalID, People.NationalID)
// 	}
// 	if People.DateOfBirth == nil || !People.DateOfBirth.Equal(dob) {
// 		t.Fatalf("expected date of birth %v, got %#v", dob, People.DateOfBirth)
// 	}
// }

// func TestGetPeopleByID_Failure_NotFound(t *testing.T) {
// 	mockRepo := NewMockPeopleRepository()

// 	service := PeopleService{repo: mockRepo}

// 	_, err := service.GetPeopleByID("xxx")
// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}
// }

// func TestPeopleFindList_Success(t *testing.T) {
// 	mockRepo := NewMockPeopleRepository()
// 	service := PeopleService{repo: mockRepo}

// 	lastName := "Phan"
// 	firstName1 := "Somchai1"
// 	firstName2 := "Somchai2"
// 	nationalID1 := "1234567890123"
// 	nationalID2 := "1234567890124"
// 	nhn1 := "HN-001"
// 	nhn2 := "HN-002"
// 	dob1 := time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC)
// 	dob2 := time.Date(1991, time.February, 20, 0, 0, 0, 0, time.UTC)

// 	Peoples := []tables.People{
// 		{
// 			FirstNameTH: &firstName1,
// 			LastNameTH:  &lastName,
// 			NationalID:  &nationalID1,
// 			PeopleHN:   &nhn1,
// 			DateOfBirth: &dob1,
// 		},
// 		{
// 			FirstNameTH: &firstName2,
// 			LastNameTH:  &lastName,
// 			NationalID:  &nationalID2,
// 			PeopleHN:   &nhn2,
// 			DateOfBirth: &dob2,
// 		},
// 	}

// 	mockRepo.mockPeopleFindList = func(req *tables.People) (*[]tables.People, error) {
// 		if req == nil || req.LastNameTH == nil || *req.LastNameTH != lastName {
// 			return nil, errors.New("unexpected search request")
// 		}
// 		return &Peoples, nil
// 	}

// 	res, err := service.FindPeoplesList(&request.PeopleRequest{LastName: &lastName})
// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}
// 	if res == nil || len(*res) != 2 {
// 		t.Fatalf("expected 2 Peoples, got %#v", res)
// 	}
// 	if (*res)[0].FirstNameTH == nil || *(*res)[0].FirstNameTH != firstName1 {
// 		t.Fatalf("expected first People name %q, got %#v", firstName1, (*res)[0].FirstNameTH)
// 	}
// 	if (*res)[1].PeopleHN == nil || *(*res)[1].PeopleHN != nhn2 {
// 		t.Fatalf("expected second People HN %q, got %#v", nhn2, (*res)[1].PeopleHN)
// 	}
// }

// func TestPeopleFindList_Failure_Invalid(t *testing.T) {
// 	mockRepo := NewMockPeopleRepository()
// 	service := PeopleService{repo: mockRepo}

// 	_, err := service.FindPeoplesList(nil)
// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}
// }

// func TestPeopleFindList_Failure_NotFound(t *testing.T) {
// 	mockRepo := NewMockPeopleRepository()
// 	lastName := "xxx"
// 	service := PeopleService{repo: mockRepo}

// 	_, err := service.FindPeoplesList(&request.PeopleRequest{LastName: &lastName})
// 	if err == nil {
// 		t.Fatal("expected error, got nil")
// 	}
// }
