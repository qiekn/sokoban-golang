package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/qiekn/constants"
	"github.com/qiekn/managers"
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

func (g *Game) InitManagers() {
	managers.GetLevelManager()
	managers.GetEntityManager()
	managers.GetTextureManager()
}

func (g *Game) Update() error {
	// scene update
	if err := g.sm.Update(); err != nil {
		return err
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sm.DrawScene(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.ScreenWidth, constants.ScreenHeight
}
