package main

import (
	"chess/chess"
	"fmt"
	"strconv"
	"strings"
)

type GameConfig struct {
	fieldSize int
	player1   string
	player2   string
}

func makeConfig() GameConfig {
	const (
		defaultFieldSize = 8
		minFieldSize     = 1
		maxFieldSize     = 26
		defaultPlayer1   = "Player1"
		defaultPlayer2   = "Player2"
	)
	var err error
	var fieldSize int
	var player1, player2, fieldSizeStr string
	fmt.Print("Insert field size: ")
	if _, err = fmt.Scan(&fieldSizeStr); err != nil {
		fieldSize = defaultFieldSize
	}
	fieldSize, err = strconv.Atoi(fieldSizeStr)
	if err != nil || fieldSize < minFieldSize || fieldSize > maxFieldSize {
		fieldSize = defaultFieldSize
	}
	fmt.Print("Insert first player name: ")
	if _, err = fmt.Scan(&player1); err != nil {
		player1 = defaultPlayer1
	}
	fmt.Print("Insert second player name: ")
	if _, err = fmt.Scan(&player2); err != nil {
		player2 = defaultPlayer2
	}
	return GameConfig{
		fieldSize: fieldSize,
		player1:   player1,
		player2:   player2,
	}
}

func makeField(pCfg *GameConfig) string {
	var strBuilder strings.Builder
	indent := len(strconv.Itoa(pCfg.fieldSize))
	for i := 0; i < indent; i++ {
		strBuilder.WriteRune(' ')
	}
	for i, r := 0, 'a'; i < pCfg.fieldSize; i, r = i+1, r+1 {
		strBuilder.WriteRune(r)
	}
	strBuilder.WriteByte('\n')
	pieceColor := chess.ColorEmpty
	for i := 0; i < pCfg.fieldSize; i++ {
		switch i {
		case 0, 1:
			pieceColor = chess.Black
		case 6, 7:
			pieceColor = chess.White
		default:
			pieceColor = chess.ColorEmpty
		}
		strBuilder.WriteString(fmt.Sprintf("%*d", indent, pCfg.fieldSize-i))
		var isBlack bool = i%2 == 1
		for j := 0; j < pCfg.fieldSize; j++ {
			piece := chess.PosColor2Piece(j, pieceColor)
			if piece != chess.PieceEmpty {
				strBuilder.WriteRune(piece.Rune())
			} else if isBlack {
				strBuilder.WriteRune('#')
			} else {
				strBuilder.WriteRune(' ')
			}
			isBlack = !isBlack
		}
		if i == 0 {
			strBuilder.WriteString(" " + pCfg.player1)
		}
		if i == 1 {
			strBuilder.WriteString(" " + pCfg.player2)
		}
		strBuilder.WriteRune('\n')
	}
	return strBuilder.String()
}

func main() {
	cfg := makeConfig()
	field := makeField(&cfg)
	fmt.Print(field)
}
