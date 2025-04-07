package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type Todo_old struct {
	ID          uuid.UUID `json:"id"`
	IsCompleted bool      `json:"completed"`
	Body        string    `json:"body"`
	DateCreated time.Time `json:"date_created"`
}

func main_old() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}

	PORT := os.Getenv("PORT")

	var todos []Todo_old

	//List all To do
	app.Get("/api_old/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	//Create a To do
	app.Post("/api_old/todos", func(c *fiber.Ctx) error {

		var todo = &Todo_old{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"msg": "body is empty"})
		}

		todo.ID = uuid.New()
		todo.DateCreated = time.Now()
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	//Update a To do status
	app.Patch("/api_old/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].IsCompleted = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"msg": "Todo id not found"})
	})

	app.Put("/api_old/todos/:id", func(c *fiber.Ctx) error {
		var todo = &Todo_old{}
		id := c.Params("id")

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"msg": "body is empty"})
		}

		for i, todoIndex := range todos {
			if fmt.Sprint(todoIndex.ID) == id {
				todos[i].Body = todo.Body
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"msg": "Todo id not found"})
	})

	//Delete a To do
	app.Delete("/api_old/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"msg": "success"})
			}

		}

		return c.Status(404).JSON(fiber.Map{"msg": "Todo id not found"})
	})
	log.Fatal(app.Listen(":" + PORT))
}
