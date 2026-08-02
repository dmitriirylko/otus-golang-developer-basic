package main

import (
	"bufio"
	"chess/internal/model/game"
	"chess/internal/model/move"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	cfg := game.MakeConfig()
	g, err := game.NewGame(cfg)
	if err != nil {
		fmt.Println("ERROR: Can't start new game")
		return
	}

	linesNum := renderAndPrint(g)
	scanner := bufio.NewScanner(os.Stdin)

	for g.IsPlaying {
		player := g.CurrentPlayer()
		fmt.Printf("\033[K%s > ", player.Name)

		if !scanner.Scan() {
			break
		}
		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "" {
			rerender(g, linesNum)
			continue
		}

		parts := strings.Fields(userInput)

		switch parts[0] {
		case "surrender":
			g.IsPlaying = false

		case "automove":
			autoMoveCnt := 1
			if len(parts) > 1 {
				if n, err := strconv.Atoi(parts[1]); err == nil && n > 0 {
					autoMoveCnt = n
				}
			}
			for i := 0; i < autoMoveCnt; i++ {
			}

		default:
			mv, err := move.ParseMove(userInput)
			if err != nil {
				rerender(g, linesNum)
				continue
			}
			isApplied := g.ApplyMove(mv)
			if !isApplied {
				rerender(g, linesNum)
				continue
			}
		}
		rerender(g, linesNum)
	}
}

func renderAndPrint(g game.Game) int {
	s := g.Board.Render()
	fmt.Print(s)
	return strings.Count(s, "\n")
}

func rerender(g game.Game, lines int) {
	fmt.Printf("\033[%dA", lines+1)
	fmt.Print(g.Board.Render())
}
