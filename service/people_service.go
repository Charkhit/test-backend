package service

import (
	"fmt"
	"log"
	"time"

	"github.com/chakhit/backend/models/request"
	"github.com/chakhit/backend/models/response"
	"github.com/chakhit/backend/models/tables"
	"github.com/chakhit/backend/repository"
)

type PeopleService struct {
	repo repository.PeopleRepository
}

func NewPeopleService(repo repository.PeopleRepository) *PeopleService {
	return &PeopleService{repo: repo}
}

func (s *PeopleService) GetPeopleList() (*[]response.PeopleResponse, error) {
	data, err := s.repo.PeopleFindAll()
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, nil
	}

	var res []response.PeopleResponse
	for _, d := range *data {
		res = append(res, response.PeopleResponse{
			ID:             d.ID,
			FirstName:      d.FirstName,
			LastName:       d.LastName,
			DateBirth:      convert(d.DateOfBirth),
			Address:        d.Address,
			CreateDateTime: d.CreatedAt,
		})
	}

	return &res, nil
}

func convert(input *string) *string {
	if input == nil {
		return nil
	}
	t, err := time.Parse("2006-01-02T15:04:05Z", *input)
	if err != nil {
		log.Fatal(err)
	}
	formatted := t.Format("2006-01-02")
	return &formatted
}

func (s *PeopleService) AddPeoples(request *request.PeopleAddRequest) error {
	if request == nil {
		return fmt.Errorf("People request is required")
	}

	req := &tables.People{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		DateOfBirth: request.DateBirth,
		Address:     request.Address,
	}

	err := s.repo.PeopleAdd(req)
	if err != nil {
		return fmt.Errorf("add Peoples: %w", err)
	}

	return nil
}
