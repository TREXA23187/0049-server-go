package main

import (
	"0049-server-go/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/gin_chat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("fail to connect db")
	}

	db.AutoMigrate(&models.UserModel{})

	user := &models.UserModel{}
	user.UserName = "UserName"
	db.Create(user)

	fmt.Println(db.First(user, 1))

	db.Model(user).Update("Password", "1234")
}
