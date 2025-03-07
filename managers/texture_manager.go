package managers

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type TextureManager struct {
	textures map[string]*ebiten.Image
}

var (
	textureManagerInstance *TextureManager
	textureManagerOnce     sync.Once
)

func GetTextureManager() *TextureManager {
	textureManagerOnce.Do(func() {
		textureManagerInstance = &TextureManager{
			textures: make(map[string]*ebiten.Image),
		}
		textureManagerInstance.loadTextures()
	})
	return textureManagerInstance
}

func (tm *TextureManager) loadTextures() {
	tm.textures["player"] = tm.loadTexture("assets/images/player.png")
	tm.textures["wall"] = tm.loadTexture("assets/images/wall.png")
	tm.textures["box"] = tm.loadTexture("assets/images/box.png")
	tm.textures["target"] = tm.loadTexture("assets/images/target.png")
}

func (tm *TextureManager) loadTexture(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func (tm *TextureManager) GetTexture(name string) *ebiten.Image {
	return tm.textures[name]
}
