package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	PatientID   uint
	Patient     Patient `gorm:"foreignKey:PatientID"`
	DoctorID    uint
	Doctor      Doctor    `gorm:"foreignKey:DoctorID"`
	AllotedTime time.Time `gorm:"index;not null"`
	EndTime     time.Time
	Status      string `gorm:"size:20;default:'scheduled'"` // scheduled, completed, cancelled
	Notes       string `gorm:"type:text"`
	RoomNumber  string `gorm:"size:20"`
	IsFollowUp  bool   `gorm:"default:false"`
}
