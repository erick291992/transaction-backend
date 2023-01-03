package routes

import (
	"fmt"

	"github.com/erick291992/transaction-backend/src/controllers"
	db "github.com/erick291992/transaction-backend/src/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

func AccountRoutes(router *fiber.Router, repo *db.Repo) {

	fmt.Println("AccountRoutes")

	ac := controllers.AccountController{
		Repo: repo,
	}

	(*router).Post("/accounts", ac.CreateAccounts)
	(*router).Get("/accounts/:accountId", ac.GetAccount)
}
