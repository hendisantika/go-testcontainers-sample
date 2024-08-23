package test

import (
	"context"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestKafkaIntegration(t *testing.T) {
	ctx := context.Background()

	// Define a Kafka container
	req := testcontainers.ContainerRequest{
		Image:        "confluentinc/cp-kafka:7.7.0",
		ExposedPorts: []string{"9092/tcp"},
		WaitingFor:   wait.ForLog("listeners started on advertised listener"),
	}

	// Create the container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer container.Terminate(ctx)

	// Your Kafka test code here

	// Clean up
	if err := container.Terminate(ctx); err != nil {
		t.Fatal(err)
	}
}
