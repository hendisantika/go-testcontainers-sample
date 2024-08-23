package mem

import (
	"go-testcontainers-sample/domain"
	"sync"
)

type AllGames struct {
	db sync.Map
}

func NewAllGames() *AllGames {
	return &AllGames{db: sync.Map{}}
}

func (a *AllGames) All() (allGames []*domain.Game) {
	a.db.Range(func(key, value interface{}) bool {
		allGames = append(allGames, value.(*domain.Game))
		return true
	})
	return
}

func (a *AllGames) Add(game *domain.Game) {
	a.db.Store(game.Id, game)
}
