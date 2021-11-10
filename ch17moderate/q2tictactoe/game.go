package q2tictactoe

type TicTacToePlayer int

const (
	NoPlayer TicTacToePlayer = iota
	PlayerA
	PlayerB
)

type TicTacToeBoard [3][3]TicTacToePlayer

type TicTacToeGame struct {
	board TicTacToeBoard
}

func (g *TicTacToeGame) GameWinner() TicTacToePlayer {
	b := g.board
	// row (inner array)
	for _, line := range b {
		if winner := lineWinner(line[0], line[1], line[2]); winner != NoPlayer {
			return winner
		}
	}
	// column
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

func lineWinner(e1, e2, e3 TicTacToePlayer) TicTacToePlayer {
	if e1 != NoPlayer && e1 == e2 && e1 == e3 {
		return e1
	}
	return NoPlayer
}
