package components

import "github.com/hajimehoshi/ebiten/v2"

type Position struct {
	X, Y int
}

type Movable struct{}

type Renderable struct {
	Char rune
}

type Sprite struct {
	Texture       *ebiten.Image
	Width, Height float64
}
