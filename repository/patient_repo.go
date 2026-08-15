package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/yourname/backend/database"
	"github.com/yourname/backend/models/tables"
)

type PatientRepository interface {
	PatientFindByID(id string) (*tables.Patient, error)
	PatientFindList(request *tables.Patient) (*[]tables.Patient, error)
}

func PatientFindByID(id string) (*tables.Patient, error) {
	var patient tables.Patient
	query := `
		SELECT *
		FROM patient
		WHERE national_id = $1 OR passport_id = $1
		LIMIT 1
	`

	row := database.DB.QueryRow(query, id)
	err := row.Scan(
		&patient.FirstNameTH,
		&patient.MiddleNameTH,
		&patient.LastNameTH,
		&patient.FirstNameEN,
		&patient.MiddleNameEN,
		&patient.LastNameEN,
		&patient.DateOfBirth,
		&patient.PatientHN,
		&patient.NationalID,
		&patient.PassportID,
		&patient.PhoneNumber,
		&patient.Email,
		&patient.Gender,
		&patient.CreatedAt,
		&patient.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("search patients by criteria: %w", err)
	}
	return &patient, nil
}

func PatientFindList(request *tables.Patient) (*[]tables.Patient, error) {
	conditionsAnd := []string{"1=1"} // Always true condition to simplify query building
	conditionsOr := []string{}
	args := []any{}
	index := 1

	if request.FirstNameTH != nil {
		conditionsOr = append(conditionsOr, fmt.Sprintf("first_name_th = $%d", index))
		args = append(args, request.FirstNameTH)
		index++
	}
	if request.MiddleNameTH != nil {
		conditionsOr = append(conditionsOr, fmt.Sprintf("middle_name_th = $%d", index))
		args = append(args, request.MiddleNameTH)
		index++
	}
	if request.LastNameTH != nil {
		conditionsOr = append(conditionsOr, fmt.Sprintf("last_name_th = $%d", index))
		args = append(args, request.LastNameTH)
		index++
	}
	if request.FirstNameEN != nil {
		conditionsOr = append(conditionsOr, fmt.Sprintf("first_name_en = $%d", index))
		args = append(args, request.FirstNameEN)
		index++
	}
	if request.MiddleNameEN != nil {
		conditionsOr = append(conditionsOr, fmt.Sprintf("middle_name_en = $%d", index))
		args = append(args, request.MiddleNameEN)
		index++
	}
	if request.LastNameEN != nil {
		conditionsOr = append(conditionsOr, fmt.Sprintf("last_name_en = $%d", index))
		args = append(args, request.LastNameEN)
		index++
	}

	if request.NationalID != nil {
		conditionsAnd = append(conditionsAnd, fmt.Sprintf("national_id = $%d", index))
		args = append(args, request.NationalID)
		index++
	}
	if request.PassportID != nil {
		conditionsAnd = append(conditionsAnd, fmt.Sprintf("passport_id = $%d", index))
		args = append(args, request.PassportID)
		index++
	}
	if request.DateOfBirth != nil {
		conditionsAnd = append(conditionsAnd, fmt.Sprintf("date_of_birth = $%d", index))
		args = append(args, request.DateOfBirth)
		index++
	}
	if request.PhoneNumber != nil {
		conditionsAnd = append(conditionsAnd, fmt.Sprintf("phone_number = $%d", index))
		args = append(args, request.PhoneNumber)
		index++
	}
	if request.Email != nil {
		conditionsAnd = append(conditionsAnd, fmt.Sprintf("email = $%d", index))
		args = append(args, request.Email)
		index++
	}

	query := "SELECT * FROM patient WHERE " + strings.Join(conditionsOr, " OR ") + " AND " + strings.Join(conditionsAnd, " AND ") + " LIMIT 50"
	fmt.Println(query)
	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("search patients by criteria: %w", err)
	}
	defer rows.Close()

	patients := []tables.Patient{}
	for rows.Next() {
		var patient tables.Patient
		if err := rows.Scan(
			&patient.FirstNameTH,
			&patient.MiddleNameTH,
			&patient.LastNameTH,
			&patient.FirstNameEN,
			&patient.MiddleNameEN,
			&patient.LastNameEN,
			&patient.DateOfBirth,
			&patient.PatientHN,
			&patient.NationalID,
			&patient.PassportID,
			&patient.PhoneNumber,
			&patient.Email,
			&patient.Gender,
			&patient.CreatedAt,
			&patient.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan patient: %w", err)
		}
		patients = append(patients, patient)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate patients: %w", err)
	}
	return &patients, nil
}
