package main

import (
	"math/rand"
	"time"
)

var _PULSABLE_MOV_SPEED = time.Duration(1000 * time.Millisecond)

type Board struct {
	entities  [][]Entity
	player    *Player
	generator IGenerator
	stopGen   chan (struct{})
}

func NewBoard(width, height int) *Board {
	b := &Board{
		entities: make([][]Entity, height),
	}
	for i := range b.entities {
		b.entities[i] = NewEmptyBoardRow(width)
	}
	return b
}

func (b *Board) SetGenerator(gen IGenerator) {
	if b.generator != nil {
		b.stopGen <- struct{}{}
	}

	b.stopGen = make(chan struct{})
	b.generator = gen

	go func() {
	GEN:
		for {
			select {
			case <-time.After(b.generator.GetGenSpeed()):
				p := NewPosition(rand.Intn(b.Height()-1)+1, b.Width()-2)
				pulsable := NewRandomPulsable(p)
				go b.PulsableHandler(pulsable)
			case <-b.stopGen:
				// TODO: Clear previous' generator pulsables
				break GEN
			}
		}
	}()
}

func (b *Board) PulsableHandler(p *Pulsable) {
	firstPrintedAll := false
	for i := 1; i < b.Width()-1; i++ {
		row, col := p.GetPosition().Row(), p.GetPosition().Col()
		insertedAll := b.InsertWholePulsable(*p, p.GetPosition())
		p.SetPosition(NewPosition(row, col-1))
		if insertedAll {
			if firstPrintedAll {
				endPos := NewPosition(row, col+p.Len())
				b.entities[endPos.Row()][endPos.Col()] = EMPTY
			} else {
				firstPrintedAll = true
			}
		}
		time.Sleep(_PULSABLE_MOV_SPEED)
	}
}

func (b *Board) InsertWholePulsable(p Pulsable, pos Position) bool {
	for i := range p.GetBody() {
		if !b.IsPositionWithinBorders(pos) {
			return false
		}
		if b.entities[pos.Row()][pos.Col()] == EMPTY {
			b.entities[pos.Row()][pos.Col()] = Entity(p.GetBody()[i])
		}
	}
	if !b.IsPositionWithinBorders(NewPosition(pos.Row(), pos.Col()+p.Len()+1)) {
		return false
	}
	return true
}

func (b *Board) Width() int {
	if b.Height() == 0 {
		return 0
	}
	return len(b.entities[0])
}

func (b *Board) Height() int {
	return len(b.entities)
}

func (b *Board) GetEntity(p Position) Entity {
	return b.entities[p.row][p.col]
}

func (b *Board) IsPositionEmpty(p Position) bool {
	return b.GetEntity(p).IsEmpty()
}

func (b *Board) IsPositionBorder(p Position) bool {
	if p.Col() == 0 || p.Row() == 0 || p.Col() == b.Width()-1 || p.Row() == b.Height()-1 {
		return true
	}
	return false
}

func (b *Board) GetNewPulsableOrigin() Position {
	for {
		pos := NewPosition(len(b.entities)-2, rand.Intn(len(b.entities)-1))
		if b.entities[pos.Row()][pos.Col()] == EMPTY {
			return pos
		}
	}
}

func (b *Board) GetEntityPosition(e Entity) (Position, bool) {
	for row := range b.entities {
		for col := range b.entities[row] {
			if b.entities[row][col] == e {
				return Position{row, col}, true
			}
		}
	}
	return Position{}, false
}

/**
 * Player manipulation methods
 */
func (b *Board) SetPlayer(player *Player) bool {
	if b.player != nil {
		return false
	}
	b.player = player
	return true
}

func (b *Board) GetPlayer() (*Player, bool) {
	if b.player == nil {
		return nil, false
	}
	return b.player, true
}

func (b *Board) MovePlayerUp() *Board {
	if b.IsPositionBorder(b.player.position) {
		b.player.position.row = b.Height() - 1
		return b
	}
	b.player.position.row--
	return b
}

func (b *Board) MovePlayerDown() *Board {
	if b.IsPositionBorder(b.player.position) {
		b.player.position.row = 0
		return b
	}
	b.player.position.row++
	return b
}

func (b *Board) MovePlayerLeft() *Board {
	if b.IsPositionBorder(b.player.position) {
		b.player.position.col = b.Width() - 1
		return b
	}
	b.player.position.col--
	return b
}

func (b *Board) MovePlayerRight() *Board {
	if b.IsPositionBorder(b.player.position) {
		b.player.position.col = 0
		return b
	}
	b.player.position.col++
	return b
}

func (b *Board) EntityHasValidPosition(e Entity) bool {
	p, ok := b.GetEntityPosition(e)
	if !ok {
		return false
	}
	isWithinBorders := b.IsPositionWithinBorders(p)
	hasCollision := b.IsPositionEmpty(p)

	return isWithinBorders && !hasCollision
}

func (b *Board) IsPositionWithinBorders(p Position) bool {
	return p.col >= 0 && p.row >= 0 && b.Width() > 0 && p.col < b.Width() && p.row < b.Height()
}

func (b *Board) GetDefaultPlayerPosition() Position {
	return Position{
		col: b.Width() / 4,
		row: b.Height() / 2,
	}
}

func NewEmptyBoardRow(width int) []Entity {
	row := make([]Entity, width)
	for i := range row {
		row[i] = EMPTY
	}
	return row
}
