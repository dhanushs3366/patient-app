package models

import "gorm.io/gorm"

type Specialisation struct {
	gorm.Model
	Name        string   `gorm:"size:100;uniqueIndex;not null"`
	Description string   `gorm:"type:text"`
	Doctors     []Doctor `gorm:"foreignKey:SpecialisationID"`
}
