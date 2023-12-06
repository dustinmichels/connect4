package game

type Match struct {
	Board         *Board
	Moves         []int
	Players       []Player
	player1Active bool
}

func NewMatch(players []Player) *Match {
	if len(players) != 2 {
		panic("Must have exactly 2 players")
	}
	return &Match{
		Board:         NewBoard(),
		Moves:         []int{},
		Players:       players,
		player1Active: true,
	}
}

func (m *Match) IsPlayer1Active() bool {
	return m.player1Active
}

// Apply move for current player
func (m *Match) ApplyMove(col int) error {
	err := m.Board.Update(m.player1Active, col)
	if err != nil {
		return err
	}
	m.Moves = append(m.Moves, col)
	m.player1Active = !m.player1Active
	return nil
}
