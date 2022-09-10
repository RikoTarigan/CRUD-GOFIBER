package model

import (
	"time"

	"gorm.io/gorm"
)

type Karyawan struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password     string `json:"password"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Role     string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time`json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}