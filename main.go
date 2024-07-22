package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/codefuentes/gyro"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	WINDOW_TITLE = "PULSAR"

	BORDER_WIDTH = 1
)

var game *tview.Application
var screen *tview.Table
var board *Board

func main() {
	recoverFunc := func(rvr any) { exit(fmt.Errorf("%v", rvr)) }
	defer func() {
		if rvr := recover(); rvr != nil {
			recoverFunc(rvr)
		}
	}()

	gameLoop := gyro.NewLoop().
		SetDebug(true).
		SetTargetFps(30).
		SetUpdateFunc(update).
		SetRenderFunc(render).
		SetRecoverFunc(recoverFunc)

	initCh := make(chan struct{})
	go initialize(initCh)

	// Initialization must finish before game loop
	// starts to avoid using nil pointers
	<-initCh
	err := gameLoop.Start()
	exit(err)
}

func exit(err error) {
	if err != nil {
		fmt.Println(err)
		writeToFile(err.Error())
	}
	writeToFile("No error.")
	fmt.Println("Press any button to exit...")
}

func writeToFile(msg string) {
	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	_, err = file.WriteString(msg)
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}
}

func initialize(done chan struct{}) {
	screen = tview.NewTable().
		SetBorders(false)

	screen.Box.SetBorder(false).SetTitle(WINDOW_TITLE)

	game = tview.NewApplication().
		SetRoot(screen, true).
		SetInputCapture(input)

	width, height := 50+BORDER_WIDTH, 25+BORDER_WIDTH
	board = NewBoard(width, height)

	for row := 0; row < board.Height(); row++ {
		endCol := board.Width() - 1
		for col := 0; col <= endCol; col++ {
			if col == 0 || col == endCol {
				board.entities[row][col] = BORDER_Y
			}
			if row == 0 || row == board.Height()-1 {
				board.entities[row][col] = BORDER_X
			}
		}
	}

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
	case tcell.KeyEsc:
		game.Stop()
		exit(nil)
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
				renderCellFromBoard(row, col)
			}
		}
	})
}

/*func renderCell(row, col int, entity Entity) {
	screen.SetCell(row, col, tview.NewTableCell(
		entity.String(),
	))
}*/

func renderCellFromBoard(row, col int) {
	screen.SetCell(row, col, tview.NewTableCell(
		string(board.GetEntity(Position{row, col})),
	))
}
