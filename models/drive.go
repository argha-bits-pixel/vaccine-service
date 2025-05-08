package models

import "time"

type VaccineDrive struct {
	Id          int       `json:"id"`
	VaccineName string    `json:"vaccine_name"`
	DriveDate   time.Time `json:"drive_date"`
	Doses       int       `json:"doses"`
	Classes     string    `gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
