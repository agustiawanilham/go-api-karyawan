// Package handlers ...
package handlers

import (
	"fmt"

	"github.com/agustiawanilham/go-api-karyawan/db"
	"github.com/agustiawanilham/go-api-karyawan/helpers"
	"github.com/agustiawanilham/go-api-karyawan/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAll Get all data karyawan with pagination
func GetAll(c *fiber.Ctx) error {
	var karyawans []models.Karyawan
	fetchParam := helpers.FetchParamsPaginationFromRequest(c)
	db.DB.Limit(int(fetchParam.Limit)).Find(&karyawans, "ID > ?", fetchParam.CursorID)

	response := helpers.StandarResponse{
		Data:    &karyawans,
		Message: "success",
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetData Get data karyawan by ID
func GetData(c *fiber.Ctx) error {
	paramsID := c.Params("id")
	var karyawan models.Karyawan

	response := helpers.StandarResponse{
		Data:    nil,
		Message: "success",
	}

	if err := db.DB.First(&karyawan, paramsID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			response.Message = "Data not found"
			return c.Status(fiber.StatusNotFound).JSON(response)
		}

		response.Message = fmt.Sprintf("Failed to Get Data with id : %s", paramsID)
		return c.Status(fiber.StatusInternalServerError).JSON(response)

	}

	response.Data = &karyawan

	return c.Status(fiber.StatusOK).JSON(response)
}

// Create Insert new data karyawan
func Create(c *fiber.Ctx) error {
	var karyawan models.Karyawan

	response := helpers.StandarResponse{
		Data:    nil,
		Message: "Success add",
	}

	if err := c.BodyParser(&karyawan); err != nil {
		response.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if err := db.DB.Create(&karyawan).Error; err != nil {
		response.Message = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Update Update data karyawan
func Update(c *fiber.Ctx) error {
	paramsID := c.Params("id")
	var karyawan models.Karyawan

	response := helpers.StandarResponse{
		Data:    nil,
		Message: "Success data",
	}

	if err := c.BodyParser(&karyawan); err != nil {
		response.Message = err.Error()
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	if db.DB.Where("id = ? ", paramsID).Updates(&karyawan).RowsAffected == 0 {
		response.Message = "Cannot update data"
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Delete data karyawan
func Delete(c *fiber.Ctx) error {
	paramsID := c.Params("id")
	var karyawan models.Karyawan

	response := helpers.StandarResponse{
		Data:    nil,
		Message: "Success add",
	}

	if db.DB.Delete(&karyawan, paramsID).RowsAffected == 0 {
		response.Message = "Failed to delete data"
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
