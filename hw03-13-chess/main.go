package main

import (
	"chess/internal/model/game"
	"fmt"
)

func main() {
	cfg := game.MakeConfig()
	g, err := game.NewGame(cfg)
	if err != nil {
		fmt.Println("ERROR: Can't start new game")
		return
	}
	fmt.Print(g.Board.Render())
}
