package infra

import (
	"github.com/google/uuid"
	"go-testcontainers-sample/domain"
	"testing"
)

func Test_Add(t *testing.T) {
	gameId := domain.GameId(uuid.NewString())

	allGames.Add(&domain.Game{
		Id:    gameId,
		Title: "Assassin's Creed Valhalla",
		PEGI:  domain.Eighteen,
	})

	game := allGames.By(gameId)
	assert.Equal(t, gameId, game.Id)
	assert.Equal(t, "Assassin's Creed Valhalla", game.Title)
	assert.Equal(t, domain.Eighteen, game.PEGI)
}

func Test_Replace(t *testing.T) {
	gameId := domain.GameId(uuid.NewString())
	allGames.Add(&domain.Game{
		Id:    gameId,
		Title: "Rayman Raving Rabbids",
		PEGI:  domain.Seven,
	})

	allGames.Add(&domain.Game{
		Id:    gameId,
		Title: "Rainbow Six Siege",
		PEGI:  domain.Eighteen,
	})

	game := allGames.By(gameId)
	assert.Equal(t, gameId, game.Id)
	assert.Equal(t, "Rainbow Six Siege", game.Title)
	assert.Equal(t, domain.Eighteen, game.PEGI)
}
