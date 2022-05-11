package models

import(
	"github.com/jinzhu/gorm"
	"github.com/mtrosin/go-bookstore/pkg/config"
)

type User struct {
	gorm.Model
	Login string `gorm:""json:"login"`
	Password string `json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func GetUser(Login string, Password string) (*User, *gorm.DB) {
	var user User
	db := db.Where("login=? AND password=?", Login, Password).Find(&user)
	return &user, db
}