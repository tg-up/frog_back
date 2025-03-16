package main

import (
	"icecreambash/tgup_backend/internal/configs"
	"icecreambash/tgup_backend/pkg/database"
	"log"
)

func init() {
	configs.LoadConfig()
	err := database.InitDB()
	if err != nil {
		log.Fatalf("init db err: %v", err)
	}
	_, ch, err := database.LoadRabbitMQ()

	if err != nil {
		log.Fatalf("load rabbitmq err: %v", err)
	}
	database.RabbitMQ = ch
}

func main() {
	msgs, err := database.RabbitMQ.Consume(
		"order_queue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	log.Println("Waiting for messages...")
	for msg := range msgs {
		log.Printf("Received a task ID: %s \n", string(msg.Body))
	}
}
