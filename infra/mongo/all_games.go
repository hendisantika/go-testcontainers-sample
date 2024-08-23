package mongo

import "go.mongodb.org/mongo-driver/mongo"

type AllGames struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewAllGames(client *mongo.Client) *AllGames {
	collection := client.Database("testing").Collection("games")
	return &AllGames{client: client, coll: collection}
}
