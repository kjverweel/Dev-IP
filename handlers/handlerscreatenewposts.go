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

	cookie, err := e.Cookie("User") //pakt uit het koekje de id van de user die is ingelogd
	// parse cookie string value to uint
	userId, err := strconv.ParseUint(cookie.Value, 10, 64) //zet het koekje om van struct naar een Uint
	if err != nil {
		log.Println("couldn't get cookie")
		return e.Render(http.StatusUnauthorized, "register", echo.Map{"NuhUh": "sorry, maar er gaat wat fout"}) //logt de error en stuurt de user terug aar de registratiepagina
	}
	//Pak het UserID voor laters
	userID := uint(userId)
	user := &models.Users{}
	err = repositories.GetUser(uint(userId), &user)
	if err != nil {
		log.Println("handlercreatenewpost.go:Couldn't get cookie")
	}
	//Pak de naam van de groep
	Groepname := e.FormValue("GroupName")
	log.Println("handlerscreatenwepost:", Groepname) //prints groepname for confirmation

	//zoekt de ID die bij de groepname hoort
	GroupID, err := repositories.CompareGroupname(Groepname)
	if err != nil {
		log.Println("handlercreatenewpost.go:couldn't find matching ID")
		return err
	}
	//hier pakt hij de files die in de formfield "image zitten
	file, err := e.FormFile("Image")

	filename := file.Filename //deze slaat hij op in de var filename

	// TODO: Elke bestandsnaam een unieke naam /id geven, je kan nu niet alle foto's uploaden

	//idee := naam omzetten naar een int, de laatse image eruit pakken, en int +1 doen. Oneindig lang doorgaan
	//daarvoor: filepath en name opslpitsen. filepath is fixed, name = int
	if err != nil {
		// stuurt een 400 error terug, als er geen image bij zat of er ging wat fout :)
		return echo.NewHTTPError(http.StatusBadRequest, "Image file not found")
	}
	src, err := file.Open() //het openen van de file is belangrijk om ervoor te zorgen dat de file goed wordt geuploadt
	if err != nil {         //als er iets fout gaat tijdens het verwerken van een afbeelding krijg je een 500, internalserver error
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer src.Close() //Het sluiten van de "src" zorgt ervoor dat in het verdere process geen problemen komen

	// os.Create slaat de filename op, met daarvoor de relatieve path voor waar de images worden opgeslagen
	dst, err := os.Create("./uploads/" + filename)
	if err != nil {
		//als er iets fout gaat tijdens het verwerken van een afbeelding krijg je een 500, internalserver error
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	//io.copy zorgt ervoor dat de afbeelding van de source, 'src', naar de destination, 'dst' word gezet
	if _, err := io.Copy(dst, src); err != nil {
		//als er iets fout gaat tijdens het verwerken van een afbeelding krijg je een 500, internalserver error
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	filelocation := "./uploads/" + filename //filelocation wordt later gebruikt in de struct om de afbeelding weer op te kunnen roepen

	//struct om de data in op te slaan die uit de forms zijn verzameld
	Post := &models.Posts{ //de struct die wordt gebruikt om de posts aan te maken
		PostContent:       e.FormValue("PostContent"),
		UserID:            int(userID),
		GroepID:           GroupID,
		PostImageLocation: filelocation,
	}
	//geeft de Post data door aan de repo om een nieuwe post aan te maken
	err = repositories.NewPost(Post)
	if err != nil {
		log.Println("Oh Oh hier gaat iets fout")
	} else {
		log.Println("Succes!")
	}
	return e.Redirect(http.StatusSeeOther, "/home")
}
