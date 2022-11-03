package routes

import (
	authHandler "github.com/agustiawanilham/go-api-karyawan/handlers"
	karyawanHandler "github.com/agustiawanilham/go-api-karyawan/handlers"
	"github.com/agustiawanilham/go-api-karyawan/middleware"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

// SetupRoutes setup all routes
func SetupRoutes(app *fiber.App) {
	// Grouping by api
	api := app.Group("/api/v1")

	// Setup swagger
	route := api.Group("/swagger")
	route.Get("*", swagger.HandlerDefault)
	route.Post("*", swagger.HandlerDefault)
	route.Patch("*", swagger.HandlerDefault)
	route.Delete("*", swagger.HandlerDefault)

	// Auth endpoints
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.LoginUser)
	auth.Post("/register", authHandler.RegisterUser)

	// Karyawan protected with JWT middleware endpoints
	karyawan := api.Group("/karyawan")
	karyawan.Use(middleware.JWTProtected())
	karyawan.Get("/", karyawanHandler.GetAll)
	karyawan.Get("/:id", karyawanHandler.GetData)
	karyawan.Post("/", karyawanHandler.Create)
	karyawan.Patch("/:id", karyawanHandler.Update)
	karyawan.Delete("/:id", karyawanHandler.Delete)
}
