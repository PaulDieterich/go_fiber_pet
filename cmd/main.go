package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"pauldieterich/go_fiber_pet/pkg/database"
	"pauldieterich/go_fiber_pet/pkg/endpoints"
)

type Pet struct {
	Id   int
	Name string
	Age  int
}

func helloWorld(c *fiber.Ctx) error {
	pets := []Pet{{1, "Floki", 5}, {2, "Fiete", 2}}

	return c.JSON(pets)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := database.Db.DB(); err != nil {
		log.Fatal("could not connect to database")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", helloWorld)
	app.Get("/pets", endpoints.HandelGetAllPets)
	app.Post("/pets", endpoints.SavePet)
	app.Get("/pets/:id", endpoints.HandleGetPet)
	app.Delete("/pets/:id", endpoints.HandelDeltePet)

	app.Listen(":3000")

}
