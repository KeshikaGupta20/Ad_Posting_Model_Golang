package routes

import (
	

	a "github.com/KeshikaGupta20/Ad_Posting_Model_Golang/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app fiber.Router) {

	app.Post("/signup", a.Signup)

	app.Get("/getuser", a.GetUser)

	app.Post("/login", a.Login)

	app.Post("/addpost", a.CreatePost)

	app.Get("/getpost", a.GetPost)

	app.Delete("/deletepost", a.DeletePost)

	app.Post("/publishpost", a.PublishPost)

	
}
