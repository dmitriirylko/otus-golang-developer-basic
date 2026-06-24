package move

import (
	"strings"
	"fmt"
	"chess/internal/model/board"
)

type MoveFlag uint8

const (
	FlagQuiet MoveFlag = 0
	FlagCapture MoveFlag = 1 << iota
)

type Disambiguation struct {
	File bool
	Rank bool
}

type Move struct {
	from board.Square
	to board.Square
	flag MoveFlag
}

func NewMove(from, to board.Square) Move { return Move{from: from, to: to, flag: FlagQuiet} }

func (m Move) WithCapture() Move { m.flag = FlagCapture; return m }

func (m Move) isQuiet() bool { return m.flag == FlagQuiet }

func (m Move) isCapture() bool { return m.flag == FlagCapture }

func (m Move) ToSan(p *board.Piece, d *Disambiguation) string {
	if p == nil { return "???" }

	var sb strings.Builder

	if p.Type != board.Pawn { sb.WriteString(p.SanLetter()) }
	
	if d != nil {
		if d.File { sb.WriteString(board.ColumnLetter(m.from.X)) }
		if d.Rank { sb.WriteString(fmt.Sprintf("%d", m.from.Y+1)) }
	}

	if m.isCapture() {
		if p.Type == board.Pawn { sb.WriteString(board.ColumnLetter(m.from.X)) }
		sb.WriteString("x")
	}

	sb.WriteString(fmt.Sprintf("%s%d", board.ColumnLetter(m.to.X), m.to.Y+1))
	return sb.String()
}