package game

import (
	"chess/internal/model/board"
	"chess/internal/model/move"
)

type Game struct {
	IsPlaying bool
	Board     board.Board
	turn      board.Color
	history   []move.Move
	players   map[board.Color]Player
}

func NewGame(cfg GameConfig) (Game, error) {
	brd, err := board.NewBoard(cfg.fieldSize, cfg.player1, cfg.player2)
	if err != nil {
		return Game{}, err
	}
	g := Game{
		IsPlaying: true,
		Board:     brd,
		turn:      board.White,
		history:   make([]move.Move, 10),
		players:   make(map[board.Color]Player),
	}
	g.players[board.White] = Player{Name: cfg.player1}
	g.players[board.Black] = Player{Name: cfg.player2}
	return g, nil
}

func (g *Game) ApplyMove(mv move.Move) {
	if !mv.IsValid(&g.Board) {
		return
	}
	// Перестановка фигуры
	g.Board.SetToSquare(mv.To, g.Board.PieceAtSquare(mv.From))
	g.Board.SetToSquare(mv.From, nil)
	// Логгирование ходов
	g.history = append(g.history, mv)
	// Переход хода от белых к черным и наоборот
	g.turn = g.turn.Opposite()
}

func (g *Game) CurrentPlayer() Player {
	return g.players[g.turn]
}
