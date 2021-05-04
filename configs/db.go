package configs

import (
	"fmt"
	"main/database"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() {

	var err error
	dsn := "root@tcp(127.0.0.1:3306)/wallet?charset=utf8mb4&parseTime=True&loc=Local"
	database.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}
