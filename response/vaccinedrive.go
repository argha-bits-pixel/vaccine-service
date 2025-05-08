package response

import (
	"fmt"
	"log"
	"time"
	"vaccination-service/models"
	"vaccination-service/request"
	"vaccination-service/utils/validator"

	"github.com/labstack/echo/v4"
)

type VaccineDriveResponseHandler interface {
	ProcessVaccineDriveResponse(req interface{}, data interface{}) interface{}
	ProcessErrorResponse(err error) interface{}
}

type VaccineDriveResposne struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error,omitempty"`
	Links   interface{} `json:"links,omitempty"`
}
type VaccineDriveGetResponse struct {
	Id        int         `json:"id"`
	Vaccine   string      `json:"vaccine_name"`
	DriveDate string      `json:"drive_date"`
	Doses     int         `json:"doses"`
	Classes   string      `json:"classes"`
	CreatedAt string      `json:"created_at"`
	UpdatedAt string      `json:"updated_at,omitempty"`
	Links     interface{} `json:"_links,omitempty"`
}

func (r VaccineDriveResposne) ProcessErrorResponse(err error) interface{} {
	resp := VaccineDriveResposne{}
	switch v := err.(type) {
	case *validator.ValidationError:
		log.Println("error type", v)
		resp.Message = "Invalid Input"
		resp.Data = []string{}
		resp.Error = err.(*validator.ValidationError).Fields
	case *echo.HTTPError:
		log.Println("error type", v)
		if err.(*echo.HTTPError).Code == 415 {
			resp.Message = "Invalid Request"
			resp.Data = []string{}
			resp.Error = map[string]string{
				"error": "Unsupported Media Type. Please use application/json in request header Content-Type",
			}
		}

	default:
		log.Println("error type", v)
		resp.Message = "unable to schedule vaccination drive"
		resp.Data = []string{}
		resp.Error = err.Error()
	}
	return resp
}

func (resp VaccineDriveResposne) ProcessVaccineDriveResponse(req, data interface{}) interface{} {
	r := VaccineDriveResposne{}
	switch req.(type) {
	case *request.VaccineDriveCreateRequest:
		r.Message = "Drive Scheduled Successfully"
		r.Data = data
		r.Links = map[string]interface{}{
			"self": map[string]string{
				"href":   fmt.Sprintf("http://localhost:8080/vaccine/drives/%d", data.(*models.VaccineDrive).Id),
				"method": "GET",
			},
			"edit": map[string]string{
				"href":   fmt.Sprintf("http://localhost:8080/vaccine/drives/%d", data.(*models.VaccineDrive).Id),
				"method": "PATCH",
			},
		}
	case *request.GetVaccineDriveRequest, *request.VaccineDriveUpdateRequest:
		if len(data.([]models.VaccineDrive)) < 1 {
			r.Message = "No Upcoming Drives in 30 days"
			r.Data = []string{}
		} else {
			r.Message = "Drive Information fetched Succesully"
			collectionData := []VaccineDriveGetResponse{}
			for _, j := range data.([]models.VaccineDrive) {
				vaccineDriveResponse := VaccineDriveGetResponse{}
				vaccineDriveResponse.Id = j.Id
				vaccineDriveResponse.Vaccine = j.VaccineName
				vaccineDriveResponse.DriveDate = j.DriveDate.Format("2006-01-02")
				vaccineDriveResponse.Doses = j.Doses
				vaccineDriveResponse.Classes = j.Classes
				vaccineDriveResponse.CreatedAt = j.CreatedAt.Format("2006-01-02 15:04:05")
				vaccineDriveResponse.UpdatedAt = j.UpdatedAt.Format("2006-01-02 15:04:05")
				vaccineDriveResponse.Links = geneRateHateOasForVaccination(j)
				collectionData = append(collectionData, vaccineDriveResponse)
			}
			r.Data = collectionData
			if len(collectionData) == 1 {
				r.Data = collectionData[0]
			}
		}

	}
	return r
}
func geneRateHateOasForVaccination(data models.VaccineDrive) interface{} {
	hateOas := map[string]interface{}{}
	hateOas["self"] = map[string]string{
		"href":   fmt.Sprintf("http://localhost:8080/vaccine/drives/%d", data.Id),
		"method": "GET",
	}
	if time.Until(data.DriveDate) > 0 {
		hateOas["edit"] = map[string]string{
			"href":   fmt.Sprintf("http://localhost:8080/vaccine/drives/%d", data.Id),
			"method": "PATCH",
		}
	}
	return hateOas
}

func NewVaccineDriveResponseHandler() VaccineDriveResponseHandler {
	return VaccineDriveResposne{}
}
