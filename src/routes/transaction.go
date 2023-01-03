package routes

import (
	"fmt"

	"github.com/erick291992/transaction-backend/src/controllers"
	db "github.com/erick291992/transaction-backend/src/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

func TransactionRoutes(router *fiber.Router, repo *db.Repo) {

	fmt.Println("TransactionRoutes")

	tc := controllers.TransactionController{
		Repo: repo,
	}

	(*router).Post("/transactions", tc.CreateTransaction)
	(*router).Get("/transactions/:transactionId", tc.GetTransaction)
}
