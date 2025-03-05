package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PauseScene struct {
	isloaded bool
}

func NewPauseScene() *PauseScene {
	return &PauseScene{
		isloaded: false,
	}
}

func (p *PauseScene) IsLoaded() bool { return p.isloaded }

func (p *PauseScene) OnEnter() {}

func (p *PauseScene) OnExit() { p.isloaded = false }

func (p *PauseScene) Start() { p.isloaded = true }

func (p *PauseScene) Update() {}

func (p *PauseScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 255, 0, 255})
	ebitenutil.DebugPrint(screen, "Pause Menu")
}

func (p *PauseScene) UpdateSceneId() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		return GameSceneId
	}
	return PauseSceneId
}

var _ Scene = (*PauseScene)(nil)
