package q02tictactoe

import "testing"

func TestGameWinner(t *testing.T) {
	N := NoPlayer
	A := PlayerA
	B := PlayerB
	cases := []struct {
		b Board
		w Player
	}{
		// empty board
		{Board{{N, N, N}, {N, N, N}, {N, N, N}}, N},
		// row (inner array) winner
		{Board{{A, A, A}, {N, N, N}, {N, N, N}}, A},
		{Board{{A, A, N}, {B, B, B}, {N, A, A}}, B},
		{Board{{N, N, N}, {N, N, N}, {A, A, A}}, A},
		// column winner
		{Board{{A, N, N}, {A, N, N}, {A, N, N}}, A},
		{Board{{N, B, A}, {N, B, N}, {N, B, A}}, B},
		{Board{{N, N, A}, {N, N, A}, {N, N, A}}, A},
		// diagonal winner
		{Board{{N, N, A}, {N, A, N}, {A, N, N}}, A},
		{Board{{B, N, N}, {A, B, N}, {A, N, B}}, B},
		// no winner
		{Board{{N, N, N}, {N, N, N}, {A, A, B}}, N},
		{Board{{A, N, N}, {A, N, N}, {B, N, N}}, N},
		{Board{{N, N, A}, {N, A, N}, {N, N, N}}, N},
	}
	for _, c := range cases {
		g := Game{c.b}
		if got := g.GameWinner(); got != c.w {
			t.Errorf("Got %v as winner, but expected %v", got, c.w)
		}
	}
}
