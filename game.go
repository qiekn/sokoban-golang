package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{141, 65, 65, 255})

	/* My Avator
	img, _, err := ebitenutil.NewImageFromFile("assets/images/avatar.png")
	if err != nil {
		log.Fatal(err)
	}
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(0.25, 0.25)
	opt.GeoM.Translate(100, 50)
	screen.DrawImage(img, opt)
	*/

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 320, 240
}
