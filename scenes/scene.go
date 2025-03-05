package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneId uint

const (
	StartSceneId SceneId = iota
	GameSceneId
	PauseSceneId
	ExitSceneId
)

type Scene interface {
	Start()
	Update()
	Draw(screen *ebiten.Image)
	OnEnter()
	OnExit()
	IsLoaded() bool
	UpdateSceneId() SceneId
}

type SceneManager struct {
	activeSceneId SceneId
	sceneMap      map[SceneId]Scene
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		activeSceneId: StartSceneId,
		sceneMap: map[SceneId]Scene{
			StartSceneId: NewStartScene(),
			GameSceneId:  NewGameScene(),
			PauseSceneId: NewPauseScene(),
		},
	}
}

func (m *SceneManager) Update() error {
	// update current scene
	activeScene := m.sceneMap[m.activeSceneId]
	if !activeScene.IsLoaded() {
		activeScene.Start()
		activeScene.OnEnter()
	}
	activeScene.Update()

	// check scene switch
	nextSceneId := activeScene.UpdateSceneId()
	if nextSceneId == ExitSceneId {
		activeScene.OnExit()
		return ebiten.Termination
	}
	if nextSceneId != m.activeSceneId {
		activeScene.OnExit()
		m.activeSceneId = nextSceneId
	}
	return nil
}

func (m *SceneManager) DrawScene(screen *ebiten.Image) {
	m.sceneMap[m.activeSceneId].Draw(screen)
}
