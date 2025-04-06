package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Name             string `gorm:"size:100;not null"`
	Email            string `gorm:"size:100;uniqueIndex"`
	Phone            string `gorm:"size:20"`
	SpecialisationID uint
	Password         string
	Specialisation   Specialisation `gorm:"foreignKey:SpecialisationID"`
	Appointments     []Appointment  `gorm:"foreignKey:DoctorID"`
	YearsExperience  int            `gorm:"default:0"`
	IsAvailable      bool           `gorm:"default:true"`
}
