package vaccinedrive

import (
	"errors"
	"fmt"
	"log"
	"time"
	"vaccination-service/models"
	"vaccination-service/repository"
	"vaccination-service/request"
	"vaccination-service/usecase"
)

type VaccineDriveUsecase struct {
	repo repository.VaccineRepositoryHandler
}

func (v *VaccineDriveUsecase) GetVaccineDriveDetails(drive *models.VaccineDrive) ([]models.VaccineDrive, error) {
	return v.repo.GetVaccineDrive(createFilterstring(drive))
}
func (v *VaccineDriveUsecase) CreatevaccineDrive(drive *models.VaccineDrive) error {
	//check if any vaccine drive is scheduled in that day
	log.Println("drive being created", drive)
	data, err := v.repo.GetVaccineDrive(fmt.Sprintf(`drive_date = "%s"`, drive.DriveDate.Format("2006-01-02")))
	if err != nil {
		log.Println("Unable to schedule vaccination drive ", err.Error())
		return errors.New("unable to schedule drive please try again later")
	}
	log.Println("data being fetched", data)
	if len(data) > 0 {
		return fmt.Errorf("vaccination drive exists on %s, drive id: %d", data[0].DriveDate.Format("2006-01-02"), data[0].Id)
	}
	//schedule drive
	return v.repo.CreateDrive(drive)

}

func (v *VaccineDriveUsecase) EditVaccineDrive(drive *request.VaccineDriveUpdateRequest) error {
	var err error
	//check if a drive with same id already exists or not
	driveDetails, err := v.repo.GetVaccineDrive(fmt.Sprintf("id = %d", drive.Id))
	if err != nil {
		log.Println("error fetching drive details", err.Error())
		return err
	}
	if len(driveDetails) == 0 {
		return fmt.Errorf("no drive exists with id %d ", drive.Id)
	}
	//check if already any drive is scheduled on the new date
	if drive.DriveDate != nil {
		data, err := v.repo.GetVaccineDrive(fmt.Sprintf(`drive_date = "%s"`, drive.DriveDate.Format("2006-01-02")))
		if err != nil {
			log.Println("Unable to schedule vaccination drive ", err.Error())
			return errors.New("unable to schedule drive please try again later")
		}
		log.Println("data being fetched", data)
		if len(data) > 0 && data[0].Id != driveDetails[0].Id {
			return fmt.Errorf("vaccination drive exists on %s, drive id: %d", data[0].DriveDate.Format("2006-01-02"), data[0].Id)
		}
		//check if they are editing a completed drive
		if time.Now().After(driveDetails[0].DriveDate) {
			return fmt.Errorf("drive with id %d already completed on %s", driveDetails[0].Id, driveDetails[0].DriveDate)
		}
		//check if they are prescheduling the drive
		if !drive.DriveDate.After(driveDetails[0].DriveDate) {
			return fmt.Errorf("drive prescheduling is not possible, please schedule after %s", driveDetails[0].DriveDate)
		}
	}
	return v.repo.UpdateVaccineDrive(drive)
}
func createFilterstring(drive *models.VaccineDrive) string {
	filterString := ""
	if drive.Id != 0 {
		filterString += fmt.Sprintf("id = %d", drive.Id)
	} else if drive.VaccineName != "" {
		filterString += fmt.Sprintf("vaccine_name LIKE '%s'", drive.VaccineName)
	} else {
		filterString += "drive_date <= DATE_ADD(NOW(), INTERVAL 30 DAY)"
	}
	return filterString
}

func NewVaccineRepositoryHandler(repo repository.VaccineRepositoryHandler) usecase.VaccineDriveUsecasehandler {
	return &VaccineDriveUsecase{
		repo: repo,
	}
}
