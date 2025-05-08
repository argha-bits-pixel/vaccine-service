package vaccinationservice

import (
	"log"
	"net/http"
	"vaccination-service/controller"
	"vaccination-service/models"
	"vaccination-service/request"
	"vaccination-service/response"
	"vaccination-service/usecase"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	req  request.VaccineDriveRequestHandler
	uc   usecase.VaccineDriveUsecasehandler
	resp response.VaccineDriveResponseHandler
}

func (v Controller) GetVaccinationDriveDetails(c echo.Context) error {
	var err error
	req := new(request.GetVaccineDriveRequest)
	model := new(models.VaccineDrive)
	if err = v.req.Bind(c, req, model); err != nil {
		log.Println("error in binding request")
		return c.JSON(http.StatusBadRequest, v.resp.ProcessErrorResponse(err))
	}
	var data []models.VaccineDrive
	if data, err = v.uc.GetVaccineDriveDetails(model); err != nil {
		log.Println("error in creating drive", err.Error())
		return c.JSON(http.StatusBadRequest, v.resp.ProcessErrorResponse(err))
	}
	return c.JSON(http.StatusOK, v.resp.ProcessVaccineDriveResponse(req, data))
}

func (v Controller) CreateVaccinationDrive(c echo.Context) error {
	var err error
	req := new(request.VaccineDriveCreateRequest)
	model := new(models.VaccineDrive)
	if err = v.req.Bind(c, req, model); err != nil {
		log.Println("error in binding request")
		return c.JSON(http.StatusBadRequest, v.resp.ProcessErrorResponse(err))
	}
	if err = v.uc.CreatevaccineDrive(model); err != nil {
		log.Println("error in creating drive", err.Error())
		return c.JSON(http.StatusBadRequest, v.resp.ProcessErrorResponse(err))
	}
	return c.JSON(http.StatusCreated, v.resp.ProcessVaccineDriveResponse(req, model))
}
func (v Controller) EditVaccinationDrive(c echo.Context) error {
	var err error
	req := new(request.VaccineDriveUpdateRequest)
	model := new(models.VaccineDrive)
	if err = v.req.Bind(c, req, model); err != nil {
		log.Println("error in binding request")
		return c.JSON(http.StatusBadRequest, v.resp.ProcessErrorResponse(err))
	}
	if err = v.uc.EditVaccineDrive(req); err != nil {
		log.Println("error in creating drive", err.Error())
		return c.JSON(http.StatusBadRequest, v.resp.ProcessErrorResponse(err))
	}
	var data []models.VaccineDrive
	model.Id = req.Id
	if data, err = v.uc.GetVaccineDriveDetails(model); err != nil {
		log.Println("error in getting drive", err.Error())
		return c.JSON(http.StatusBadRequest, v.resp.ProcessErrorResponse(err))
	}
	return c.JSON(http.StatusOK, v.resp.ProcessVaccineDriveResponse(req, data))
}

func NewVaccineServiceController(e *echo.Echo, req request.VaccineDriveRequestHandler, uc usecase.VaccineDriveUsecasehandler, resp response.VaccineDriveResponseHandler) controller.VaccineServiceController {
	vaccineServiceController := Controller{
		req:  req,
		uc:   uc,
		resp: resp,
	}
	e.POST("/vaccine/drives", vaccineServiceController.CreateVaccinationDrive)
	e.GET("/vaccine/drives", vaccineServiceController.GetVaccinationDriveDetails)
	e.GET("/vaccine/drives/:id", vaccineServiceController.GetVaccinationDriveDetails)
	e.PATCH("/vaccine/drives", vaccineServiceController.EditVaccinationDrive)
	return e
}
