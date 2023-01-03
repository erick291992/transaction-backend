package controllers

import (
	db "github.com/erick291992/transaction-backend/src/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	Repo *db.Repo
}

type AccountData struct {
	DocumentNumber string `json:"documentNumber"`
}

func (ac *AccountController) CreateAccounts(c *fiber.Ctx) error {
	// Parse the request body and get the AccountData struct.

	var data AccountData
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if len(data.DocumentNumber) == 0 {
		return c.JSON(fiber.NewError(400, "missing documentNumber"))
	}
	// Call the CreateAccount method on the Repo field of the AccountController struct.
	account, err := ac.Repo.CreateAccount(c.Context(), data.DocumentNumber)
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
		// return c.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}

	// Return the created account.
	return c.JSON(account)
}

func (ac *AccountController) GetAccount(c *fiber.Ctx) error {

	accountId, err := c.ParamsInt("accountId")
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	account, err := ac.Repo.GetAccount(c.Context(), int64(accountId))
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	// Return the account.
	return c.JSON(account)
}
