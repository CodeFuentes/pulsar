package main

type Board struct {
	entities [][]Entity
	player   *Player
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
	if !b.IsPositionEmpty(p) {
		return false
	}
	b.entities[p.row][p.col] = EMPTY
	return true
}

func (b *Board) UpdateEntity(e Entity, p Position) bool {
	if !b.IsPositionEmpty(p) {
		return false
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
	if b.entities[b.player.position.row][b.player.position.col].IsBorder() {
		b.player.position.row = b.Height() - 1 - BORDER_WIDTH
	}
	return b
}

func (b *Board) MovePlayerDown() *Board {
	b.player.position.row++
	if b.entities[b.player.position.row][b.player.position.col].IsBorder() {
		b.player.position.row = 0 + BORDER_WIDTH
	}
	return b
}

func (b *Board) MovePlayerLeft() *Board {
	b.player.position.col--
	if b.entities[b.player.position.row][b.player.position.col].IsBorder() {
		b.player.position.col = b.Width() - 1 - BORDER_WIDTH
	}
	return b
}

func (b *Board) MovePlayerRight() *Board {
	b.player.position.col++
	if b.entities[b.player.position.row][b.player.position.col].IsBorder() {
		b.player.position.col = 0 + BORDER_WIDTH
	}
	return b
}

func (b *Board) EntityHasValidPosition(e Entity) bool {
	p, ok := b.GetEntityPosition(e)
	if !ok {
		return false
	}
	isWithinBorders := p.col >= 0 && p.row >= 0 && b.Width() > 0 && p.col < b.Width() && p.row < b.Height()
	hasCollision := b.IsPositionEmpty(p)

	return isWithinBorders && !hasCollision
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
