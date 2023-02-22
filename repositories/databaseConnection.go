package repositories

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main.go/models"
)

var db *gorm.DB
var err error

func DatabaseConnection() {
	dsn := "Fredmeister:DikkieDik@(localhost:3306)/opdrachtip?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		fmt.Println("You did a fucky wucky senpai san")
	} else {
		fmt.Println("Database succesfully done diggery doo")
	}
	err = db.AutoMigrate(&models.Users{}, &models.Posts{}, &models.Groups{}, &models.Comments{}, &models.Likes{}, &models.Groupmembers{})
	if err != nil {
		fmt.Println("You did a fucky wucky senpai san")
	} else {
		fmt.Println("migrations complete")
	}

}
