package handlers

import (
	"github.com/agustiawanilham/go-api-karyawan/db"
	"github.com/agustiawanilham/go-api-karyawan/middleware"
	"github.com/agustiawanilham/go-api-karyawan/models"
	"github.com/gofiber/fiber/v2"
)

// TokenRequest struct
type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginUser             godoc
// @Summary      Login to account user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login {object} body TokenRequest true "User login"
// @Success      200  {object}  TokenRequest
// @Router       /api/v1/auth/login [post]
func LoginUser(c *fiber.Ctx) error {
	var request TokenRequest
	var user models.User

	// parse body json
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// check username
	if err := db.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// check password
	credentialPasswordError := user.CheckPassword(request.Password)
	if credentialPasswordError != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid credential",
		})
	}

	// generate token JWT
	token, err := middleware.GenerateJWT(user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid credential",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

// RegisterUser             godoc
//
// @Summary      Register account user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user body models.User  true  "User Data"
// @Success      200  {object}  models.User
// @Router       /api/v1/auth/register [post]
func RegisterUser(c *fiber.Ctx) error {
	var user models.User

	// parse body json
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := user.HashPassword(user.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
