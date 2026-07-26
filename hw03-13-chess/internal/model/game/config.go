package game

import (
	"fmt"
	"strconv"
)

type GameConfig struct {
	fieldSize int
	player1   string
	player2   string
}

func MakeConfig() GameConfig {
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
