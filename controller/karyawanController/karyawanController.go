package karyawancontroller

import (
	"github.com/agustiawanilham/go-api-karyawan/models"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var karyawans []models.Karyawan
	models.DB.Find(&karyawans)

	return c.Status(fiber.StatusOK).JSON(karyawans)
}

func Show(c *fiber.Ctx) error {
	return nil
}

func Create(c *fiber.Ctx) error {
	var karyawan models.Karyawan
	if err := c.BodyParser(&karyawan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := models.DB.Create(karyawan).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return nil
}

func Update(c *fiber.Ctx) error {
	return nil
}

func Delete(c *fiber.Ctx) error {
	return nil
}
