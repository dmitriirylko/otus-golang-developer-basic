package move

import (
	"chess/internal/model/board"
)

type Move struct {
	Color board.Color
	Piece board.PieceType
	From  board.Square
	To    board.Square
}

func NewMove(color board.Color, piece board.PieceType, from, to board.Square) Move {
	return Move{Color: color, Piece: piece, From: from, To: to}
}

func (m Move) EntityType() string { return "move" }

// TODO: проверка правил передвижения фигур
func (m Move) IsValid(b *board.Board) bool {
	if b.PieceAtSquare(m.From) == nil {
		return false
	}
	if !b.InBounds(m.To.X, m.To.Y) {
		return false
	}
	return true
}
