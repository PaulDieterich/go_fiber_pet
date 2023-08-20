package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
	"os"
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

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := database.Connect(config)
	if err != nil {
		log.Fatal("could not connect to database")
	}
	petRepository := database.Repository{
		DB: db,
	}
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", helloWorld)
	app.Get("/pets/:id", endpoints.HandleGetPet)
	app.Delete("/pets/:id", endpoints.HandelDeltePet)

	app.Listen(":3000")

}
