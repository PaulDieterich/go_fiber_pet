package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"pauldieterich/go_fiber_pet/pkg/models"
	"pauldieterich/go_fiber_pet/pkg/queries"
	"strconv"
)

func HandelDeltePet(c *fiber.Ctx) error {
	id := c.Params("id")
	if id != "" {
		idInt, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
		}
		err = queries.DeletePet(idInt)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Could not delete pet")

		}
		return fiber.NewError(fiber.StatusOK, "Pet deleted")
	}
	return fiber.NewError(fiber.StatusBadRequest, "ID is emptry")
}

func HandleGetPet(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}
	pet, err := queries.GetPetById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not get pet")
	}
	return c.JSON(pet)
}
func HandelGetAllPets(c *fiber.Ctx) error {
	//TODO implement
	return nil
}
func SavePet(c *fiber.Ctx) error {
	newPet := models.Pet{}
	c.Accepts("application/json")
	if err := c.BodyParser(&newPet); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := queries.Save(&newPet); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(newPet)
}
