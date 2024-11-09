package repository

import (
	"database/sql"

	"medical-portal/internal/models"
)

type PatientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) CreatePatient(patient *models.Patient) error {
	_, err := r.db.Exec("INSERT INTO patients (name, age, condition) VALUES ($1, $2, $3)",
		patient.Name, patient.Age, patient.Condition)
	return err
}

func (r *PatientRepository) UpdatePatient(patient *models.Patient) error {
	_, err := r.db.Exec("UPDATE patients SET name = $1, age = $2, condition = $3 WHERE id = $4",
		patient.Name, patient.Age, patient.Condition, patient.ID)
	return err
}

func (r *PatientRepository) DeletePatient(id string) error {
	_, err := r.db.Exec("DELETE FROM patients WHERE id = $1", id)
	return err
}

func (r *PatientRepository) GetAllPatients() ([]models.Patient, error) {
	rows, err := r.db.Query("SELECT id, name, age, condition FROM patients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []models.Patient
	for rows.Next() {
		var p models.Patient
		if err := rows.Scan(&p.ID, &p.Name, &p.Age, &p.Condition); err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}
	return patients, nil
}