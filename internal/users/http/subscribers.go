package http

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync/atomic"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

var projectIDSub string

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", err)
	}
	projectIDSub = os.Getenv("PROJECT_ID")
}

func createUserSubs(w io.Writer, UsersDep *UsersDependencies) error {
	subID := "create-user-sub"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectIDSub)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
		atomic.AddInt32(&received, 1)

		// CALL HANDER MS
		// answerData := UsersDep.uh.CreateUser(msg.Data)
		// if err != nil {
		// 	log.Println(err)
		// }

		// fmt.Println(answerData)

		msg.Ack()
	})

	if err != nil {
		return fmt.Errorf("sub.Receive: %w", err)
	}

	fmt.Fprintf(w, "Received %d messages\n", received)

	return nil
}

func deleteUserSubs(w io.Writer, UsersDep *UsersDependencies) error {
	subID := "delete-user-sub"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectIDSub)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
		atomic.AddInt32(&received, 1)

		// CALL HANDER MS

		msg.Ack()
	})

	if err != nil {
		return fmt.Errorf("sub.Receive: %w", err)
	}

	fmt.Fprintf(w, "Received %d messages\n", received)

	return nil
}
