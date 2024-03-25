package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/StephanSuarez/chat-rooms-users-ms/internal/users/http/dtos"
	"github.com/joho/godotenv"
)

var projectIDPub string

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	projectIDPub = os.Getenv("PROJECT_ID")
}

func createUserPubResponse(w io.Writer, data *dtos.AnswerDTO) error {

	topicID := "create-user-response"

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectIDPub)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	t := client.Topic(topicID)

	bodyJSON, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal UsResponse to JSON: %v", err)
	}

	result := t.Publish(ctx, &pubsub.Message{
		Data: bodyJSON,
	})

	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("failed to publish: %v", err)
	}

	fmt.Fprintf(w, "Published message; msg ID: %v\n", id)

	return nil
}
