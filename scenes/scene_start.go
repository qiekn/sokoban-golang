package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type StartScene struct {
	isloaded bool
}

func NewStartScene() *StartScene {
	return &StartScene{
		isloaded: false,
	}
}

func (s *StartScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 0, 0, 255})
	ebitenutil.DebugPrint(screen, "Press enter to start")
}

func (s *StartScene) IsLoaded() bool { return s.isloaded }

func (s *StartScene) Start() { s.isloaded = true }

func (s *StartScene) Update() {}

func (s *StartScene) NextSceneId() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return GameSceneId
	}
	return StartSceneId
}

func (s *StartScene) OnEnter() {}

func (s *StartScene) OnExit() {}

var _ Scene = (*StartScene)(nil)
