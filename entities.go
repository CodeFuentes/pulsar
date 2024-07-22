package main

type Entity rune

const (
	PLAYER = '>'
	EMPTY  = ' '

	BORDER_X = '─'
	BORDER_Y = '│'
)

func (e Entity) IsPlayer() bool {
	return e == PLAYER
}

func (e Entity) IsEmpty() bool {
	return e == EMPTY
}

func (e Entity) IsBorder() bool {
	return e == BORDER_X || e == BORDER_Y
}

func (e Entity) String() string {
	return string(e)
}
