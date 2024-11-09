package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Id       int    `json: id`
	Question string `json: question`
	Answer   string `json: answer`
}

func main() {
	fmt.Println("Hello World")
	app := fiber.New()

	messages := []Message{}

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

		//simulate processing
		answer := strings.ToUpper(message.Question)
		message.Answer = answer

		messages = append(messages, *message)

		return c.Status(201).JSON(message)
	})

	log.Fatal(app.Listen(":4000"))
}
