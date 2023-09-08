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
	api := app.Group("/api")
	v1 := api.Group("/v1", apiHandler)
	app.Get("/", helloWorld)

	pets := v1.Group("/pets")
	pets.Get("/", endpoints.HandelGetAllPets)
	pets.Post("/", endpoints.SavePet)
	pets.Get("/:id", endpoints.HandleGetPet)
	pets.Delete("/:id", endpoints.HandelDeltePet)

	users := v1.Group("/users")
	users.Get("/", endpoints.HandelGetAllUsers)
	//TODO: implement
	users.Post("/")
	users.Get("/:id")
	users.Put("/:id")
	users.Delete("/id")
	users.Get("/:id/details")

	app.Listen(":3000")

}

func apiHandler(ctx *fiber.Ctx) error {
	ctx.Set("Version", "v1")
	return ctx.Next()
}
