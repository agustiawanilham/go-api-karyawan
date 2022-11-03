// main
package main

import (
	"log"

	karyawancontroller "github.com/agustiawanilham/go-api-karyawan/controller/karyawanController"
	"github.com/agustiawanilham/go-api-karyawan/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()
	api := app.Group("/api")
	karyawan := api.Group("/karyawan")

	karyawan.Get("/", karyawancontroller.Index)
	karyawan.Get("/:id", karyawancontroller.Show)
	karyawan.Post("/", karyawancontroller.Create)
	karyawan.Put("/:id", karyawancontroller.Update)
	karyawan.Delete("/:id", karyawancontroller.Delete)

	log.Fatal(app.Listen(":3000"))
}
