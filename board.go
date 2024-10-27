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
				p := NewPosition(rand.Intn(b.Height()-1)+1, b.Width()-1)
				pulsable := NewRandomPulsable(p)
				b.PulsableHandler(pulsable)
			case <-b.stopGen:
				break GEN
			}
		}
	}()
}

func (b *Board) PulsableHandler(p *Pulsable) {
	firstPrintedAll := false
	for i := 1; i < b.Width()-1; i++ {
		row, col := p.GetPosition().Row(), p.GetPosition().Col()
		printedAll := b.InsertPulsable(*p, p.GetPosition())
		p.SetPosition(NewPosition(row, col-1))
		if printedAll {
			if firstPrintedAll {
				endPos := NewPosition(row, col+p.Len())
				board.RemoveEntity(endPos)
			} else {
				firstPrintedAll = true
			}
		}
		time.Sleep(_PULSABLE_MOV_SPEED)
	}
}

func (b *Board) InsertPulsable(p Pulsable, pos Position) bool {
	for i, e := range p.GetBody() {
		ok := b.UpdateEntity(Entity(e), NewPosition(pos.Row(), pos.Col()+i))
		if !ok {
			return false
		}
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

func (b *Board) InsertEntity(e Entity, p Position) bool {
	if !b.EntityHasValidPosition(e) {
		return false
	}

	if b.GetEntity(p) != EMPTY {
		return false
	}

	b.entities[p.row][p.col] = e
	return true
}

func (b *Board) RemoveEntity(p Position) bool {
	if b.IsPositionEmpty(p) {
		return false
	}
	b.entities[p.row][p.col] = EMPTY
	return true
}

func (b *Board) UpdateEntity(e Entity, p Position) bool {
	if !b.IsPositionWithinBorders(p) {
		return false
	}
	if b.IsPositionEmpty(p) {
		return b.InsertEntity(e, p)
	}
	currentEntityPos, ok := b.GetEntityPosition(e)
	if ok {
		b.RemoveEntity(currentEntityPos)
	}
	b.entities[p.row][p.col] = e
	return true
}

func (b *Board) GetEntity(p Position) Entity {
	return b.entities[p.row][p.col]
}

func (b *Board) IsPositionEmpty(p Position) bool {
	return b.GetEntity(p).IsEmpty()
}

func (b *Board) IsPositionBorder(p Position) bool {
	return b.entities[p.row][p.col].IsBorder()
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
	b.player.position.row--
	if b.IsPositionBorder(b.player.position) {
		b.player.position.row = b.Height() - 1 - BORDER_WIDTH
	}
	return b
}

func (b *Board) MovePlayerDown() *Board {
	b.player.position.row++
	if b.IsPositionBorder(b.player.position) {
		b.player.position.row = 0 + BORDER_WIDTH
	}
	return b
}

func (b *Board) MovePlayerLeft() *Board {
	b.player.position.col--
	if b.IsPositionBorder(b.player.position) {
		b.player.position.col = b.Width() - 1 - BORDER_WIDTH
	}
	return b
}

func (b *Board) MovePlayerRight() *Board {
	b.player.position.col++
	if b.IsPositionBorder(b.player.position) {
		b.player.position.col = 0 + BORDER_WIDTH
	}
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
