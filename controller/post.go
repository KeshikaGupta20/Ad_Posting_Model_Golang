package controller

import (
	"fmt"
	"github.com/KeshikaGupta20/Ad_Posting_Model_Golang/database"
	"github.com/KeshikaGupta20/Ad_Posting_Model_Golang/models"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {

	db := database.DB

	post := new(models.Post)

	err := c.BodyParser(post)

	if err != nil {
		return err
	}

	result := db.Create(&post)

	if result != nil {
		fmt.Println("Post sucessfully uploaded")

		return c.JSON(fiber.Map{
			"message": true,
			"data":    post,
		})

	}

	return c.JSON(fiber.Map{
		"message": true,
	})

}


func GetPost(c *fiber.Ctx) error {

	db := database.DB

	var post []models.Post

	db.Find(&post)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "All users",
		"data":    post,
	})
}

