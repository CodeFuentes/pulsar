package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	WINDOW_TITLE = "PULSAR"

	_FONT_SIZE_WIDTH  = 8
	_FONT_SIZE_HEIGHT = 14
)

type Game struct{}

var _BOARD *Board
var b string

func (g *Game) Update() error {
	err := input()
	if err != nil {
		return err
	}

	player, hasPlayer := _BOARD.GetPlayer()
	playerPos := player.GetPosition()
	if hasPlayer {
		_BOARD.entities[playerPos.row][playerPos.col] = PLAYER
	}

	b = ""
	for row := range _BOARD.entities {
		for col := range _BOARD.entities[row] {
			entity := _BOARD.entities[row][col]
			b += entity.String()
		}
		b += "\n"
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, b, 0, 0)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return _BOARD.Width() * _FONT_SIZE_WIDTH, _BOARD.Height() * _FONT_SIZE_HEIGHT
}

func main() {
	_BOARD = NewBoard(60, 30)
	player := NewPlayer(PLAYER, _BOARD.GetDefaultPlayerPosition())
	_BOARD.SetPlayer(player)
	_BOARD.SetGenerator(NewDefaultGenerator())

	ebiten.SetWindowSize(_BOARD.Width()*_FONT_SIZE_WIDTH, _BOARD.Height()*_FONT_SIZE_HEIGHT)
	ebiten.SetWindowTitle(WINDOW_TITLE)

	ebiten.SetTPS(30)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
		exit(err)
	}
}

func input() error {

	player, ok := _BOARD.GetPlayer()
	if !ok {
		panic("no player set")
	}

	inputChars := []rune{}
	ebiten.AppendInputChars(inputChars)
	for _, r := range inputChars {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			player.Shoot(Entity(r))
		}
	}

	removePlayer := func() {
		currentPlayerPosition := player.GetPosition()
		_BOARD.entities[currentPlayerPosition.row][currentPlayerPosition.col] = EMPTY
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		removePlayer()
		_BOARD.MovePlayerUp()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		removePlayer()
		_BOARD.MovePlayerDown()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		removePlayer()
		_BOARD.MovePlayerLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		removePlayer()
		_BOARD.MovePlayerRight()
	}
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("User pressed close button...")
	}

	return nil
}

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
