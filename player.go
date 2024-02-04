package main

const (
	DEFAULT_PLAYER_SPEED = 10
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
