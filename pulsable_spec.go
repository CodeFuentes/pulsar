package main

import "math/rand"

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
