package controller

import (
	"fmt"
	
	"log"
	"regexp"

	"github.com/KeshikaGupta20/Ad_Posting_Model_Golang/database"
	"github.com/KeshikaGupta20/Ad_Posting_Model_Golang/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// CreatePerson will create thw person in the people table
func Signup(c *fiber.Ctx) error {

	//tmpl := template.Must(template.ParseFiles("/view/login.html"))
	

	reg := new(models.Register)

	db := database.DB

	err := c.BodyParser(&reg)
	if err != nil {
		return err
	}

	reg.Password = getHash([]byte(reg.Password))

	ValidateEmail := reg.Email

	CheckEmail := regexp.MustCompile(`^[a-z0-9._\-]+@[a-z0-9\-]+\.[a-z]{2,4}$`)

	Check := (CheckEmail.MatchString(ValidateEmail))

	if Check == true {

		 db.Find(&reg.Email)

		result := db.Create(&reg)

		if result != nil {
			fmt.Println("User Registered")

			return c.JSON(fiber.Map{
				"message": true,
				"data":    reg,
			})

		}

	}
	return c.JSON(fiber.Map{
		"message": true,
	})
}

func GetUser(c *fiber.Ctx) error {

	db := database.DB

	var people []models.Register

	db.Find(&people)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All users",
		"data":    people,
	})
}

func Login(c *fiber.Ctx) error {

	db := database.DB

	var dbuser models.Register
	user := new(models.Register)

    //dbuser = db.Query("select password from register where email=user.email ")

	//dbuser.Password = "123456"

	fmt.Println(dbuser.Password)

	err := c.BodyParser(user)
	if err != nil {
		return err
	}

	 db.Find(user.Email, dbuser.Password)
	 
	//db.Find(dbuser.Password)

	fmt.Println(user.Email)
	fmt.Println(dbuser.Password)

	userPass := []byte(user.Password)
	dbPass := []byte(dbuser.Password)

	fmt.Println(userPass)
	fmt.Println(dbPass) 

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	fmt.Println(passErr)

	if passErr != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Wrong Password",
		})
	}
	/* validToken, err := CreateToken(u.Email)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	} */

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"email":  user.Email,
			"status": "log in",
			//"token": validToken,
		},
	})

}
