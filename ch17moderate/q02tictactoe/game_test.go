package q02tictactoe

import "testing"

func TestGameWinner(t *testing.T) {
	t.Parallel()

	N := NoPlayer
	A := PlayerA
	B := PlayerB
	tests := []struct {
		board  Board
		winner Player
		name   string
	}{
		{Board{{N, N, N}, {N, N, N}, {N, N, N}}, N, "empty board"},
		{Board{{A, A, A}, {N, N, N}, {N, N, N}}, A, "row 1 winner"},
		{Board{{A, A, N}, {B, B, B}, {N, A, A}}, B, "row 2 winner"},
		{Board{{N, N, N}, {N, N, N}, {A, A, A}}, A, "row 3 winner"},
		{Board{{A, N, N}, {A, N, N}, {A, N, N}}, A, "column 1 winner"},
		{Board{{N, B, A}, {N, B, N}, {N, B, A}}, B, "column 2 winner"},
		{Board{{N, N, A}, {N, N, A}, {N, N, A}}, A, "column 3 winner"},
		{Board{{N, N, A}, {N, A, N}, {A, N, N}}, A, "diagonal TR-BL winner"},
		{Board{{B, N, N}, {A, B, N}, {A, N, B}}, B, "diagonal TL-BR winner"},
		{Board{{N, N, N}, {N, N, N}, {A, A, B}}, N, "no winner - bottom row"},
		{Board{{A, N, N}, {A, N, N}, {B, N, N}}, N, "no winner - column 1"},
		{Board{{N, N, A}, {N, A, N}, {N, N, N}}, N, "no winner - diagonal TR-BL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			g := Game{tt.board}
			if got := g.GameWinner(); got != tt.winner {
				t.Errorf("Got %v as winner, but expected %v", got, tt.winner)
			}
		})
	}
}
