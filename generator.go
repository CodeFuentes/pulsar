package main

import (
	"time"
)

const (
	_DEFAULT_SPEED = 3000 * time.Millisecond
)

type IGenerator interface {
	SetBoard(board *Board)
	SetGenSpeed(genSpeed time.Duration)
	GetGenSpeed() time.Duration
}

type Generator struct {
	IGenerator

	genSpeed time.Duration
	board    *Board
}

func NewGenerator(speed time.Duration, board *Board) *Generator {
	return &Generator{
		genSpeed: speed,
		board:    board,
	}
}

func (g *Generator) SetBoard(board *Board) {
	g.board = board
}

func (g *Generator) SetGenSpeed(speed time.Duration) {
	g.genSpeed = speed
}

func (g *Generator) GetGenSpeed() time.Duration {
	return g.genSpeed
}
