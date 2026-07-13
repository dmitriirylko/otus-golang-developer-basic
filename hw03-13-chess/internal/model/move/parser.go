package move

import (
	"chess/internal/model/board"
	"errors"
	"strings"
)

// Ходы будем сериализовывать/десериализовывать следующим образом
// <color><pieceSan>-<from>-<to>
// color - цвет фигуры (b/w)
// pieceSan - SAN представление типа фигуры (N/B/R/Q/K)
// from - откуда ход в формате столбец-строка (например, a0)
// to - куда ходв формате столбец-строка (например, a1)

var (
	InvalidSeparation = errors.New("Invalid separation of string")
	InvalidFormat     = errors.New("Invalid format")
	InvalidColor      = errors.New("Invalid color")
	InvalidPiece      = errors.New("Invalid piece")
	InvalidSquare     = errors.New("Invalid square")
	OutOfBounds       = errors.New("Square is outside of the board")
)

func (m Move) String() string {
	var sb strings.Builder
	switch m.Color {
	case board.Black:
		sb.WriteRune('b')
	case board.White:
		sb.WriteRune('w')
	default:
		return ""
	}

	if m.Piece == board.NoneType {
		return ""
	}
	sb.WriteRune(m.Piece.SanLetter())
	sb.WriteRune('-')
	sb.WriteString(m.From.String())
	sb.WriteRune('-')
	sb.WriteString(m.To.String())

	return sb.String()
}

func parseMove(s string) (Move, error) {
	var m Move
	parts := strings.Split(s, "-")
	if len(parts) != 3 {
		return Move{}, InvalidSeparation
	}
	if len(parts[0]) != 2 || len(parts[1]) < 2 || len(parts[2]) < 2 {
		return Move{}, InvalidFormat
	}
	firstRunes := []rune(parts[0])
	switch firstRunes[0] {
	case 'b':
		m.Color = board.Black
	case 'w':
		m.Color = board.White
	default:
		return Move{}, InvalidColor
	}
	piece := board.SanLetterToPieceType(firstRunes[1])
	if piece == board.NoneType {
		return Move{}, InvalidPiece
	}
	m.Piece = piece

	sq, err := board.ParseSquare(parts[1])
	if err != nil {
		return Move{}, InvalidSquare
	}
	m.From = sq

	sq, err = board.ParseSquare(parts[2])
	if err != nil {
		return Move{}, InvalidSquare
	}
	m.To = sq

	return m, nil
}
