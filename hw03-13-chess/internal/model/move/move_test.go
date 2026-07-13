package move_test

import (
	"chess/internal/model/board"
	"chess/internal/model/move"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveToString(t *testing.T) {
	cases := []struct {
		name      string
		color     board.Color
		pieceType board.PieceType
		from      board.Square
		to        board.Square
		expected  string
	}{
		{"1", board.White, board.Pawn, board.Square{X: 0, Y: 1}, board.Square{X: 0, Y: 2}, "wP-a2-a3"},
		{"2", board.White, board.Bishop, board.Square{X: 6, Y: 6}, board.Square{X: 4, Y: 4}, "wB-g7-e5"},
		{"3", board.Black, board.Rook, board.Square{X: 0, Y: 7}, board.Square{X: 3, Y: 7}, "bR-a8-d8"},
		{"4", board.Black, board.Pawn, board.Square{X: 6, Y: 4}, board.Square{X: 5, Y: 5}, "bP-g5-f6"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			move := move.NewMove(tc.color, tc.pieceType, tc.from, tc.to)
			str := move.String()
			assert.Equal(t, str, tc.expected)
		})
	}
}
