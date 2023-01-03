package routes

import (
	"fmt"

	"github.com/erick291992/transaction-backend/src/controllers"
	db "github.com/erick291992/transaction-backend/src/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

func OperationsTypeRoutes(router *fiber.Router, repo *db.Repo) {

	fmt.Println("OperationsTypeRoutes")

	opc := controllers.OperationsTypeController{
		Repo: repo,
	}

	(*router).Post("/operationsType", opc.CreateOperationsType)
	(*router).Get("/operationsType/:operationsTypeId", opc.GetOperationsType)
}
