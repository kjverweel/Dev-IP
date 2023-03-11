package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"log"
	"main.go/handlers"
	"main.go/models"
	"main.go/repositories"
)

func main() {
	e := echo.New()
	tpl, err := template.ParseGlob("./templates/*html")
	if err != nil {
		log.Println("Error loading templates: ", err)
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
	e.GET("/groups", handlers.Groups)
	e.POST("/home", handlers.CreateGroup)
	e.GET("/member", handlers.Member)
	e.POST("/member", handlers.GetNewMemberInfo)
	e.GET("/newpost", handlers.Posts)
	e.POST("/newpost", handlers.CreateNewPost)
	if err := e.Start(":1325"); err != nil {
		log.Println("Error starting the server: ", err)
	}
}
