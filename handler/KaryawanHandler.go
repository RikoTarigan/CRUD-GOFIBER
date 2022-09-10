package handler

import (
	"CURD_KARYAWAN/database"
	"CURD_KARYAWAN/model"
	"CURD_KARYAWAN/utills"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func KaryawanHandlerGetAll(ctx *fiber.Ctx) error {
	var kar []model.Karyawan
	sql := "SELECT * FROM KARYAWANS"

	s := ctx.Query("s")
	if s != "" {
		sql = fmt.Sprintf("%s WHERE NAME LIKE '%%%s%%' OR EMAIL LIKE '%%%s%%'",sql,s,s)
	}
	log.Println(sql)
	

	page,_ := strconv.Atoi(ctx.Query("page","1"))
	perPage := 9
	var totalData int64
	database.DB.Raw(sql).Count(&totalData);

	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d",sql,perPage,(page-1) * perPage)

	database.DB.Raw(sql).Scan(&kar)

	affected := database.DB.Raw(sql).Scan(&kar)
	if affected.Error != nil{
		log.Println(affected.Error)
	}

	log.Println("DEBUGGER!");
	
	last_page := math.Ceil(float64(totalData/int64(perPage)))

	return ctx.JSON(fiber.Map{
		"data":kar,
		"total":totalData,
		"page":page,
		"last_page": last_page,
	})
	
	
	// var users []model.Karyawan

	// result := database.DB.Debug().Find(&users)
	// if result.Error != nil {
	// 	log.Println(result.Error)
	// }
	
	// return ctx.JSON(users)
}

func KaryawanHandlerCreate(ctx *fiber.Ctx) error {
	
	karyawan := new(model.KaryawanCreateRequest)
	if err := ctx.BodyParser(karyawan); err != nil {
		log.Println(err)
		return err
	}

	//set dafault role to be 2 , can not do in validator because not supported. 1:for Admin, 2:for user
	if karyawan.Role ==""{
		karyawan.Role="2"
	}

	if karyawan.Role == "1" || karyawan.Role =="2"{
		//todo 
	} else {
		return ctx.Status(400).JSON(fiber.Map{
			"message":"failed to create Karyawan ",
			"reason":"please check the role input ["+karyawan.Role+"]",
		})
	}

	//Validation
	validate := validator.New()
	errValidate := validate.Struct(karyawan)
	if errValidate != nil{
		return ctx.Status(400).JSON(fiber.Map{
			"message":"failed",
			"erorr":errValidate.Error(),
		})
	}

	newKaryawan := model.Karyawan{
		Name : karyawan.Name,
		Email : karyawan.Email,
		Address : karyawan.Address,
		Phone : karyawan.Phone,
		Role : karyawan.Role,
	}

	hashedPass,err := utills.PasswordUtilsHashing(karyawan.Password)
	if err != nil{
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal Server Erorr",
		})
	}

	newKaryawan.Password = hashedPass
	
	errCreate := database.DB.Create(&newKaryawan).Error
	log.Println(errCreate)
	if errCreate != nil{
		return ctx.Status(500).JSON(fiber.Map{
			"message":"failed create karyawan",
		})
	}

	return ctx.JSON(fiber.Map{
		"message":"succes created",
		"data" : newKaryawan,
	})
}


func KaryawanHandlerGetById(ctx *fiber.Ctx) error {
	idKaryawan := ctx.Params("id")
	
	var karyawan model.Karyawan
	err := database.DB.First(&karyawan , "id = ?", idKaryawan).Error
	if err != nil{
		return ctx.Status(404).JSON(fiber.Map{
			"message":"Karyawan not found",
		})
	}

	karyawanRespon := model.KaryawanCreateRespon{
		ID: karyawan.ID,
		Name: karyawan.Name,
		Email: karyawan.Email,
		Address: karyawan.Address,
		Phone: karyawan.Phone,
		CreatedAt: karyawan.CreatedAt,
		UpdatedAt: karyawan.UpdatedAt,
	}

	return ctx.JSON(fiber.Map{
		"message":"success",
		"data":karyawanRespon,
	})
}


func KaryawanHandlerUpdateById(ctx *fiber.Ctx) error {

	//Checck authentication
	userInfo := ctx.Locals("karyawanInfo").(jwt.MapClaims)
	if userInfo["role"] != "1"{
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden acces",
			"reason": "Only for Admin",
		})
	}

	karyawanRequest := new(model.KaryawanUpdateRequest)
	if err := ctx.BodyParser(karyawanRequest); err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(fiber.Map{
			"message":"Bad Request",
		})
	}

	var karyawan model.Karyawan

	//Check if Karyawan exist or not
	idKaryawan := ctx.Params("id")
	err := database.DB.First(&karyawan , "id = ?", idKaryawan).Error
	if err != nil{
		return ctx.Status(404).JSON(fiber.Map{
			"message":"Karyawan not found",
		})
	}

	//Update Data
	if karyawanRequest.Name != ""{
		karyawan.Name = karyawanRequest.Name
	}
	karyawan.Address = karyawanRequest.Address
	karyawan.Phone = karyawanRequest.Phone
	
	errUpdate := database.DB.Save(&karyawan).Error
	if errUpdate != nil{
		return ctx.Status(500).JSON(fiber.Map{
			"message":"Internal Server Erorr",
		})
	}

	return ctx.JSON(fiber.Map{
		"message":"success",
		"data":karyawan,
	})
}


func KaryawanHandlerDeleteById(ctx *fiber.Ctx) error {
	//Checck authentication
	userInfo := ctx.Locals("karyawanInfo").(jwt.MapClaims)
	if userInfo["role"] != "1"{
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden acces",
			"reason": "Only for Admin",
		})
	}
	
	idKaryawan := ctx.Params("id")
	var karyawan model.Karyawan

	//Check if Karyawan exist or not
	err := database.DB.Debug().First(&karyawan,"id = ? ",idKaryawan).Error
	if err != nil{
		return ctx.Status(404).JSON(fiber.Map{
			"message":"Karyawan Not Found",
		})
	}
	
	//Delete User
	errDelete := database.DB.Delete(&karyawan).Error
	if errDelete != nil{
		return ctx.Status(500).JSON(fiber.Map{
			"message":"Internal Server Errorr",
		})
	}

	return ctx.JSON(fiber.Map{
		"message":"Karyawan Deleted",
	})
}