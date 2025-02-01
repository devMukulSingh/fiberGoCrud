package router

import (
	"github.com/devMukulSingh/fiberGoCrud.git/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	app.Get("/",controller.HomeController)
	app.Post("/post",controller.PostController)
}