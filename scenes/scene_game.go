package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/qiekn/systems"
)

type GameScene struct {
	isloaded bool
}

func NewGameScene() *GameScene {
	gameScene := &GameScene{
		isloaded: false,
	}
	return gameScene
}

func (g *GameScene) IsLoaded() bool { return g.isloaded }

func (g *GameScene) OnEnter() {}

func (g *GameScene) OnExit() {}

func (g *GameScene) Start() {
	g.isloaded = true
	systems.InitCurrentLevel()
}

func (g *GameScene) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		systems.SwitchToNextLevel()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		systems.SwitchToPrevLevel()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		systems.Reset()
	}
	systems.MoveInputUpdate()
	if systems.Win() {
		systems.SwitchToNextLevel()
	}
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	systems.Render(screen)
}

func (g *GameScene) NextSceneId() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return ExitSceneId
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		return PauseSceneId
	}
	return GameSceneId
}

var _ Scene = (*GameScene)(nil)
