package main

import (
	"time"

	"github.com/codefuentes/gyro"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	WINDOW_TITLE = "PULSAR"
)

var game *tview.Application
var screen *tview.Table
var board *Board

func main() {
	gameLoop := gyro.NewLoop().
		SetDebug(true).
		SetTargetFps(30).
		SetUpdateFunc(update).
		SetRenderFunc(render)

	initCh := make(chan struct{})
	go initialize(initCh)

	// Initialization must finish before game loop
	// starts to avoid using nil pointers
	<-initCh
	gameLoop.Start()
}

func initialize(done chan struct{}) {
	screen = tview.NewTable().
		SetBorders(false)

	screen.Box.SetBorder(true).SetTitle(WINDOW_TITLE)

	game = tview.NewApplication().
		SetRoot(screen, true).
		SetInputCapture(input)

	width, height := 80, 20
	board = NewBoard(width, height)
	player := NewPlayer(PLAYER, board.GetDefaultPlayerPosition())
	board.SetPlayer(player)

	// Done initializing
	close(done)

	if err := game.Run(); err != nil {
		panic(err)
	}
}

func input(event *tcell.EventKey) *tcell.EventKey {

	switch event.Key() {
	case tcell.KeyUp:
		board.MovePlayerUp()
	case tcell.KeyDown:
		board.MovePlayerDown()
	case tcell.KeyLeft:
		board.MovePlayerLeft()
	case tcell.KeyRight:
		board.MovePlayerRight()
	}

	return nil
}

func update(deltaTime time.Duration) {
	player, hasPlayer := board.GetPlayer()
	if hasPlayer {
		board.UpdateEntity(player.GetEntity(), player.GetPosition())
	}
}

func render() {
	game.QueueUpdateDraw(func() {
		for row := range board.entities {
			for col := range board.entities[row] {
				renderCell(row, col)
			}
		}
	})
}

func renderCell(row, col int) {
	screen.SetCell(row, col, tview.NewTableCell(
		string(board.GetEntity(Position{row, col})),
	))
}
