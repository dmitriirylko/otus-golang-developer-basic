package board

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	InvalidFormat     = errors.New("Invalid format")
	InvalidConversion = errors.New("Invalid string to number conversion")
	InvalidSize       = errors.New("Invalid board size")
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

func ParseSquare(s string) (Square, error) {
	var sq Square
	if len(s) < 2 {
		return Square{}, InvalidFormat
	}
	sq.X = LetterToColumn(rune(s[0]))
	res, err := strconv.ParseInt(s[1:], 10, 32)
	if err != nil {
		return Square{}, InvalidConversion
	}
	sq.Y = int(res)
	return sq, nil
}

type Board struct {
	size    int
	name1   string
	name2   string
	squares []*Piece
}

func NewBoard(sz int, name1, name2 string) (Board, error) {
	if sz <= 0 || sz > 26 {
		return Board{}, InvalidSize
	}
	b := Board{
		size:    sz,
		name1:   name1,
		name2:   name2,
		squares: make([]*Piece, sz*sz),
	}
	b.setInitialPosition()
	return b, nil
}

func (b *Board) setInitialPosition() {
	standard := []PieceType{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}
	for i := 0; i < b.size && i < len(standard); i++ {
		if b.size > 2 {
			b.SetToCoords(i, 1, NewPiece(Pawn, White))
			b.SetToCoords(i, b.size-2, NewPiece(Pawn, Black))
		}
		b.SetToCoords(i, 0, NewPiece(standard[i], White))
		b.SetToCoords(i, b.size-1, NewPiece(standard[i], Black))
	}
}

func (b *Board) Index(x, y int) int {
	return y*b.size + x
}

func (b *Board) InBounds(x, y int) bool {
	return x >= 0 && x < b.size && y >= 0 && y < b.size
}

func (b *Board) PieceAtCoord(x, y int) *Piece {
	if !b.InBounds(x, y) {
		return nil
	}
	return b.squares[b.Index(x, y)]
}

func (b *Board) PieceAtSquare(sq Square) *Piece {
	return b.PieceAtCoord(sq.X, sq.Y)
}

func (b *Board) SetToCoords(x, y int, p *Piece) {
	if !b.InBounds(x, y) {
		return
	}
	b.squares[b.Index(x, y)] = p
}

func (b *Board) SetToSquare(sq Square, p *Piece) {
	b.SetToCoords(sq.X, sq.Y, p)
}

func LetterToColumn(r rune) int {
	if r < 'a' || r > 'z' {
		return -1
	}
	return int(r - 'a')
}

func ColumnToLetter(x int) string {
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
	sb.WriteRune('\n')

	return sb.String()
}
