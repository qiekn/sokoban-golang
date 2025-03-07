package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(960, 640)
	ebiten.SetWindowTitle("Sokoban")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	game := NewGame()
	game.InitManagers()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
