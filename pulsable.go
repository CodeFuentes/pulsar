package main

import (
	"math/rand"
)

type Pulsable struct {
	body     string
	position Position
}

func NewPulsable(body string, position Position) *Pulsable {
	return &Pulsable{
		body,
		position,
	}
}

func NewRandomPulsable(position Position) *Pulsable {
	return NewPulsable(
		Pulsables[rand.Intn(len(Pulsables))],
		position,
	)
}

var Pulsables []string = []string{
	"bin",
	"mobile",
	"eternal",
	"roar",
	"brode",
	"cheap",
	"accumulation",
	"zero",
	"sport",
	"conviction",
	"jaw",
	"misplace",
	"profound",
	"grand",
	"hide",
}
