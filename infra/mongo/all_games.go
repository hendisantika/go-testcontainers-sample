package mongo

import "go.mongodb.org/mongo-driver/mongo"

type AllGames struct {
	client *mongo.Client
	coll   *mongo.Collection
}
