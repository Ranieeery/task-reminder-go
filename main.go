package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Todo struct {
	ID          uuid.UUID `json:"_id" bson:"_id"`
	IsCompleted bool      `json:"completed"`
	Body        string    `json:"body"`
	DateCreated time.Time `json:"date_created"`
}

var collection *mongo.Collection

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env: ", err)
	}

	PORT := os.Getenv("PORT")
	MongodbUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MongodbUri)
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	fmt.Println("Connected to MongoDB!")

	collection = client.Database("golang_todo_db").Collection("todo")

	app.Get("/api/todos", getTodo)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

	if PORT == "" {
		PORT = "4000"
	}

	log.Fatal(app.Listen(":" + PORT))
}

func getTodo(ctx *fiber.Ctx) error {
	return nil
}

func createTodo(ctx *fiber.Ctx) error {
	return nil
}

func updateTodo(ctx *fiber.Ctx) error {
	return nil
}

func deleteTodo(ctx *fiber.Ctx) error {
	return nil
}
