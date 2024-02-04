package main

type Player struct {
	entity   Entity
	position Position
}

func NewPlayer(e Entity, p Position) *Player {
	return &Player{e, p}
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
