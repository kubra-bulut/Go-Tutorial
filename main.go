package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var db *sql.DB

func main() {
	// Initialize the database connection
	initDB()
	app := fiber.New()

	// Routes
	app.Get("/api/todos", getAllTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

	log.Fatal(app.Listen(":4000"))

}

func initDB() {
	var err error
	connStr := "user=postgres dbname= mydb sslmode=disable password=postgres"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected!")
}

func getAllTodos(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT id, body, completed FROM todos")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch todos"})
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Body, &todo.Completed); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "error scanning todos"})
		}
		todos = append(todos, todo)
	}

	return c.Status(200).JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	todo := &Todo{}
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "could not parse body"})
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "todo body is required"})
	}

	err := db.QueryRow("INSERT INTO todos (body, completed) VALUES ($1, $2) RETURNING id", todo.Body, todo.Completed).Scan(&todo.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not create todo"})
	}

	return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid todo ID"})
	}

	result, err := db.Exec("UPDATE todos SET completed = $1 WHERE id = $2", true, todoID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not update todo"})
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid todo ID"})
	}

	result, err := db.Exec("DELETE FROM todos WHERE id = $1", todoID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not delete todo"})
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}
