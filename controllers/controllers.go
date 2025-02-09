package controllers

import (
	"net/http"
	"strconv"

	"github.com/tr1sm0s1n/fiber-postgres-api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateOne(c *fiber.Ctx, db *gorm.DB) error {
	var newCertificate models.Certificate
	if err := c.BodyParser(&newCertificate); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	result := db.Create(&newCertificate)
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).SendString(result.Error.Error())
	}

	return c.Status(http.StatusCreated).JSON(newCertificate)
}

func ReadAll(c *fiber.Ctx, db *gorm.DB) error {
	var certificates []models.Certificate
	result := db.Find(&certificates)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(certificates)
}

func ReadOne(c *fiber.Ctx, db *gorm.DB) error {
	var oldCertificate models.Certificate
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	result := db.First(&oldCertificate, "id = ?", id)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}

	return c.Status(http.StatusOK).JSON(oldCertificate)
}

func UpdateOne(c *fiber.Ctx, db *gorm.DB) error {
	var update models.Certificate
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	result := db.First(&update, "id = ?", id)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}

	if err := c.BodyParser(&update); err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	result = db.Save(&update)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(update)
}

func DeleteOne(c *fiber.Ctx, db *gorm.DB) error {
	var trash models.Certificate
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	result := db.First(&trash, "id = ?", id)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	}

	result = db.Delete(&models.Certificate{}, id)
	if result.Error != nil {
		return c.Status(http.StatusInternalServerError).SendString(result.Error.Error())
	}

	return c.Status(http.StatusOK).JSON(trash)
}
