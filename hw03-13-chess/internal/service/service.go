package service

import (
	"chess/internal/model/board"
	"chess/internal/model/game"
	"chess/internal/model/move"
	"chess/internal/repository"
)

func Serve() {
	repository.Store(board.NewPiece(board.Pawn, board.White))
	repository.Store(board.Square{X: 1, Y: 1})
	repository.Store(move.NewMove(board.White, board.Pawn, board.Square{X: 0, Y: 0}, board.Square{X: 1, Y: 0}))
	repository.Store(game.Player{Name: "John"})
}
