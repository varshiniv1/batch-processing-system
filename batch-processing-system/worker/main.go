package main

import (
	"context"
	"log"
	"os/exec"

	"cloud.google.com/go/pubsub"
)

func processJob(jobID string) {
	// Simulate job processing
	log.Printf("Processing job: %s", jobID)

	// Use cmd.exe for Windows compatibility
	cmd := exec.Command("cmd.exe", "/C", "echo Processing job")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Job %s failed: %v", jobID, err)
		return
	}

	log.Printf("Job %s completed: %s", jobID, string(output))
}

func main() {
	ctx := context.Background()

	// Create a Pub/Sub client
	client, err := pubsub.NewClient(ctx, "batch-processing-system") // Use your project ID here
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Subscribe to the Pub/Sub topic
	sub := client.Subscription("jobs-subscription")
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		// Process the job
		processJob(string(msg.Data))

		// Acknowledge the message
		msg.Ack()
	})

	if err != nil {
		log.Fatalf("Failed to receive messages: %v", err)
	}
}
