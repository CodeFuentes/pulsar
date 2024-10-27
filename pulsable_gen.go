package main

import (
	"math/rand"
	"time"
)

type PulsableGenerator struct {
	*Generator
}

type DefaultPulsableGenerator struct {
	*PulsableGenerator
}

func NewDefaultGenerator(b *Board) *DefaultPulsableGenerator {
	return &DefaultPulsableGenerator{
		PulsableGenerator: NewPulsableGenerator(_DEFAULT_SPEED, b),
	}
}

func NewPulsableGenerator(s time.Duration, b *Board) *PulsableGenerator {
	return &PulsableGenerator{
		NewGenerator(s, b),
	}
}

func (g *PulsableGenerator) SetGenSpeed(genSpeed time.Duration) {
	g.genSpeed = genSpeed
}

func (g PulsableGenerator) GetGenSpeed() time.Duration {
	return g.genSpeed
}

func NewRandomPulsable(position Position) *Pulsable {
	return NewPulsable(
		Pulsables[rand.Intn(len(Pulsables))],
		position,
	)
}
