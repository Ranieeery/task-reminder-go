package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
)

type Todo struct {
	ID          uuid.UUID `json:"id"`
	isCompleted bool      `json:"completed"`
	Body        string    `json:"body"`
}

func main() {
	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"msg": "body is empty"})
		}

		todo.ID = uuid.New()
		todos = append(todos, *todo)

		return c.Status(201).JSON(fiber.Map{"msg": "ok"})
	})

	log.Fatal(app.Listen(":4000"))
}
