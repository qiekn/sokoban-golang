package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/qiekn/constants"
	"github.com/qiekn/scenes"
)

type Game struct {
	sceneManager *scenes.SceneManager
}

func NewGame() *Game {
	return &Game{
		sceneManager: scenes.NewSceneManager(),
	}
}

func (g *Game) Update() error {
	res := g.sceneManager.Update()
	if res == ebiten.Termination {
		return res
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.DrawScene(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// return len(g.board[0]) * g.tileSize, len(g.board) * g.tileSize
	return constants.ScreenWidth, constants.ScreenHeight
}
