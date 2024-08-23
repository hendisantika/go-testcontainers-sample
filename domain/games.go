package domain

type PEGI int
type GameId string

const (
	Three PEGI = iota
	Seven
	Twelve
	Sixteen
	Eighteen
)

type Game struct {
	Id    GameId
	Title string
	PEGI  PEGI
}
