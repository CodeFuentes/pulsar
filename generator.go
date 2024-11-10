package main

import (
	"time"
)

const (
	_DEFAULT_SPEED = 3 * time.Second
)

type IGenerator interface {
	SetGenSpeed(genSpeed time.Duration)
	GetGenSpeed() time.Duration
}

type Generator struct {
	IGenerator

	genSpeed time.Duration
}

func NewGenerator(speed time.Duration) *Generator {
	return &Generator{
		genSpeed: speed,
	}
}

func (g *Generator) SetGenSpeed(speed time.Duration) {
	g.genSpeed = speed
}

func (g *Generator) GetGenSpeed() time.Duration {
	return g.genSpeed
}
