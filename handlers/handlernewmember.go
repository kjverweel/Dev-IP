package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"main.go/models"
	"main.go/repositories"
	"net/http"
	"strconv"
)

func GetNewMemberInfo(e echo.Context) error {
	//Cookiecode
	// get cookie from request
	cookie, err := e.Cookie("User")
	// parse cookie string value to uint
	userId, err := strconv.ParseUint(cookie.Value, 10, 64)
	if err != nil {
		//if an error occurs in Cookiecode this usually means that the user isn't logged in properly.
		//this e.Render causes a direct to the index page, where you can log in or register an account.
		log.Println("handlerhome.go:Couldn't get cookie")
		e.Render(http.StatusOK, "index", nil)
	}
	user := &models.Users{}
	err = repositories.GetUser(uint(userId), &user)
	if err != nil {
		log.Println("handlerhome.go:Couldn't get cookie")
	}
	//end of Cookiecode
	//start of the actual NewMember code
	if e.FormValue("UserName") == "" || e.FormValue("GroupName") == "" {
		return e.Render(http.StatusUnauthorized, "member", nil)
	}

	Usernickname := &models.Users{
		UserNickname: e.FormValue("UserName"),
	}
	Groepname := &models.Groups{
		Groepname: e.FormValue("GroupName"),
	}

	log.Println("handlernewmember.go:", Usernickname)
	log.Println("handlernewmember.go:", Groepname)

	UserID, err := repositories.CompareUsername(Usernickname)
	if err != nil {
		log.Println("handlernewmember.go:couldn't find matching ID")
		return err
	}

	GroupID, err := repositories.CompareGroupname(Groepname)
	if err != nil {
		log.Println("handlernewmember.go:couldn't find matching ID")
		return err
	}
	//helps confirm that the ID's are correctly pulled from the database
	log.Println("handlernewmember.go:UserID is", UserID)
	log.Println("handlernewmember.go:GroupID is", GroupID)

	Groupmembers := &models.Groupmembers{
		UserID:  UserID,
		GroepID: GroupID,
	}

	IsMember, err := repositories.CheckGroupMembers(Groupmembers)
	if err != nil {
		log.Println("handlernewmember.go:error checking group members:", err)
		return nil
	}
	if IsMember == true {
		log.Println("user is not yet a member of the group")
		err = repositories.NewMember(Groupmembers)
		if err != nil {
			log.Println("handlercreategroup.go:Repository got fucked")
		} else {
			log.Println("handlercreategroup.go:Succesfully called")
		}
		log.Println("Now he is a member")
	} else if IsMember == false {
		log.Println("User is a member and he can fuck off ")
	}
	//End of the NewMember Code

	groups, err := repositories.GetGroup()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get groups",
		})
	}

	RecentPosts, err := repositories.GetRecentPosts()
	if err != nil {
		log.Println("handlerhome.go:Couldn't get recents posts")
	}
	log.Println("handlerhome:", RecentPosts)

	if groups == nil {
		e.Render(http.StatusOK, "home", echo.Map{"Groups": "Unfortunately, there are no groups yet", "RecentPosts": RecentPosts})
	}
	err = e.Render(http.StatusOK, "home", echo.Map{"Nem": user.UserNickname, "Groups": groups})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
