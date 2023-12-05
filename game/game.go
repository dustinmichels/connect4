package game

type Game struct {
	Match   *Match
	Players []Player
}

func NewGame() *Game {
	players := []Player{
		{Name: "Player 1"},
		{Name: "Player 2"},
	}
	return &Game{
		Match:   NewMatch(players),
		Players: players,
	}
}

func (g *Game) Start() {
	g.Match.Start()
}
