package main

type Entity rune

const (
	PLAYER = '>'
	EMPTY  = ' '
)

func (e Entity) IsPlayer() bool {
	return e == PLAYER
}

func (e Entity) IsEmpty() bool {
	return e == EMPTY
}

func (e Entity) String() string {
	return string(e)
}
