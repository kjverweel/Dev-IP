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
	cookie, err := e.Cookie("User") //get User_ID from cookie
	if err != nil {
		log.Println("couldn't get cookie")
	}

	GroepadminID, err := strconv.ParseUint(cookie.Value, 10, 64) //convert from cookie
	Groupies := &models.Groups{
		Groepname: e.FormValue("GroupName"),
	}
	Usersies := &models.Users{
		UserNickname: e.FormValue("UserName"),
	}

	UserID, err := repositories.CompareUsername(Usersies)
	GroupID, err := repositories.CompareGroupname(Groupies)
	CheckForAdmin := repositories.IsAnAdmin(int(GroepadminID), GroupID)
	GroepID, err := repositories.GetGroupsFromMembers(int(GroepadminID))
	RecentPosts, err := repositories.GetRecentPosts(GroepID)
	groups, err := repositories.GetGroup()
	AllUsers, err := repositories.GetAllUsers()

	if CheckForAdmin == false {
		HaHaGeenAdminLoser := "Sorry, maar je bent geen admin voor deze groep. Je kan niemand toevoegen."
		return e.Render(http.StatusUnauthorized, "member", echo.Map{"Admin_message": HaHaGeenAdminLoser, "Groups": groups, "RecentPosts": RecentPosts, "Users": AllUsers})
	} else if CheckForAdmin == true {
		IsMember, err := repositories.CheckIfInGroup(UserID, GroupID)
		if IsMember == true {
			UserIsAlMember := "Sorry, maar deze user is al toegevoegd aan de groep, check de naam nog een keer"
			return e.Render(http.StatusUnauthorized, "member", echo.Map{"Admin_message": UserIsAlMember, "Groups": groups, "RecentPosts": RecentPosts, "Users": AllUsers})
		} else if IsMember == false {

			WelGeenAdmin := e.FormValue("Admin")
			Admin := false
			if WelGeenAdmin == "Wel Admin" {
				Admin = true // set Admin to true if "Wel Admin" was submitted
			}
			NewMembie := &models.Groupmembers{
				UserID:  UserID,
				GroepID: GroupID,
				Admin:   Admin, // use the Admin variable declared outside of the if block
			}
			err = repositories.NewMember(NewMembie)
			err = e.Redirect(http.StatusSeeOther, "/home")
			if err != nil {
				log.Println("Error redirecting:", err)
			}
		}
		return nil
	}
	return nil
}
