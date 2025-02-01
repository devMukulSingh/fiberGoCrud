package main

import (
	"log"
	"github.com/devMukulSingh/fiberGoCrud.git/db"
	"github.com/devMukulSingh/fiberGoCrud.git/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil{
		log.Fatal("Error loading env")
	}

	database.ConnectDB()

	sqlDb,err := database.DbConn.DB()

	if err!= nil{
		panic("Error in sql")
	}

	defer sqlDb.Close()

	app := fiber.New();

	app.Use(logger.New())

	router.SetupRoutes(app);



	app.Listen(":8000")
}