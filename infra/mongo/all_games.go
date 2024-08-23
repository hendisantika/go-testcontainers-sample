package mongo

import (
	"context"
	"go-testcontainers-sample/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type AllGames struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewAllGames(client *mongo.Client) *AllGames {
	collection := client.Database("testing").Collection("games")
	return &AllGames{client: client, coll: collection}
}

func (a AllGames) All() []*domain.Game {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter, err := a.coll.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var allGamesDocs []bson.M
	if err = filter.All(ctx, &allGamesDocs); err != nil {
		log.Fatal(err)
	}
	var allGames []*domain.Game
	for _, game := range allGamesDocs {
		allGames = append(allGames, toDomain(game))
	}
	return allGames
}
