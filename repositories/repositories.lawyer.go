package repositories

import (
	"errors"
	"fmt"

	"github.com/Practicum-1/lawyer-client-backend.git/db"
	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"gorm.io/gorm"
)

func GetAllLawyers() (interface{}, error) {
	db := db.GetDB()
	var lawyer []models.Lawyer
	result := db.Find(&lawyer)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("404")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return lawyer, nil
}

func GetLawyerById(id uint64, lawyer *models.Lawyer) error {
	db := db.GetDB()
	// result := db.Preload(clause.Associations).Where("id = ?", id).First(&lawyer)
	result := db.Where("id = ?", id).First(&lawyer)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("404")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateLawyer(lawyer *models.Lawyer) error {
	db := db.GetDB()
	if result := db.Create(&lawyer); result.Error != nil {
		return result.Error
	}
	return nil
}

func getIDs(lawyers []models.Lawyer) []uint {
	var resultIDs []uint
	for _, lawyer := range lawyers {
		resultIDs = append(resultIDs, lawyer.ID)
	}
	return resultIDs
}

func GetLawyerByFilter(filters map[string]string) ([]models.Lawyer, error) {
	db := db.GetDB()
	var lawyers []models.Lawyer
	var noOfAssociationFilters int = 0
	var statement string = "SELECT * FROM lawyers WHERE 1=1"
	if filters["experience"] != "" {
		statement += " AND `experience` = " + filters["experience"]
	}
	if filters["location_id"] != "" {
		statement += " AND `location_id` = " + filters["location_id"]
	}
	if filters["gender"] != "" {
		statement += " AND `gender` = \"" + filters["gender"] + "\""
	}
	if filters["practice_area_id"] != "" {
		colName := func() string {
			if noOfAssociationFilters == 0 {
				return " `id`"
			} else {
				return " `lawyer_id`"
			}
		}()
		statement += " AND" + colName + " IN ( SELECT `lawyer_id` FROM `lawyer_practice_areas` WHERE `practice_area_id` = " + filters["practice_area_id"]
		noOfAssociationFilters += 1
	}
	if filters["court_id"] != "" {
		colName := func() string {
			if noOfAssociationFilters == 0 {
				return "`id`"
			} else {
				return "`lawyer_id`"
			}
		}()
		statement += " AND" + colName + " IN ( SELECT `lawyer_id` FROM `lawyer_courts` WHERE `court_id` = " + filters["court_id"]
		noOfAssociationFilters += 1
	}
	if filters["language_id"] != "" {
		colName := func() string {
			if noOfAssociationFilters == 0 {
				return "`id`"
			} else {
				return "`lawyer_id`"
			}
		}()
		statement += " AND" + colName + " IN ( SELECT `lawyer_id` FROM `lawyer_languages` WHERE `language_id` = " + filters["language_id"]
		noOfAssociationFilters += 1
	}
	for i := 0; i < noOfAssociationFilters; i++ {
		statement += ")"
	}
	statement += ";"
	fmt.Println(statement)
	result := db.Raw(statement).Find(&lawyers)
	fmt.Println("RowsAffected", result.RowsAffected)

	return lawyers, nil
}
