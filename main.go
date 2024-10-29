package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	WINDOW_TITLE = "PULSAR"

	FONT_SIZE_WIDTH  = 8
	FONT_SIZE_HEIGHT = 14
)

type Game struct{}

var _BOARD *Board
var b []rune
var lastUpdate time.Time

func (g *Game) Update() error {
	if time.Since(lastUpdate) > time.Millisecond*300 {
		player, hasPlayer := _BOARD.GetPlayer()
		if hasPlayer {
			_BOARD.UpdateEntity(player.GetEntity(), player.GetPosition())
		}

		b = []rune{}
		for row := range _BOARD.entities {
			for col := range _BOARD.entities[row] {
				entity := _BOARD.entities[row][col]
				b = append(b, rune(entity))
			}
			b = append(b, '\n')
		}
		lastUpdate = time.Now()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, string(b), 0, 0)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return _BOARD.Width() * FONT_SIZE_WIDTH, _BOARD.Height() * FONT_SIZE_HEIGHT
}

func main() {
	_BOARD = NewBoard(60, 30)
	// player := NewPlayer(PLAYER, _BOARD.GetDefaultPlayerPosition())
	player := NewPlayer(PLAYER, NewPosition(0, 0))
	_BOARD.SetPlayer(player)
	_BOARD.SetGenerator(NewDefaultGenerator(_BOARD)) // TODO: Refactor

	ebiten.SetWindowSize(1024, 512)
	ebiten.SetWindowTitle(WINDOW_TITLE)

	ebiten.SetTPS(30)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
		exit(err)
	}
}

/*
func input(event *tcell.EventKey) *tcell.EventKey {

		player, ok := _BOARD.GetPlayer()
		if !ok {
			panic("no player set")
		}

		currentPlayerPosition := player.GetPosition()
		_BOARD.RemoveEntity(currentPlayerPosition)
		switch event.Key() {
		case tcell.KeyUp:
			_BOARD.MovePlayerUp()
		case tcell.KeyDown:
			_BOARD.MovePlayerDown()
		case tcell.KeyLeft:
			_BOARD.MovePlayerLeft()
		case tcell.KeyRight:
			_BOARD.MovePlayerRight()
		case tcell.KeyEsc:
			game.Stop()
			exit(nil)
		case tcell.KeyRune:
			r := event.Rune()
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
				player.Shoot(Entity(r))
			}
		}

		return nil
	}
*/

func exit(err error) {
	if err != nil {
		fmt.Println(err)
		writeToFile(err.Error())
	} else {
		writeToFile("No error.")
	}
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
