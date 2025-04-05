package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Todo struct {
	ID          uuid.UUID `json:"_id,omitempty" bson:"_id,omitempty"`
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

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	if err = client.Ping(context.Background(), nil); err != nil {
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

func getTodo(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer func(cursor *mongo.Cursor, c context.Context) {
		err := cursor.Close(c)
		if err != nil {
			log.Fatal(err)
		}
	}(cursor, context.Background())

	for cursor.Next(context.Background()) {
		var todo Todo

		if err := cursor.Decode(&todo); err != nil {
			return err
		}

		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"msg": "body is empty"})
	}
	
	todo.ID = uuid.New()
	todo.DateCreated = time.Now()

	_, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	return nil
}

func deleteTodo(c *fiber.Ctx) error {
	return nil
}
