package main

const (
	DEFAULT_PLAYER_SPEED = 1
)

type Player struct {
	entity   Entity
	position Position

	speed int
}

func NewPlayer(e Entity, p Position) *Player {
	return &Player{
		entity:   e,
		position: p,
		speed:    DEFAULT_PLAYER_SPEED,
	}
}

func (p *Player) GetEntity() Entity {
	return p.entity
}

func (p *Player) GetPosition() Position {
	return p.position
}

func (p *Player) SetPosition(position Position) *Player {
	p.position = position
	return p
}

func (p *Player) IsAt(position Position) bool {
	return position.row == p.position.row && position.col == p.position.col
}

func (p *Player) Shoot(e Entity) {
	ePos := NewPosition(p.position.row+1, p.position.col)
	_BOARD.entities[ePos.Row()][ePos.Col()] = e // TODO: Make this an event trigger
}
