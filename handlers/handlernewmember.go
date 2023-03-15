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
	cookie, err := e.Cookie("User") //pakt uit het koekje de id van de user die is ingelogd
	if err != nil {
		log.Println("couldn't get cookie")
		return e.Render(http.StatusUnauthorized, "member", echo.Map{"Admin_message": "sorry, maar je koekje kan niet geladen worden."}) //logt de error en stuurt de user terug aar de login
	}

	GroepadminID, err := strconv.ParseUint(cookie.Value, 10, 64) //zet het koekje om van struct naar een Uint en slaat het op in AdminID
	Groupies := e.FormValue("GroupName")                         //slaat de GroupName formfield op in var Groupies
	Usersies := &models.Users{ //slaat de UserName op in de Users Struct in var Usersies
		UserNickname: e.FormValue("UserName"),
	}

	UserID, err := repositories.CompareUsername(Usersies)                //gebruikt Usersies om de ID van de user te krijgen, en dit op slaan in UserID
	GroupID, err := repositories.CompareGroupname(Groupies)              //Gebruikt Groepies om de GroupID te krijgen van de ingevulde groep, en dit opslaan in GroupID
	CheckForAdmin := repositories.IsAnAdmin(int(GroepadminID), GroupID)  //Gebruikt GroepadminID (van de cookie) en GroupID om te kijken of de user die de andere user probeert toe te voegen een admin is van de groep, en dit als boolean opslaan in IsAdmin
	GroepID, err := repositories.GetGroupsFromMembers(int(GroepadminID)) //gebruikt de GroepAdminID om alle Groepen waar de user(elke user, niet alleen de admin, maar dit is hoe ik de user_cookie heb genoemd) deel van uit maakt opslaat in GroupID als int slice
	RecentPosts, err := repositories.GetRecentPosts(GroepID)             //gebruikt de GroupID uit de lijn daarbove om alle meest recente posts the halen die zijn gemaakt in de groepen met de ID's die overeenkomen met de ID's uit de slice
	groups, err := repositories.GetGroup()                               //GetGroup pakt alle groepen en slaat die op in de slice string groups
	AllUsers, err := repositories.GetAllUsers()                          //GetAllUsers pakt alle Users en slaat deze op in de slice string AllUsers

	if CheckForAdmin == false { //als de boolean CheckForAdmin false is betekent dit dat de user geen admin is en dus geen user kan toevoegen
		HaHaGeenAdminLoser := "Sorry, maar je bent geen admin voor deze groep. Je kan niemand toevoegen." //hierin sturen we een bericht mee aan de e.render met een bericht waarom hij/zij nieman toe kan voegen
		return e.Render(http.StatusUnauthorized, "member", echo.Map{"Admin_message": HaHaGeenAdminLoser, "Groups": groups, "RecentPosts": RecentPosts, "Users": AllUsers})
	} else if CheckForAdmin == true { //als de boolean CheckForAdmin true is, gaat de code verder met het proces van toevoegen van de users
		IsMember, err := repositories.CheckIfInGroup(UserID, GroupID)
		if IsMember == true { //als de boolean IsMember true is betekent dit dat de user al in de groep zit en dus niet meer toegevoegd kan worden. Het proces wordt afgebroken en je wordt teruggestuurd naar /member met een bericht
			UserIsAlMember := "Sorry, maar deze user is al toegevoegd aan de groep, check de naam nog een keer"
			return e.Render(http.StatusUnauthorized, "member", echo.Map{"Admin_message": UserIsAlMember, "Groups": groups, "RecentPosts": RecentPosts, "Users": AllUsers})
		} else if IsMember == false {
			//als IsMember false is, is de user nog geen lid van de groep, dan kunnen we verder met het toevoegen
			WelGeenAdmin := e.FormValue("Admin") //WelGeenAdmin slaat op uit het HTML formulier of de User een admin is
			Admin := false                       //standaard is admin False, zo verminderen we de hoeveelheid code
			if WelGeenAdmin == "Wel Admin" {
				Admin = true //Admin slaat op of het False of True is
			}
			NewMembie := &models.Groupmembers{ //NewMembie zet alle gegevens in de struct voor Groupmembers, hierin staat de UserID vanuit CompareUsers, de GroepID vanuit CompareGroupname en Admin, die op basis van de form wel of geen admin invult
				UserID:  UserID,
				GroepID: GroupID,
				Admin:   Admin, // use the Admin variable declared outside of the if block
			}
			err = repositories.NewMember(NewMembie)        //NewMember pakt de struct en zet het in de database query, en voegt de user toe!
			err = e.Redirect(http.StatusSeeOther, "/home") //als alles goed is redirect ie naar Home
			if err != nil {
				log.Println("Error redirecting:", err) //anders print hij een log met een error
			}
		}
		return nil
	}
	return nil
}
