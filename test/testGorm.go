package main

import (
	"fmt"
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	// 迁移schame
	db.AutoMigrate(&models.UserBasic{})

	user := models.UserBasic{}
	user.Name = "gin"
	db.Create(user)
	//根据整型主键查找
	fmt.Println(db.First(user, 1))
	// 查找code字段值为42的记录
	db.First(user, "code=?", "D42")
	// update
	db.Model(user).Update("Password", "1234")

}
