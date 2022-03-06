package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/Practicum-1/lawyer-client-backend.git/helpers"
	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"github.com/Practicum-1/lawyer-client-backend.git/repositories"
)

func GetAllLawyer(c *fiber.Ctx) error {

	lawyer, err := repositories.GetAllLawyers()
	if err != nil {
		if err.Error() == "404" {
			return helpers.SendResponse(c, fiber.StatusNotFound, "No lawyer found", nil)
		} else {
			return helpers.SendResponse(c, fiber.StatusBadRequest, err.Error(), err)
		}
	}
	return helpers.SendResponse(c, fiber.StatusOK, "success", lawyer)
}

func GetLawyerById(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusBadRequest, "Invalid Id", err)
	}

	var lawyer models.Lawyer
	err = repositories.GetLawyerById(id, &lawyer)
	if err != nil {
		if err.Error() == "404" {
			return helpers.SendResponse(c, fiber.StatusNotFound, "Lawyer not found", nil)
		} else {
			return helpers.SendResponse(c, fiber.StatusBadRequest, err.Error(), err)
		}
	}
	return helpers.SendResponse(c, fiber.StatusOK, "success", lawyer)

}

func CreateLawyer(c *fiber.Ctx) error {
	newLawyer := &models.Lawyer{}

	//Parse the body
	if err := c.BodyParser(newLawyer); err != nil {
		return helpers.SendResponse(c, fiber.StatusBadRequest, "Invalid input", err)
	}
	fmt.Println("newLawyer", newLawyer)
	//Create Client
	err := repositories.CreateLawyer(newLawyer)
	fmt.Println("Error: ", err)
	if err != nil {
		return helpers.SendResponse(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return helpers.SendResponse(c, fiber.StatusCreated, "Lawyer Created Successfully", nil)
}

func GetLawyerByFilter(c *fiber.Ctx) error {
	//Parse the body
	//get query
	location_id, _ := strconv.ParseUint(c.Query("location_id"), 10, 64)
	gender := c.Query("gender")
	experience, _ := strconv.ParseUint(c.Query("experience"), 10, 64)
	language_id, _ := strconv.ParseUint(c.Query("language_id"), 10, 64)
	practice_area_id, _ := strconv.ParseUint(c.Query("practice_area_id"), 10, 64)
	court_id, _ := strconv.ParseUint(c.Query("court_id"), 10, 64)

	//fiber map
	var filters models.Filters
	filters.LocationID = uint(location_id)
	filters.Gender = gender
	filters.Experience = uint(experience)
	filters.LanguageID = uint(language_id)
	filters.PracticeAreaID = uint(practice_area_id)
	filters.CourtID = uint(court_id)

	lawyers, _ := repositories.GetLawyerByFilter(filters)

	// if err != nil {
	// 	if err.Error() == "404" {
	// 		return helpers.SendResponse(c, fiber.StatusNotFound, "Lawyer not found", nil)
	// 	} else {
	// 		return helpers.SendResponse(c, fiber.StatusBadRequest, err.Error(), err)
	// 	}
	// }
	return helpers.SendResponse(c, fiber.StatusOK, "success", lawyers)
}
