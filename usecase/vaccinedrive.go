package usecase

import (
	"vaccination-service/models"
	"vaccination-service/request"
)

type VaccineDriveUsecasehandler interface {
	GetVaccineDriveDetails(drive *models.VaccineDrive) ([]models.VaccineDrive, error)
	CreatevaccineDrive(drive *models.VaccineDrive) error
	EditVaccineDrive(drive *request.VaccineDriveUpdateRequest) error
}
