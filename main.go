package main

import (
	"fmt"
	"time"

	"github.com/codefuentes/gyro"
)

func main() {
	gameLoop := gyro.NewLoop().
		SetDebug(true).
		SetTargetFps(3).
		SetUpdateFunc(update)

	fmt.Println(gameLoop.GetTargetFps())
	gameLoop.Start()
}

func update(deltaTime time.Duration) {
	fmt.Println("delta", deltaTime.String())
}
