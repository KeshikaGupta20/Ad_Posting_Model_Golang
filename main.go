package main

import (
	
	"log"
	
	

	"github.com/KeshikaGupta20/Ad_Posting_Model_Golang/database"
	"github.com/KeshikaGupta20/Ad_Posting_Model_Golang/routes"

	"github.com/joho/godotenv"
    
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/* var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("view/*.html"))

}

func index(w http.ResponseWriter, r *http.Request) {
     tpl.ExecuteTemplate(w, "login.html", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "post.html", nil)
 
}

func about(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.html", nil)

} */

func main() {

	app := fiber.New()

	//use logger for detailed view
	app.Use(logger.New())

	//for intialiseing thr env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	/* http.HandleFunc("/signup", index)
	http.HandleFunc("/addpost", home)
	 */
	routes.RegisterRoutes(app)

	
	

	log.Fatal(app.Listen(":4000"))

	

}
