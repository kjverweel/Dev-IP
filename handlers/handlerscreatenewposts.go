package handlers

import (
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"main.go/models"
	"main.go/repositories"
	"net/http"
	"os"
	"strconv"
)

func CreateNewPost(e echo.Context) error {

	//Cookiecode
	// get cookie from request
	cookie, err := e.Cookie("User")
	// parse cookie string value to uint
	userId, err := strconv.ParseUint(cookie.Value, 10, 64)
	log.Println("handlerscreatenwepost:", userId)
	if err != nil {
		//if an error occurs in Cookiecode this usually means that the user isn't logged in properly.
		//this e.Render causes a direct to the index page, where you can log in or register an account.
		log.Println("handlercreatenewpost.go:Couldn't get cookie")
		err := e.Render(http.StatusOK, "index", nil)
		if err != nil {
			return err
		}
	}

	//take userid for laters
	userID := uint(userId)
	user := &models.Users{}
	err = repositories.GetUser(uint(userId), &user)
	if err != nil {
		log.Println("handlercreatenewpost.go:Couldn't get cookie")
	}
	log.Println("handlerscreatenwepost:", user) //prints userID for confirmation

	//end of Cookiecode
	//takes the name of the group
	Groepname := &models.Groups{
		Groepname: e.FormValue("GroupName"),
	}
	log.Println("handlerscreatenwepost:", Groepname) //prints groepname for confirmation

	//query's the groupname into the database to find the ID
	GroupID, err := repositories.CompareGroupname(Groepname)
	if err != nil {
		log.Println("handlercreatenewpost.go:couldn't find matching ID")
		return err
	}
	log.Println("handlercreatenewpost.go:GroupID is", GroupID)

	file, err := e.FormFile("Image")
	filename := file.Filename
	log.Println("filename: ", filename)

	if err != nil {
		// Return a 400 Bad Request error to the client if no file was uploaded or if there was an error parsing the form data.
		return echo.NewHTTPError(http.StatusBadRequest, "Image file not found")
	}
	src, err := file.Open()
	if err != nil {
		// Return a 500 Internal Server Error to the client if there was an error opening the file.
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("./uploads/" + filename)
	if err == nil {
		log.Println("THIS BETTER FUCKING WORK")
	}
	if err != nil {
		// Return a 500 Internal Server Error to the client if there was an error creating the file.
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	if _, err := io.Copy(dst, src); err != nil {
		// Return a 500 Internal Server Error to the client if there was an error copying the contents of the file.
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	filelocation := "./uploads/" + filename
	log.Println(filelocation)

	//data uploaded to the database
	Post := &models.Posts{
		PostContent:       e.FormValue("PostContent"),
		UserID:            int(userID),
		GroepID:           GroupID,
		PostImageLocation: filelocation,
	}

	log.Println(Post) //for confirmation the correct data was passed
	//appends the function that uploads the data to the database
	err = repositories.NewPost(Post)
	if err != nil {
		log.Println("you done fucked up now boi")
	} else {
		log.Println("Succes!")
	}
	return e.Redirect(http.StatusSeeOther, "/home")
}
