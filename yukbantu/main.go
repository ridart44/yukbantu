package main

import (
	"log"
	"yukbantu/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/yukbantu?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	user := user.User{
		Name:           "InYourDream",
		Occupation:     "Pro Esport Player",
		Email:          "kiky.casa@gmail.com",
		PasswordHash:   "kikyInYourDream",
		AvatarFileName: "IYD.jpg",
		Role:           "user",
	}

	userRepository.Save(user)
}
