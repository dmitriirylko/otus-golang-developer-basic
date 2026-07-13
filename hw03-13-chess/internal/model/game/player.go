package game

import (
	"chess/internal/model/move"
	"errors"
	"fmt"
)

var (
	InvalidUserInput = errors.New("Can not get user input from terminal")
	ParsingFail      = errors.New("Fail to parse user input")
)

type Player struct {
	name string
}

func (p Player) GetMove() (move.Move, error) {
	var userInput string
	fmt.Printf("%s: ", p.name)
	if _, err := fmt.Scan(&userInput); err != nil {
		return move.Move{}, InvalidUserInput
	}
	mv, er := move.ParseMove(userInput)
	if er != nil {
		return move.Move{}, ParsingFail
	}
	return mv, nil
}
