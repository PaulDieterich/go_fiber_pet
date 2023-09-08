package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"pauldieterich/go_fiber_pet/pkg/models"
	"pauldieterich/go_fiber_pet/pkg/queries"
	"strconv"
)

func HandelGetAllUsers(ctx *fiber.Ctx) error {

	return ctx.JSON("getAllUsers")
}

func HandelGetUser(ctx *fiber.Ctx) error {
	idString := ctx.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}
	user, err := queries.GetPetById(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not get user")
	}
	return ctx.JSON(user)

}

func HandelSaveUser(ctx *fiber.Ctx) error {
	idString := ctx.Params("id")
	//TODO: Implement
	return ctx.JSON(idString)
}

func HandelEditUser(ctx *fiber.Ctx) error {
	newUser := models.User{}
	ctx.Accepts("application/json")
	if err := ctx.BodyParser(&newUser); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	if err := queries.SaveUser(&newUser); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}

func HandelDeleteUser(ctx *fiber.Ctx) error {
	//TODO: Implement
	return nil
}

func HandelUserDerails(ctx *fiber.Ctx) error {
	//TODO: Implement
	return nil
}
