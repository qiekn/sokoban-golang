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
	NextSceneId() SceneId
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
	activeScene.Update()
	nextSceneId := activeScene.NextSceneId()

	// check scene switch
	if nextSceneId == ExitSceneId {
		activeScene.OnExit()
		return ebiten.Termination
	}
	if nextSceneId != m.activeSceneId {
		activeScene.OnExit()
		m.activeSceneId = nextSceneId
		nextScene := m.sceneMap[nextSceneId]
		if !nextScene.IsLoaded() {
			nextScene.Start()
		}
		nextScene.OnEnter()
	}
	return nil
}

func (m *SceneManager) DrawScene(screen *ebiten.Image) {
	m.sceneMap[m.activeSceneId].Draw(screen)
}
