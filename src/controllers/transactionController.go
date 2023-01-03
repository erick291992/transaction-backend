package controllers

import (
	db "github.com/erick291992/transaction-backend/src/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	Repo *db.Repo
}

func (tc *TransactionController) CreateTransaction(c *fiber.Ctx) error {

	var data db.CreateTransactionParams
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// if len(data.DocumentNumber) == 0 {
	// 	return c.JSON(fiber.NewError(400, "missing documentNumber"))
	// }

	transaction, err := tc.Repo.CreateTransaction(c.Context(), data)
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	// Return the created transaction.
	return c.JSON(transaction)
}

func (tc *TransactionController) GetTransaction(c *fiber.Ctx) error {

	transactionId, err := c.ParamsInt("transactionId")
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	transaction, err := tc.Repo.GetTransaction(c.Context(), int64(transactionId))
	if err != nil {
		return c.JSON(fiber.NewError(fiber.StatusInternalServerError, err.Error()))
	}

	// Return the transaction.
	return c.JSON(transaction)
}
