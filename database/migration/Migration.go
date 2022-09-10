package migration

import (
	"CURD_KARYAWAN/database"
	"CURD_KARYAWAN/model"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&model.Karyawan{})

	if err != nil{
		log.Println(err)
	}
	fmt.Println("Database Migrate")	
}