package board

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	InvalidFormat = errors.New("Invalid format")
)

type Square struct {
	X int
	Y int
}

func (s Square) String() string {
	if s.X < 0 || s.Y < 0 || s.X > 25 || s.Y > 25 {
		return ""
	}
	file := string('a' + rune(s.X))
	rank := s.Y + 1
	return fmt.Sprintf("%s%d", file, rank)
}

func parseSqaure(s string) (Square, error) {
	var sq Square
	if len(s) < 2 {
		return Square{}, InvalidFormat
	}

	return sq, nil
}

type Board struct {
	size    int
	name1   string
	name2   string
	squares []*Piece
}

func NewBoard(sz int, name1, name2 string) *Board {
	if sz <= 0 || sz > 26 {
		return nil
	}
	return &Board{
		size:    sz,
		name1:   name1,
		name2:   name2,
		squares: make([]*Piece, sz*sz),
	}
}

func (b *Board) index(x, y int) int {
	return y*b.size + x
}

func (b *Board) inBounds(x, y int) bool {
	return x >= 0 && x < b.size && y >= 0 && y < b.size
}

func (b *Board) PieceAtCoord(x, y int) *Piece {
	if !b.inBounds(x, y) {
		return nil
	}
	return b.squares[b.index(x, y)]
}

func (b *Board) pieceAtSquare(sq Square) *Piece {
	return b.PieceAtCoord(sq.X, sq.Y)
}

func (b *Board) setToCoords(x, y int, p *Piece) {
	if !b.inBounds(x, y) {
		return
	}
	b.squares[b.index(x, y)] = p
}

func (b *Board) setToSquare(sq Square, p *Piece) {
	b.setToCoords(sq.X, sq.Y, p)
}

func LetterToColumn(r rune) int {

}

func ColumnLetter(x int) string {
	if x < 0 || x >= 26 {
		return ""
	}
	return string('a' + rune(x))
}

func (b *Board) Render() string {
	var sb strings.Builder
	// Верхняя строка с буквами
	// Отступ зависит от количества цифр в количестве строк доски
	indent := len(strconv.Itoa(b.size))
	for i := 0; i < indent; i++ {
		sb.WriteRune(' ')
	}
	for i := 0; i < b.size; i++ {
		sb.WriteRune(' ')
		sb.WriteRune('a' + rune(i))
	}
	sb.WriteRune('\n')

	// Отрисовка строк доски
	// Строка начинается с номера
	for y := b.size - 1; y >= 0; y-- {
		lineNumStr := fmt.Sprintf("%*d", indent, y+1)
		sb.WriteString(lineNumStr)
		for x := 0; x < b.size; x++ {
			sb.WriteRune(' ')
			p := b.PieceAtCoord(x, y)
			if p == nil {
				if (x+y)%2 == 0 {
					sb.WriteRune('#')
				} else {
					sb.WriteRune(' ')
				}
			} else {
				sb.WriteString(p.String())
			}
		}
		// Строка заканчивается номером
		sb.WriteRune(' ')
		sb.WriteString(lineNumStr)
		sb.WriteRune('\n')
	}

	// Нижняя строка с буквами
	for i := 0; i < indent; i++ {
		sb.WriteRune(' ')
	}
	for i := 0; i < b.size; i++ {
		sb.WriteRune(' ')
		sb.WriteRune('a' + rune(i))
	}
	sb.WriteString("\n\n")
	sb.WriteString(b.name1)
	sb.WriteRune('\n')
	sb.WriteString(b.name2)

	return sb.String()
}
