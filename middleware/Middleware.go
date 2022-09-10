package middleware

import (
	"CURD_KARYAWAN/utills"
	"log"

	"github.com/gofiber/fiber/v2"
)

func MiddlewareInit(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unAuthenticated",
		})
	}


	_ , err := utills.VerifyJwt(token)
	if err != nil{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unAuthenticated",
		})
	}

	claims,err := utills.DecodeToken(token)
	if err != nil{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unAuthenticated",
		})
	}
	
	role := claims["role"].(string)
	if role == ""{
		role = "2" //default user if role is blank
	}

	log.Println(claims)
	ctx.Locals("karyawanInfo",claims)


	// if token == "" || token != "secret" {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "unauthorize",
	// 	})
	// }
	return ctx.Next()
}