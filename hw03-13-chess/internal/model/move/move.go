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
func (m Move) IsValid(b *board.Board, turn board.Color) bool {
	// Проверка -- чей ход сейчас (черных или белых)
	if turn != m.Color {
		return false
	}
	// На ячейке, откуда осуществляется ход должа быть фигура
	if b.PieceAtSquare(m.From) == nil {
		return false
	}
	// Ячейка, куда осуществляется ход должна быть в границах доски
	if !b.InBounds(m.To.X, m.To.Y) {
		return false
	}
	// Ячейка, куда осуществляется ход не должна содержать фигуру своего же цвета
	toPiece := b.PieceAtSquare(m.To)
	if toPiece != nil {
		if toPiece.Color == turn {
			return false
		}
	}
	return true
}
