package infra

import (
	"github.com/testcontainers/testcontainers-go"
	"go-testcontainers-sample/domain"
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
