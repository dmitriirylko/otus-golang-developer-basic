package game

import (
	"chess/internal/model/board"
	"chess/internal/model/move"
)

type Game struct {
	board   board.Board
	turn    board.Color
	history []move.Move
	players map[board.Color]Player
}

func (g *Game) SetPlayer(color board.Color, p Player) {
	if _, ok := g.players[color]; !ok {
		g.players[color] = p
	}
}

func (g *Game) ApplyMove(mv move.Move) {
	// TODO
}
