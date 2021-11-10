package q2tictactoe

import "testing"

func TestGameWinner(t *testing.T) {
	N := NoPlayer
	A := PlayerA
	B := PlayerB
	cases := []struct {
		b TicTacToeBoard
		w TicTacToePlayer
	}{
		// empty board
		{TicTacToeBoard{{N, N, N}, {N, N, N}, {N, N, N}}, N},
		// row (inner array) winner
		{TicTacToeBoard{{A, A, A}, {N, N, N}, {N, N, N}}, A},
		{TicTacToeBoard{{A, A, N}, {B, B, B}, {N, A, A}}, B},
		{TicTacToeBoard{{N, N, N}, {N, N, N}, {A, A, A}}, A},
		// column winner
		{TicTacToeBoard{{A, N, N}, {A, N, N}, {A, N, N}}, A},
		{TicTacToeBoard{{N, B, A}, {N, B, N}, {N, B, A}}, B},
		{TicTacToeBoard{{N, N, A}, {N, N, A}, {N, N, A}}, A},
		// diagonal winner
		{TicTacToeBoard{{N, N, A}, {N, A, N}, {A, N, N}}, A},
		{TicTacToeBoard{{B, N, N}, {A, B, N}, {A, N, B}}, B},
		// no winner
		{TicTacToeBoard{{N, N, N}, {N, N, N}, {A, A, B}}, N},
		{TicTacToeBoard{{A, N, N}, {A, N, N}, {B, N, N}}, N},
		{TicTacToeBoard{{N, N, A}, {N, A, N}, {N, N, N}}, N},
	}
	for _, c := range cases {
		g := TicTacToeGame{c.b}
		if got := g.GameWinner(); got != c.w {
			t.Errorf("Got %v as winner, but expected %v", got, c.w)
		}
	}
}
