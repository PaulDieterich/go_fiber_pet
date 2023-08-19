package main

import (
	"pauldieterich/go_fiber_pet/pkg/database"
	"pauldieterich/go_fiber_pet/pkg/endpoints"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
)

func main() {
	connStr := "postgres://postgres:password@localhost/DB_1?sslmode=disable"
	database.InitDB(connStr)
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", helloWorld)
	app.Get("/pets/:id", endpoints.HandleGetPet)
	app.Delete("/pets/:id", endpoints.HandelDeltePet)

	app.Listen(":3000")

}
