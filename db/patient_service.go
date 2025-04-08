package db

import (
	"log"

	"github.com/dhanushs3366/patient-app/db/models"
)

func (s *Store) CreateNewPatient(patient models.Patient) (uint, error) {
	result := s.DB.Create(&patient)

	if result.Error != nil {
		log.Printf("failed to insert new patient, %s\n", result.Error.Error())

		return 0, result.Error
	}

	return patient.Model.ID, nil
}

func (s *Store) GetPatientByEmail(email string) (*models.Patient, error) {
	var patient models.Patient
	result := s.DB.Where("email = ?", email).First(&patient)

	if result.Error != nil {
		log.Printf("no record found, %s\n", result.Error.Error())
		return nil, result.Error
	}

	return &patient, nil
}
