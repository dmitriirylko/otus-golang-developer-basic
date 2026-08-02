package main

import (
	"bufio"
	"chess/internal/model/game"
	"chess/internal/model/move"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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
			fmt.Printf("\033[%dA", 1) // Поднимаем курсор вверх на 1 строчку (нужно после того как пользователь нажмет Enter)
			for i := 0; i < autoMoveCnt; i++ {
				mv, ok := g.RandomMove()
				if !ok {
					g.IsPlaying = false
					break
				}
				// Вывод номер автохода
				fmt.Printf("\033[K%s > automove...(%d/%d)\n", g.CurrentPlayer().Name, i+1, autoMoveCnt)
				// Задержка 2-4 секунды на каждом автоходе
				rng := rand.New(rand.NewSource(time.Now().UnixNano()))
				delay := 2000 + rng.Intn(2001)
				time.Sleep(time.Duration(delay) * time.Millisecond)
				// Применение автохода
				isApplied := g.ApplyMove(mv)
				if !isApplied {
					rerender(g, linesNum)
					continue
				}
				rerender(g, linesNum)
			}
			continue

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
