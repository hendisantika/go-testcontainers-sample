package mem

import "sync"

type AllGames struct {
	db sync.Map
}
