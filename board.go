package main

const (
	MIN_WIDTH  = 20
	MIN_HEIGHT = 5
)

type Board struct {
	entities [][]Entity
	player   *Player
}

type Position struct {
	row int
	col int
}

func NewBoard(width, height int) *Board {
	width, height = max(width, MIN_WIDTH), max(height, MIN_HEIGHT)

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
	if !b.IsValidPosition(p) {
		return false
	}

	if b.GetEntity(p) != EMPTY {
		return false
	}

	b.entities[p.row][p.col] = e
	return true
}

func (b *Board) UpdateEntity(e Entity, p Position) bool {
	if !b.IsValidPosition(p) {
		return false
	}

	b.entities[p.row][p.col] = e
	return true
}

func (b *Board) GetEntity(p Position) Entity {
	if !b.IsValidPosition(p) {
		return EMPTY
	}
	return b.entities[p.row][p.col]
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
	if b.player.position.row < 0 {
		b.player.position.row = b.Height() - 1
	}
	return b
}

func (b *Board) MovePlayerDown() *Board {
	b.player.position.row++
	if b.player.position.row == b.Height() {
		b.player.position.row = 0
	}
	return b
}

func (b *Board) MovePlayerLeft() *Board {
	b.player.position.col--
	if b.player.position.col < 0 {
		b.player.position.col = b.Width() - 1
	}
	return b
}

func (b *Board) MovePlayerRight() *Board {
	b.player.position.col++
	if b.player.position.col == b.Width() {
		b.player.position.col = 0
	}
	return b
}

func (b *Board) IsValidPosition(p Position) bool {
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
