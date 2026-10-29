package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/chakhit/backend/models/tables"
)

type peopleRepository struct {
	db *sql.DB
}

func NewPeopleRepository(db *sql.DB) *peopleRepository {
	return &peopleRepository{db: db}
}

type PeopleRepository interface {
	PeopleFindAll() (*[]tables.People, error)
	PeopleAdd(request *tables.People) error
}

func (r *peopleRepository) PeopleFindAll() (*[]tables.People, error) {

	query := `
		SELECT 
		CONVERT(VARCHAR(36), id) AS id,
		first_name,
		last_name,
		date_of_birth,
		address,
		created_at,
		updated_at
		FROM dbo.people
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("search Peoples by criteria: %w", err)
	}
	defer rows.Close()

	var People []tables.People
	for rows.Next() {
		var p tables.People
		err := rows.Scan(
			&p.ID,
			&p.FirstName,
			&p.LastName,
			&p.DateOfBirth,
			&p.Address,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("search Peoples by criteria: %w", err)
		}
		People = append(People, p)
	}

	return &People, nil
}

func (r *peopleRepository) PeopleAdd(person *tables.People) error {
	if person == nil {
		return fmt.Errorf("person is required")
	}

	countQuery := `
		SELECT COUNT(*)
		FROM dbo.people
		WHERE first_name = @p1 AND last_name = @p2
	`

	var existingCount int
	err := r.db.QueryRow(
		countQuery,
		*person.FirstName,
		*person.LastName,
	).Scan(&existingCount)
	if err != nil {
		return fmt.Errorf("check existing person: %w", err)
	}
	if existingCount > 0 {
		return fmt.Errorf("person already exists")
	}

	date, err := time.Parse("2006-01-02", *person.DateOfBirth)
	if err != nil {
		return fmt.Errorf("parse date of birth: %w", err)
	}

	insertQuery := `
		INSERT INTO dbo.people (first_name, last_name, date_of_birth, address)
		VALUES (@p1, @p2, @p3, @p4)
	`

	_, err = r.db.Exec(
		insertQuery,
		*person.FirstName,
		*person.LastName,
		date,
		*person.Address,
	)
	if err != nil {
		return fmt.Errorf("add person: %w", err)
	}

	return nil
}
