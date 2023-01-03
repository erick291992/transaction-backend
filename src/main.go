package main

import (
	"database/sql"
	"fmt"
	"os"

	db "github.com/erick291992/transaction-backend/src/db/sqlc"
	"github.com/erick291992/transaction-backend/src/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("hello API")

	// connection, err := sql.Open("postgres", "postgresql://admin:secret@host.docker.internal:5432/main?sslmode=disable")

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connection, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@host.docker.internal:5432/%s?sslmode=disable", user, password, dbName))

	if err != nil {
		panic(err)
	}
	defer connection.Close()

	// Create a new Repo.
	repo := db.NewRepo(connection)

	app := fiber.New()
	app.Use("/", func(c *fiber.Ctx) error {
		fmt.Println("Request Url:", c.Path())
		return c.Next()
	})

	router := app.Group("/api/v1")

	routes.AccountRoutes(&router, repo)
	routes.OperationsTypeRoutes(&router, repo)
	routes.TransactionRoutes(&router, repo)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Everyone 👋!")
	})

	app.Listen(":3000")
}
