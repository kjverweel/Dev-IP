package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/models"
	"main.go/repositories"
	"net/http"
	"strconv"
)

func CreateNewPost(e echo.Context) error {

	//Cookiecode
	// get cookie from request
	cookie, err := e.Cookie("User")
	// parse cookie string value to uint
	userId, err := strconv.ParseUint(cookie.Value, 10, 64)
	log.Println(userId)
	if err != nil {
		//if an error occurs in Cookiecode this usually means that the user isn't logged in properly.
		//this e.Render causes a direct to the index page, where you can log in or register an account.
		log.Println("handlercreatenewpost.go:Couldn't get cookie")
		e.Render(http.StatusOK, "index", nil)
	}

	//take userid for laters
	userID := uint(userId)
	user := &models.Users{}
	err = repositories.GetUser(uint(userId), &user)
	if err != nil {
		log.Println("handlercreatenewpost.go:Couldn't get cookie")
	}
	log.Println(user)
	//end of Cookiecode
	//takes the name of the group
	Groepname := &models.Groups{
		Groepname: e.FormValue("GroupName"),
	}
	log.Println(Groepname)
	//query's the groupname into the database to find the ID
	GroupID, err := repositories.CompareGroupname(Groepname)
	if err != nil {
		log.Println("handlercreatenewpost.go:couldn't find matching ID")
		return err
	}
	log.Println("handlercreatenewpost.go:GroupID is", GroupID)

	Post := &models.Posts{
		PostContent: e.FormValue("PostContent"),
		UserID:      int(userID),
		GroepID:     GroupID,
	}

	log.Println(Post)

	return nil

}
