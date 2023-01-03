package controllers

import (
	db "github.com/erick291992/transaction-backend/src/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type OperationsTypeController struct {
	Repo *db.Repo
}

type OperationsTypeData struct {
	Description string `json:"description"`
}

func (otc *OperationsTypeController) CreateOperationsType(c *fiber.Ctx) error {

	var data OperationsTypeData
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if len(data.Description) == 0 {
		return c.JSON(fiber.NewError(400, "missing description"))
	}
	// Call the CreateAccount method on the Repo field of the AccountController struct.
	operationsType, err := otc.Repo.CreateOperationsType(c.Context(), data.Description)
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	// Return the created account.
	return c.JSON(operationsType)
}

func (otc *OperationsTypeController) GetOperationsType(c *fiber.Ctx) error {

	operationsTypeId, err := c.ParamsInt("operationsTypeId")
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	operationsType, err := otc.Repo.GetOperationsType(c.Context(), int64(operationsTypeId))
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	// Return the operationsType.
	return c.JSON(operationsType)
}
