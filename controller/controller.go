package controller

import (
	"log"
	database "github.com/devMukulSingh/fiberGoCrud.git/db"
	"github.com/devMukulSingh/fiberGoCrud.git/model"
	"github.com/gofiber/fiber/v2"
)

func HomeController(c *fiber.Ctx) error {

	return c.JSON("hellow world")
}

func PostController(c *fiber.Ctx) error {

	newBlog := new(model.Blog)

	if err := c.BodyParser(&newBlog); err != nil {
		log.Printf("error parsing request body %s", err.Error())
		return c.JSON("Error parsing request body")
	}

	result := database.DbConn.Create(newBlog)

	if result.Error != nil {
		log.Printf("Error in saving data %s", result.Error.Error())
		return c.JSON("Error in saving data ")
	}

	return c.JSON("saved sucess")
}

func UpdatController(c *fiber.Ctx) error {

	id := c.Params("id")

	var record model.Blog

	database.DbConn.First(&record, id)

	if record.Id == 0 {
		log.Printf("Record not found")
		c.Status(400)
		return c.JSON("Error: record not found")
	}

	result := database.DbConn.Save(record)

	if result.Error != nil {
		log.Println("Error in saving data")
		c.Status(500)
		return c.JSON("Error in saving data")
	}
	c.Status(200)
	return c.JSON("update success")

}

func DeleteController( c * fiber.Ctx) error{
	id := c.Params("id");

	var record model.Blog

	database.DbConn.First(&record,id);

	if record.Id == 0{
		log.Printf("NO record found")
		c.Status(400)
		return c.JSON("No record found")
	}

	if err := database.DbConn.Delete(&record); err!= nil{
		log.Printf("Error in deleting %s",err.Error.Error());
		c.Status(500)
		return c.JSON("error: record delte failed");
	}

	c.Status(200)
	return c.JSON("record deleted successfully")

}