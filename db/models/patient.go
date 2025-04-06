package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	Name           string `gorm:"size:100;not null"`
	Email          string `gorm:"size:100;uniqueIndex"`
	Phone          string `gorm:"size:20"`
	DateOfBirth    time.Time
	Address        string        `gorm:"type:text"`
	MedicalHistory string        `gorm:"type:text"`
	Appointments   []Appointment `gorm:"foreignKey:PatientID"`
}
