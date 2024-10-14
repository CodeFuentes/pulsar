package main

import (
	"time"
)

const (
	_DEFAULT_SPEED = 1500 * time.Millisecond
)

type PulsableGenerator interface {
	SetGenSpeed(genSpeed time.Duration)
	GetGenSpeed() time.Duration
}

type DefaultPulsableGenerator struct {
	PulsableGenerator

	genSpeed time.Duration
	board    Board
}

func NewDefaultGenerator(b Board) DefaultPulsableGenerator {
	return DefaultPulsableGenerator{
		genSpeed: _DEFAULT_SPEED,
		board:    b,
	}
}

func (g *DefaultPulsableGenerator) SetGenSpeed(genSpeed time.Duration) {
	g.genSpeed = genSpeed
}

func (g DefaultPulsableGenerator) GetGenSpeed() time.Duration {
	return g.genSpeed
}
