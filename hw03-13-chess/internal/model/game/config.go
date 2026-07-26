package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Insert field size: ")
	fieldSize := defaultFieldSize
	if scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s != "" {
			if n, err := strconv.Atoi(s); err == nil && n >= minFieldSize && n <= maxFieldSize {
				fieldSize = n
			}
		}
	}

	fmt.Print("Insert first player name: ")
	player1 := defaultPlayer1
	if scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s != "" {
			player1 = s
		}
	}

	fmt.Print("Insert second player name: ")
	player2 := defaultPlayer2
	if scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s != "" {
			player2 = s
		}
	}

	return GameConfig{
		fieldSize: fieldSize,
		player1:   player1,
		player2:   player2,
	}
}
