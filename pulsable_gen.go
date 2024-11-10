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

func NewDefaultGenerator() *DefaultPulsableGenerator {
	return &DefaultPulsableGenerator{
		PulsableGenerator: NewPulsableGenerator(_DEFAULT_SPEED),
	}
}

func NewPulsableGenerator(s time.Duration) *PulsableGenerator {
	return &PulsableGenerator{
		NewGenerator(s),
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
