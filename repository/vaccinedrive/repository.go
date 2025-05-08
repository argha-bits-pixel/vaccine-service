package vaccinedrive

import (
	"fmt"
	"log"
	"vaccination-service/adapters/mysql"
	"vaccination-service/models"
	"vaccination-service/repository"
	"vaccination-service/request"
)

type VaccineRepository struct {
	DB *mysql.MysqlConnect
}

func (v *VaccineRepository) GetVaccineDrive(filter string) ([]models.VaccineDrive, error) {
	var drives []models.VaccineDrive
	var err error
	if filter == "" {
		err = v.DB.Table("vaccination_drives").Order("drive_date ASC").Find(&drives).Error
		if err != nil {
			log.Println("error in fetching drives", err.Error())
			return drives, err
		}
		return drives, nil
	}
	fmt.Println("Filter being applied", filter)
	err = v.DB.Table("vaccination_drives").Where(filter).Order("drive_date ASC").Find(&drives).Error
	if err != nil {
		log.Println("error in fetching drives", err.Error())
		return drives, err
	}
	return drives, nil
}
func (v *VaccineRepository) CreateDrive(drive *models.VaccineDrive) error {
	return v.DB.Table("vaccination_drives").Create(drive).Error
}
func (v *VaccineRepository) UpdateVaccineDrive(drive *request.VaccineDriveUpdateRequest) error {
	updateMap := map[string]interface{}{}

	if drive.DriveDate != nil {
		updateMap["drive_date"] = drive.DriveDate
	}
	if drive.VaccineName != nil {
		updateMap["vaccine_name"] = drive.VaccineName
	}
	if drive.Doses != nil {
		updateMap["doses"] = drive.Doses
	}
	if drive.Classes != nil {
		updateMap["classes"] = drive.Classes
	}
	return v.DB.Table("vaccination_drives").Where(fmt.Sprintf("id = %d", drive.Id)).Updates(updateMap).Error
}

func NewVaccineRepositoryHandler(db *mysql.MysqlConnect) repository.VaccineRepositoryHandler {
	return &VaccineRepository{
		DB: db,
	}
}
