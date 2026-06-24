package move

import (
	"errors"
	"strings"
	"unicode"
	"chess/internal/model/board"
)

var (
	InvalidFormat = errors.New("Invalid format of Standard Algebraic Notation")
	AmbiguousSan  = errors.New("Multiple moves match")
	OutOfBounds   = errors.New("Square is outside of the board")
)

func Parse(s string, b *board.Board) (Move, error) {
	var pieceLetter rune
	var isCapture bool

	s = strings.TrimSpace(s)
	if len(s) > 0 && unicode.IsUpper(rune(s[0])) {
		pieceLetter = rune(s[0])
		s = s[1:]
	}

	if strings.Contains(s, "x") {
		isCapture = true
		s = strings.ReplaceAll(s, "x", "")
	}

	if len(s) < 2 { return Move{}, InvalidFormat }

	return Move{}, InvalidFormat
}