package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
	"time"
)

func setupPostgresContainer(t *testing.T) (testcontainers.Container, string) {
	ctx := context.Background()

	// Define a PostgreSQL container
	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:17beta3-alpine3.20"),
		postgres.WithDatabase("example-db"),
		postgres.WithUsername("yuji"),
		postgres.WithPassword("S3cret"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(30*time.Second),
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Create the container
	connStr, err := container.ConnectionString(ctx, "sslmode=disable", "application_name=test")
	assert.NoError(t, err)

	return container, connStr
}
