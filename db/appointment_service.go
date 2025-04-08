package db

import (
	"errors"
	"time"

	"github.com/dhanushs3366/patient-app/db/models"
)

func (s *Store) BookAppointment(patientID, doctorID uint) error {

	var patient models.Patient
	if err := s.DB.First(&patient, patientID).Error; err != nil {
		return errors.New("patient not found")
	}

	var doctor models.Doctor
	if err := s.DB.First(&doctor, doctorID).Error; err != nil {
		return errors.New("doctor not found")
	}

	if !doctor.IsAvailable {
		return errors.New("doctor is not available for appointments")
	}

	// Set default appointment time (next available slot - current time + 1 hour)
	startTime := time.Now().Add(time.Hour)
	// Default appointment duration (30 minutes)
	endTime := startTime.Add(30 * time.Minute)

	// Check for conflicting appointments
	var conflictCount int64
	s.DB.Model(&models.Appointment{}).
		Where("doctor_id = ? AND status = 'scheduled' AND alloted_time < ? AND end_time > ?",
			doctorID, endTime, startTime).
		Count(&conflictCount)

	if conflictCount > 0 {
		return errors.New("doctor already has an appointment during this time slot")
	}

	appointment := models.Appointment{
		PatientID:   patientID,
		DoctorID:    doctorID,
		AllotedTime: startTime,
		EndTime:     endTime,
		Status:      "scheduled",
		RoomNumber:  "TBD",
	}

	if err := s.DB.Create(&appointment).Error; err != nil {
		return err
	}

	return nil
}
