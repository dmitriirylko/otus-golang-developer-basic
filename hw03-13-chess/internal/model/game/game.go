package game

import (
	"chess/internal/model/board"
	"chess/internal/model/move"
	"math/rand"
	"time"
)

type Game struct {
	IsPlaying bool
	Board     board.Board
	Turn      board.Color
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
		Turn:      board.White,
		history:   make([]move.Move, 10),
		players:   make(map[board.Color]Player),
	}
	g.players[board.White] = Player{Name: cfg.player1}
	g.players[board.Black] = Player{Name: cfg.player2}
	return g, nil
}

func (g *Game) ApplyMove(mv move.Move) bool {
	if !mv.IsValid(&g.Board, g.Turn) {
		return false
	}
	// Перестановка фигуры
	g.Board.SetToSquare(mv.To, g.Board.PieceAtSquare(mv.From))
	g.Board.SetToSquare(mv.From, nil)
	// Логгирование ходов
	g.history = append(g.history, mv)
	// Переход хода от белых к черным и наоборот
	g.Turn = g.Turn.Opposite()
	return true
}

func (g *Game) CurrentPlayer() Player {
	return g.players[g.Turn]
}

func (g *Game) RandomMove() (move.Move, bool) {
	size := g.Board.Size()

	// Найдем все ячейки, где находятся фигуры нужного цвета
	var sources []board.Square
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			p := g.Board.PieceAtCoord(x, y)
			if p != nil && p.Color == g.Turn {
				sources = append(sources, board.Square{X: x, Y: y})
			}
		}
	}
	if len(sources) == 0 {
		return move.Move{}, false
	}

	// Генерация случайного хода
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for attempts := 0; attempts < 200; attempts++ {
		from := sources[rng.Intn(len(sources))]
		to := board.Square{
			X: rng.Intn(size),
			Y: rng.Intn(size),
		}
		if from == to {
			continue
		}
		piece := g.Board.PieceAtSquare(from)
		mv := move.NewMove(g.Turn, piece.Type, from, to)
		if mv.IsValid(&g.Board, g.Turn) {
			return mv, true
		}
	}
	return move.Move{}, false
}
