package endpoints

import (
	"pauldieterich/go_fiber_pet/pkg/queries"
	"strconv"

	"github.com/gofiber/fiber/v2"
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
