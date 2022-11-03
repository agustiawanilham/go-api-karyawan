// Package handlers ...
package handlers

import (
	"github.com/agustiawanilham/go-api-karyawan/db"
	"github.com/agustiawanilham/go-api-karyawan/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAll Get all data karyawan with pagination
func GetAll(c *fiber.Ctx) error {
	var karyawans []models.Karyawan
	db.DB.Find(&karyawans)

	return c.Status(fiber.StatusOK).JSON(karyawans)
}

// GetData Get data karyawan by ID
func GetData(c *fiber.Ctx) error {
	paramsID := c.Params("id")
	var karyawan models.Karyawan

	if err := db.DB.First(&karyawan, paramsID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Data not found",
			})

		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data not found",
		})

	}

	return c.JSON(karyawan)
}

// Create Insert new data karyawan
func Create(c *fiber.Ctx) error {
	var karyawan models.Karyawan
	if err := c.BodyParser(&karyawan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := db.DB.Create(&karyawan).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(karyawan)
}

// Update Update data karyawan
func Update(c *fiber.Ctx) error {
	paramsID := c.Params("id")
	var karyawan models.Karyawan

	if err := c.BodyParser(&karyawan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if db.DB.Where("id = ? ", paramsID).Updates(&karyawan).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cannot update data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Update success",
	})
}

// Delete data karyawan
func Delete(c *fiber.Ctx) error {
	paramsID := c.Params("id")
	var karyawan models.Karyawan

	if db.DB.Delete(&karyawan, paramsID).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cannot delete data",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Delete success",
	})
}
