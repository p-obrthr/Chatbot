package main

import (
	"fmt"
	"log"
	"os"

	"Backend.Api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Message struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func main() {
	fmt.Println("Hello World")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://127.0.0.1:3000",
		AllowMethods: "GET,POST,OPTIONS",
		AllowHeaders: "Content-Type",
	}))

	messages := []Message{}

	apiKey, exists := os.LookupEnv("OPENAI_API_KEY")
	if !exists {
		log.Fatal("no apikey found in envionment variable")
	}

	systemPrompt := "You are a very friendly and polite Chatbot. Use markdown for formatting the output."
	client := services.NewOpenAIClient(apiKey, systemPrompt)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	app.Get("/api/messages/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
		}

		for _, message := range messages {
			if message.Id == id {
				return c.Status(200).JSON(message)
			}
		}

		return c.Status(404).JSON(fiber.Map{
			"error": fmt.Sprintf("Message with id %d not found", id),
		})
	})

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		message := &Message{}

		if err := c.BodyParser(message); err != nil {
			return err
		}

		if message.Question == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Question is required"})
		}

		message.Id = len(messages) + 1

		// simulate processing
		// answer := strings.ToUpper(message.Question)
		// message.Answer = answer

		answer, err := client.SendMessage(message.Question)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		message.Answer = answer
		messages = append(messages, *message)

		return c.Status(201).JSON(message)
	})

	log.Fatal(app.Listen(":4000"))
}
