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
	fmt.Println("hello")

	// connection, err := sql.Open("postgres", "postgresql://admin:secret@host.docker.internal:5432/main?sslmode=disable")

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connection, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@host.docker.internal:5432/%s?sslmode=disable", user, password, dbName))

	if err != nil {
		panic(err)
	}
	defer connection.Close()

	// Execute the SELECT COUNT(*) query.
	var count int
	err = connection.QueryRow("SELECT COUNT(*) FROM operations_types").Scan(&count)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if the count is 4.
	if count >= 4 {
		fmt.Println("There are already 4 rows in the operations_types table.")
	} else {
		fmt.Println("There are less than 4 rows in the operations_types table.")
		_, err := connection.Exec(
			"INSERT INTO operations_types (description) VALUES ($1), ($2), ($3), ($4)",
			"Normal Purchase", "Purchase with installments", "Withdrawal", "Credit Voucher")

		if err != nil {
			fmt.Println(err.Error())
		}
	}

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
