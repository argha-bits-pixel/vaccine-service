package request

import (
	"log"
	"time"
	"vaccination-service/models"

	"github.com/labstack/echo/v4"
)

type VaccineDriveRequestHandler interface {
	Bind(c echo.Context, request interface{}, model *models.VaccineDrive) error
}

type VaccineDriveRequest struct {
}

type GetVaccineDriveRequest struct {
	Id     int    `param:"id"`
	Name   string `query:"vaccine_name"`
	Limit  int    `query:"limit"`
	Offset int    `query:"offset"`
}

type VaccineDriveCreateRequest struct {
	DriveDate   time.Time `json:"drive_date" validate:"checkValidDriveDate"`
	VaccineName string    `json:"vaccine_name" validate:"required"`
	Doses       int       `json:"doses" validate:"required"`
	Classes     string    `json:"classes" validate:"required"`
}

type VaccineDriveUpdateRequest struct {
	Id          int        `json:"id" validate:"required"`
	DriveDate   *time.Time `json:"drive_date,omitempty"`
	VaccineName *string    `json:"vaccine_name,omitempty"`
	Doses       *int       `json:"doses,omitempty"`
	Classes     *string    `json:"classes,omitempty"`
}

func (r VaccineDriveRequest) Bind(c echo.Context, req interface{}, model *models.VaccineDrive) error {
	var err error

	if err = c.Bind(req); err != nil {
		log.Println("Error in reading request", err.Error())
		return err
	}
	if err = c.Validate(req); err != nil {
		log.Println("error in validating request", err.Error())
		return err
	}
	switch v := req.(type) {
	case *VaccineDriveCreateRequest:
		model.DriveDate = req.(*VaccineDriveCreateRequest).DriveDate.UTC()
		model.VaccineName = req.(*VaccineDriveCreateRequest).VaccineName
		model.Doses = req.(*VaccineDriveCreateRequest).Doses
		model.Classes = req.(*VaccineDriveCreateRequest).Classes
	case *GetVaccineDriveRequest:
		model.Id = req.(*GetVaccineDriveRequest).Id
		model.VaccineName = req.(*GetVaccineDriveRequest).Name
		log.Println("req is ", model)
	case **VaccineDriveUpdateRequest:
		model.Id = req.(*VaccineDriveUpdateRequest).Id
	default:
		log.Println("request type Unknown for transformation", v)
	}

	return nil
}
func NewVaccineDriveRequestHandler() VaccineDriveRequestHandler {
	return VaccineDriveRequest{}
}
