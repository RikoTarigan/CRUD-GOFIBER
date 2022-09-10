package main

import (
	"CURD_KARYAWAN/database"
	"CURD_KARYAWAN/database/migration"
	"CURD_KARYAWAN/route"

	"github.com/gofiber/fiber/v2"
)

func main(){
	//INIT DATABASE
	database.DatabaseInit()

	//ORM
	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	route.RoutePath(app)

	app.Listen(":3600")
}