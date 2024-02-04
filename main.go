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
	initCh := make(chan struct{})
	gameLoop := gyro.NewLoop().
		SetDebug(true).
		SetTargetFps(30).
		SetUpdateFunc(update).
		SetRenderFunc(render)

	go initialize(initCh)

	// Initialization must finish before game loop starts
	// to avoid using nil pointers
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
	}

	return nil
}

func update(deltaTime time.Duration) {
	screen.Box.SetTitle(string(deltaTime.Milliseconds()))
}

func render() {
	player, hasPlayer := board.GetPlayer()
	game.QueueUpdateDraw(func() {
		for row := range board.entities {
			for col := range board.entities[row] {
				position := Position{row, col}
				var cellValue string
				if hasPlayer && player.IsAt(position) {
					cellValue = string(player.GetEntity())
				} else {
					cellValue = string(board.GetEntity(position))
				}

				screen.SetCell(row, col, tview.NewTableCell(cellValue))
			}
		}
	})
}
