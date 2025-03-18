package main

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
)

type Job struct {
	JobID          string `json:"job_id"`
	Script         string `json:"script"`
	InputData      string `json:"input_data"`
	OutputLocation string `json:"output_location"`
}

var jobs = make(map[string]Job)

func publishJob(job Job) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "batch-processing-system")
	if err != nil {
		log.Printf("Failed to create Pub/Sub client: %v", err)
		return err
	}
	defer client.Close()

	topic := client.Topic("jobs")
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(job.JobID),
	})

	id, err := result.Get(ctx)
	if err != nil {
		log.Printf("Failed to publish job to Pub/Sub: %v", err)
		return err
	}

	log.Printf("Job published with ID: %s", id)
	return nil
}

func main() {
	r := gin.Default()

	r.POST("/jobs", func(c *gin.Context) {
		var job Job
		if err := c.ShouldBindJSON(&job); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		jobs[job.JobID] = job
		if err := publishJob(job); err != nil {
			c.JSON(500, gin.H{"error": "Failed to publish job"})
			return
		}
		c.JSON(200, gin.H{"message": "Job submitted", "job_id": job.JobID})
	})

	r.GET("/jobs/:id", func(c *gin.Context) {
		jobID := c.Param("id")
		job, exists := jobs[jobID]
		if !exists {
			c.JSON(404, gin.H{"error": "Job not found"})
			return
		}
		c.JSON(200, job)
	})

	r.Run() // Listen and serve on 0.0.0.0:8080
}
