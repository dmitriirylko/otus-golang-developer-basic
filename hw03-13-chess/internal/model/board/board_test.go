package board_test

import (
	"chess/internal/model/board"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieceString(t *testing.T) {
	var piece board.Piece
	piece.Color = board.White
	piece.Type = board.Knight
	assert.Equal(t, piece.String(), "♘")

	piece.Color = board.Black
	piece.Type = board.Queen
	assert.Equal(t, piece.String(), "♛")

	piece.Color = board.Black
	piece.Type = board.NoneType
	assert.Equal(t, piece.String(), "?")
}

func TestSquareString(t *testing.T) {
	var square board.Square
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
	assert.Equal(t, square.String(), "")
}

func TestBoardLetterToColumn(t *testing.T) {
	assert.Equal(t, board.LetterToColumn('a'), 0)
	assert.Equal(t, board.LetterToColumn('z'), 25)
	assert.Equal(t, board.LetterToColumn('1'), -1)
}

func TestBoardIndex(t *testing.T) {
	b, err := board.NewBoard(10, "John", "Rayan")
	assert.Nil(t, err)
	assert.Equal(t, b.Index(5, 0), 5)
	assert.Equal(t, b.Index(2, 3), 32)
	assert.Equal(t, b.Index(9, 9), 99)
}

func TestBoardInBounds(t *testing.T) {
	b, err := board.NewBoard(10, "John", "Rayan")
	assert.Nil(t, err)
	assert.True(t, b.InBounds(0, 0))
	assert.True(t, b.InBounds(9, 9))
	assert.True(t, b.InBounds(4, 7))
	assert.False(t, b.InBounds(10, 1))
	assert.False(t, b.InBounds(10, 11))
	assert.False(t, b.InBounds(-1, 0))
	assert.False(t, b.InBounds(5, -1))
}

func TestBoardSetGet(t *testing.T) {
	b, err := board.NewBoard(10, "John", "Rayan")
	assert.Nil(t, err)
	p1 := board.NewPiece(board.Pawn, board.Black)
	b.SetToCoords(1, 1, p1)
	p2 := board.NewPiece(board.Knight, board.White)
	b.SetToSquare(board.Square{3, 8}, p2)
	assert.Nil(t, b.PieceAtCoord(-1, 1))
	assert.Nil(t, b.PieceAtCoord(1, 10))
	assert.Nil(t, b.PieceAtSquare(board.Square{2, 2}))
	p11 := b.PieceAtCoord(1, 1)
	assert.Equal(t, p1, p11)
	p22 := b.PieceAtSquare(board.Square{3, 8})
	assert.Equal(t, p2, p22)
}

func TestBoardRender(t *testing.T) {
	b, err := board.NewBoard(10, "John", "Rayan")
	assert.Nil(t, err)
	b.SetToCoords(2, 8, board.NewPiece(board.Rook, board.Black))
	b.SetToCoords(3, 2, board.NewPiece(board.King, board.White))
	fmt.Print(b.Render())
}
