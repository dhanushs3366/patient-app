package db

import (
	"github.com/dhanushs3366/patient-app/db/models"
)

func (s *Store) GetAllSpecialisationTypes() ([]string, error) {
	var specialisations []string
	result := s.DB.Model(&models.Specialisation{}).Select("name").Scan(&specialisations)

	if result.Error != nil {
		return nil, result.Error
	}

	return specialisations, nil
}
