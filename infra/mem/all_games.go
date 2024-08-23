package mem

import "sync"

type AllGames struct {
	db sync.Map
}

func NewAllGames() *AllGames {
	return &AllGames{db: sync.Map{}}
}
