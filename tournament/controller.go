package tournament

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/moosashah/go-crud/initializers"
	"github.com/moosashah/go-crud/models"
	"gorm.io/gorm"
)

func Create(c *fiber.Ctx) error {
	var payload *models.CreateTournamentSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	newTournament := models.Tournament{
		Name:            payload.Name,
		AddressLineOne:  payload.AddressLineOne,
		AddressLineTwo:  payload.AddressLineTwo,
		AddressCity:     payload.AddressCity,
		AddressPostCode: payload.AddressPostCode,
	}

	result := initializers.DB.Create(&newTournament)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Name already exists, please use another name"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": newTournament})
}

func FindTournaments(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var tournaments []models.Tournament
	results := initializers.DB.Limit(intLimit).Offset(offset).Find(&tournaments)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(tournaments), "tournaments": tournaments})
}

func UpdateTournament(c *fiber.Ctx) error {
	tournamentId := c.Params("tournamentId")
	var payload *models.UpdateTournamentSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var tournament models.Tournament
	result := initializers.DB.First(&tournament, "id = ?", tournamentId).Updates(&payload)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Tournament with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"tournament": tournament}})
}

func FindTournamentById(c *fiber.Ctx) error {
	tournamentId := c.Params("tournamentId")

	var tournament models.Tournament
	result := initializers.DB.First(&tournament, "id = ?", tournamentId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Tournament with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"note": tournament}})
}

func DeleteTournament(c *fiber.Ctx) error {
	tournamentId := c.Params("tournamentId")
	var tournament models.Tournament
	result := initializers.DB.Delete(&tournament, "id = ?", tournamentId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No tournament with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": result.Error})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
