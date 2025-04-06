package db

import (
	"fmt"
	"log"
	"os"

	"github.com/dhanushs3366/patient-app/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func New() (*Store, error) {
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")
	DB_HOST := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Error opening db, %s\n", err.Error())
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error accesing DB %s\n", err.Error())
		return nil, err
	}

	err = sqlDB.Ping()

	if err != nil {
		log.Printf("Error pinging db %s\n", err.Error())
		return nil, err
	}

	return &Store{DB: db}, nil
}

func (s *Store) Close() error {
	sqlDB, err := s.DB.DB()
	if err != nil {
		log.Printf("Error closing db, %s\n", err.Error())
		return err
	}
	return sqlDB.Close()
}

func (s *Store) Init() error {
	err := s.DB.AutoMigrate(
		&models.Doctor{},
		&models.Appointment{},
		&models.Patient{},
		&models.Specialisation{},
	)

	if err != nil {
		log.Printf("failed to migrate models, %s\n", err.Error())
		return err
	}
	return nil
}
