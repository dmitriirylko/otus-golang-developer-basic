package main

import (
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
		maxFieldSize     = 100
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

func makeField(size int) string {
	strLen := size*size + size
	var strBuilder strings.Builder
	strBuilder.Grow(strLen)
	for i := 0; i < size; i++ {
		var isBlack bool = i%2 == 1
		for j := 0; j < size; j++ {
			if isBlack {
				strBuilder.WriteByte('#')
			} else {
				strBuilder.WriteByte(' ')
			}
			isBlack = !isBlack
		}
		strBuilder.WriteByte('\n')
	}
	return strBuilder.String()
}

func main() {
	cfg := makeConfig()
	field := makeField(cfg.fieldSize)
	fmt.Print(field)
}
