package handler

import (
	"CURD_KARYAWAN/database"
	"CURD_KARYAWAN/model"
	"CURD_KARYAWAN/utills"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AuthHandlerLoginKaryawan(ctx *fiber.Ctx) error {
	loginRequest := new(model.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	//Validation
	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil{
		return ctx.Status(400).JSON(fiber.Map{
			"message":"failed",
			"erorr":errValidate.Error(),
		})
	}

	var karyawan model.Karyawan
	err := database.DB.First(&karyawan , "email = ?", loginRequest.Email).Error
	if err != nil{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"Wrong credential",
		})
	}

	//password checker
	isPasswordValid := utills.PasswordUtilsDecodeHashing(loginRequest.Password,karyawan.Password)
	if !isPasswordValid{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"Wrong credential",
		})
	}

	//generate token jwt
	claims := jwt.MapClaims{}
	claims["name"] = karyawan.Name
	claims["email"] = karyawan.Email
	claims["address"] = karyawan.Address
	claims["exp_time"] = time.Now().Add(time.Minute * 2).Unix() //set expirateion time
	claims["role"] = karyawan.Role
	
	token,errGenerate := utills.GenerateJwtUtils(&claims)
	if errGenerate != nil{
		log.Println(errGenerate)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"Wrong credential",
		})
	}
	
	return ctx.JSON(fiber.Map{
		"token":token,
	})
}