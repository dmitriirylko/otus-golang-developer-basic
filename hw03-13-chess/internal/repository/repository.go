package repository

import (
	"chess/internal/model"
	"chess/internal/model/board"
	"chess/internal/model/game"
	"chess/internal/model/move"
)

var (
	Pieces  []*board.Piece
	Squares []board.Square
	Moves   []move.Move
	Players []game.Player
)

func Store(e model.Entity) {
	switch v := e.(type) {
	case *board.Piece:
		Pieces = append(Pieces, v)
	case board.Square:
		Squares = append(Squares, v)
	case move.Move:
		Moves = append(Moves, v)
	case game.Player:
		Players = append(Players, v)
	}
}
