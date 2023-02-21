package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"html/template"
	"main.go/handlers"
	"main.go/models"
	"main.go/repositories"
)

func main() {
	e := echo.New()
	tpl, err := template.ParseGlob("./templates/*html")
	if err != nil {
		fmt.Println("Error loading templates: ", err)
		return
	}
	t := models.NewTemplate(tpl)
	e.Renderer = t
	repositories.DatabaseConnection()

	e.GET("/", handlers.Start)
	e.GET("/login", handlers.Loginscreen)
	e.POST("/login", handlers.Login)
	e.GET("/home", handlers.Home)
	e.GET("/register", handlers.Registerscreen)
	e.POST("/register", handlers.Register)
	e.GET("/groups", handlers.Groepjes)
	if err := e.Start(":1323"); err != nil {
		fmt.Println("Error starting the server: ", err)
	}
}
