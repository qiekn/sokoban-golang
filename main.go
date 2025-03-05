package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Sokoban")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
