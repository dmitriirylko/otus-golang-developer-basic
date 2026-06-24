package move

import (
    "testing"
	"github.com/stretchr/testify/assert"
	"chess/internal/model/board"
)

func TestMoveCtors(t *testing.T) {
	m := NewMove(board.Square{0,0}, board.Square{1,1})
	assert.Equal(t, m.flag, FlagQuiet)
	assert.True(t, m.isQuiet())
	assert.False(t, m.isCapture())
	m1 := m.WithCapture()
	assert.Equal(t, m1.flag, FlagCapture)
	assert.False(t, m1.isQuiet())
	assert.True(t, m1.isCapture())
}

func TestMoveToSan(t *testing.T) {
	// TODO: увеличить количество тест-кейсов
	cases := []struct {
		name string
		from board.Square
		to board.Square
		pieceType board.PieceType
		withCapture bool
		expected string
	}{
		{"1", board.Square{0,1}, board.Square{0,2}, board.Pawn, false, "a3"},
		{"2", board.Square{6,6}, board.Square{4,4}, board.Bishop, false, "Be5"},
		{"3", board.Square{0,7}, board.Square{3,7}, board.Rook, true, "Rxd8"},
		{"4", board.Square{6,4}, board.Square{5,5}, board.Pawn, true, "gxf6"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T){
			piece := board.NewPiece(tc.pieceType, board.White)
			move := NewMove(tc.from, tc.to)
			if tc.withCapture == true { move = move.WithCapture() }
			san := move.ToSan(piece, nil)
			assert.Equal(t, san, tc.expected)
		})
	}
}