package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Images struct {
	Player *ebiten.Image
	Wall   *ebiten.Image
	Box    *ebiten.Image
	Point  *ebiten.Image
}

func loadImages() (*Images, error) {
	player, _, err := ebitenutil.NewImageFromFile("assets/images/player.png")
	if err != nil {
		return nil, err
	}
	wall, _, err := ebitenutil.NewImageFromFile("assets/images/wall.png")
	if err != nil {
		return nil, err
	}
	box, _, err := ebitenutil.NewImageFromFile("assets/images/box.png")
	if err != nil {
		return nil, err
	}
	point, _, err := ebitenutil.NewImageFromFile("assets/images/point.png")
	if err != nil {
		return nil, err
	}
	return &Images{
		Player: player,
		Wall:   wall,
		Box:    box,
		Point:  point,
	}, nil
}
