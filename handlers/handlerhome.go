package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/models"
	"main.go/repositories"
	"net/http"
	"strconv"
)

func Home(e echo.Context) error {
	// get cookie from request
	cookie, err := e.Cookie("User")
	// parse cookie string value to uint
	userId, err := strconv.ParseUint(cookie.Value, 10, 64)
	if err != nil {
		log.Println("handlerhome.go:Couldn't get cookie")
		e.Render(http.StatusOK, "index", nil)
	}
	user := &models.Users{}
	err = repositories.GetUser(uint(userId), &user)
	if err != nil {
		log.Println("handlerhome.go:Couldn't get cookie")
	}
	groups, err := repositories.GetGroup()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get groups",
		})
	}
	if groups == nil {
		e.Render(http.StatusOK, "home", echo.Map{"Groups": "Unfortunately, there are no groups yet"})
	}

	GroepID, err := repositories.GetGroupsFromMembers(int(userId))

	RecentPosts, err := repositories.GetRecentPosts(GroepID[0])
	if err != nil {
		log.Println("handlerhome.go:Couldn't get recents posts")
	}
	log.Println("handlerhome:", RecentPosts)

	err = e.Render(http.StatusOK, "home", echo.Map{"Nem": user.UserNickname, "Groups": groups, "RecentPosts": RecentPosts})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
