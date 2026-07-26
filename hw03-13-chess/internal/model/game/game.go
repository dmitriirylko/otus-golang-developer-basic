package game

import (
	"chess/internal/model/board"
	"chess/internal/model/move"
)

type Game struct {
	Board   board.Board
	turn    board.Color
	history []move.Move
	players map[board.Color]Player
}

func NewGame(cfg GameConfig) (Game, error) {
	brd, err := board.NewBoard(cfg.fieldSize, cfg.player1, cfg.player2)
	if err != nil {
		return Game{}, err
	}
	g := Game{
		Board:   brd,
		turn:    board.White,
		history: make([]move.Move, 10),
		players: make(map[board.Color]Player),
	}
	g.players[board.White] = Player{Name: cfg.player1}
	g.players[board.Black] = Player{Name: cfg.player2}
	return g, nil
}

func (g *Game) setPlayer(color board.Color, p Player) {
	if _, ok := g.players[color]; !ok {
		g.players[color] = p
	}
}

func (g *Game) ApplyMove(mv move.Move) {
	// TODO
}
