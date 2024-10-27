package main

type Position struct {
	row int
	col int
}

func NewPosition(row, col int) Position {
	return Position{row, col}
}

func (p Position) Row() int {
	return p.row
}

func (p Position) Col() int {
	return p.col
}

func (p Position) Equals(other Position) bool {
	return p.row == other.row && p.col == other.col
}
