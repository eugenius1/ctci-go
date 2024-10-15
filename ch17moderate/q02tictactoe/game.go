package q02tictactoe

type Player int

const (
	NoPlayer Player = iota
	PlayerA
	PlayerB
)

type Board [3][3]Player

type Game struct {
	board Board
}

func (g *Game) GameWinner() Player {
	b := g.board
	// rows (inner array)
	for _, line := range b {
		if winner := lineWinner(line[0], line[1], line[2]); winner != NoPlayer {
			return winner
		}
	}

	// columns
	for i := range b {
		if winner := lineWinner(b[0][i], b[1][i], b[2][i]); winner != NoPlayer {
			return winner
		}
	}

	// diagonals
	if winner := lineWinner(b[0][0], b[1][1], b[2][2]); winner != NoPlayer {
		return winner
	}

	if winner := lineWinner(b[0][2], b[1][1], b[2][0]); winner != NoPlayer {
		return winner
	}

	return NoPlayer
}

func lineWinner(e1, e2, e3 Player) Player {
	if e1 != NoPlayer && e1 == e2 && e1 == e3 {
		return e1
	}

	return NoPlayer
}
