package board

import (
    "testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestPieceString(t *testing.T) {
	var piece Piece
	piece.Color = White
	piece.Type = Knight
	assert.Equal(t, piece.String(), "♘")

	piece.Color = Black
	piece.Type = Queen
	assert.Equal(t, piece.String(), "♛")

	piece.Color = Black
	piece.Type = None
	assert.Equal(t, piece.String(), "?")
}

func TestSquareString(t *testing.T) {
	var square Square
	square.X = 0
	square.Y = 0
	assert.Equal(t, square.String(), "a1")

	square.X = 7
	square.Y = 7
	assert.Equal(t, square.String(), "h8")

	square.X = 3
	square.Y = 5
	assert.Equal(t, square.String(), "d6")

	square.X = 10
	square.Y = 11
	assert.Equal(t, square.String(), "k12")

	square.X = 26
	square.Y = 1
	assert.Equal(t, square.String(), "(26,1)")
}

func TestBoardIndex(t *testing.T) {
	board := NewBoard(10)
	assert.NotNil(t, board)
	assert.Equal(t, board.index(5, 0), 5)
	assert.Equal(t, board.index(2, 3), 32)
	assert.Equal(t, board.index(9, 9), 99)
}

func TestBoardInBounds(t *testing.T) {
	board := NewBoard(10)
	assert.NotNil(t, board)
	assert.True(t, board.inBounds(0, 0))
	assert.True(t, board.inBounds(9, 9))
	assert.True(t, board.inBounds(4, 7))
	assert.False(t, board.inBounds(10, 1))
	assert.False(t, board.inBounds(10, 11))
	assert.False(t, board.inBounds(-1, 0))
	assert.False(t, board.inBounds(5, -1))
}

func TestBoardSetGet(t *testing.T) {
	board := NewBoard(10)
	assert.NotNil(t, board)
	p1 := NewPiece(Pawn, Black)
	board.setToCoords(1, 1, p1)
	p2 := NewPiece(Knight, White)
	board.setToSquare(Square{3, 8}, p2)
	assert.Nil(t, board.pieceAtCoord(-1, 1))
	assert.Nil(t, board.pieceAtCoord(1, 10))
	assert.Nil(t, board.pieceAtSquare(Square{2, 2}))
	p11 := board.pieceAtCoord(1, 1)
	assert.Equal(t, p1, p11)
	p22 := board.pieceAtSquare(Square{3, 8})
	assert.Equal(t, p2, p22)
}

func TestBoardRender(t *testing.T) {
	board := NewBoard(10)
	assert.NotNil(t, board)
	board.setToCoords(2, 8, NewPiece(Rook, Black))
	board.setToCoords(3, 2, NewPiece(King, White))
	fmt.Print(board.Render())
}