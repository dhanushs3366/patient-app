package db

import "github.com/dhanushs3366/patient-app/db/models"

// add super admin role such that only admin can add doctor

func (s *Store) InsertDoctor(doctor *models.Doctor) error {
	result := s.DB.Create(doctor)
	return result.Error
}

func (s *Store) UpdateDoctor(doctor *models.Doctor) error {
	result := s.DB.Save(doctor)
	return result.Error
}
func (s *Store) GetDoctorByID(id uint) (*models.Doctor, error) {
	var doctor models.Doctor
	result := s.DB.Preload("Specialisation").First(&doctor, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &doctor, nil
}

func (s *Store) GetDoctorBySpecialisation(specialisationID uint) ([]models.Doctor, error) {
	var doctors []models.Doctor
	result := s.DB.Preload("Specialisation").Where("specialisation_id = ?", specialisationID).Find(&doctors)
	if result.Error != nil {
		return nil, result.Error
	}
	return doctors, nil
}

func (s *Store) GetAllDoctors() ([]models.Doctor, error) {
	var doctors []models.Doctor
	result := s.DB.Preload("Specialisation").Find(&doctors)
	if result.Error != nil {
		return nil, result.Error
	}
	return doctors, nil
}

func (s *Store) GetDoctorByMail(username string) (*models.Doctor, error) {
	var doctor models.Doctor
	result := s.DB.Preload("Specialisation").Where("email = ?", username).First(&doctor)
	if result.Error != nil {
		return nil, result.Error
	}
	return &doctor, nil
}
