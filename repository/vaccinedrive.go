package repository

import (
	"vaccination-service/models"
	"vaccination-service/request"
)

type VaccineRepositoryHandler interface {
	GetVaccineDrive(filter string) ([]models.VaccineDrive, error)
	CreateDrive(drive *models.VaccineDrive) error
	UpdateVaccineDrive(drive *request.VaccineDriveUpdateRequest) error
}
