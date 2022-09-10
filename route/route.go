package route

import (
	"CURD_KARYAWAN/handler"
	"CURD_KARYAWAN/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)


func RoutePath(r *fiber.App){

	//CORS Setting

	// Or extend your config for customization
	r.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET,POST,DELETE,PATCH",
	}))

	// NO MIDDLEWARE
	r.Post("/create",handler.KaryawanHandlerCreate) //register karyawan
	r.Post("/login",handler.AuthHandlerLoginKaryawan) //login karyawan

	//WITH MIDDLEWARE
	r.Get("/Read",middleware.MiddlewareInit,handler.KaryawanHandlerGetAll) //get all data karyawan
	r.Get("/Read/:id",middleware.MiddlewareInit,handler.KaryawanHandlerGetById) //get karyawan by id
	r.Patch("/update/:id",middleware.MiddlewareInit,handler.KaryawanHandlerUpdateById) //update karyawan //only for admin
	r.Delete("/delete/:id",middleware.MiddlewareInit,handler.KaryawanHandlerDeleteById) //delete karyawan //only for admin
}