package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/qiekn/constants"
	"github.com/qiekn/scenes"
)

type Game struct {
	sm *scenes.SceneManager
}

func NewGame() *Game {
	return &Game{
		sm: scenes.NewSceneManager(),
	}
}

func (g *Game) Update() error {
	// scene switch update
	res := g.sm.Update()
	if res != nil {
		return res
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sm.DrawScene(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// return len(g.board[0]) * g.tileSize, len(g.board) * g.tileSize
	return constants.ScreenWidth, constants.ScreenHeight
}
