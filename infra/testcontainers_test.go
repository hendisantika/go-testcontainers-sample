package infra

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"go-testcontainers-sample/domain"
	"go-testcontainers-sample/infra/mem"
	"go-testcontainers-sample/infra/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"testing"
	"time"
)

type image struct {
	name string
	port string
}

var (
	allGames domain.AllGames

	mongoDB = image{
		name: "mongo:5.0.8",
		port: "27017",
	}
)

type container struct {
	testcontainers.Container
	URI string
}

func TestMain(m *testing.M) {
	var code = 1
	defer func() { os.Exit(code) }()

	ctx := context.Background()
	mongoContainer, err := setup(ctx, mongoDB)
	if err != nil {
		log.Printf("Unexpected error, fallback to mem repository implementations.\nerror: %s.\n", err)
		allGames = mem.NewAllGames()
	} else {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		client, err := mdriver.Connect(ctx, options.Client().ApplyURI(mongoContainer.URI))
		if err != nil {
			log.Printf("Could not connect to mongodb, fallback to mem repository implementations.\nerror: %s.\n", err)
			allGames = mem.NewAllGames()
		} else {
			allGames = mongo.NewAllGames(client)
		}
		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()
	}

	code = m.Run()
}
