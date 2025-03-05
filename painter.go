package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Painter interface {
	Draw()
	Upload()
}

type GopherPainter struct {
	name string
}

func NewGopherPinter() *GopherPainter {
	return &GopherPainter{
		name: "Gopher",
	}
}

func (g *GopherPainter) Draw() {}

func (g *GopherPainter) Upload() ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("assets/images/gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	return *img
}
