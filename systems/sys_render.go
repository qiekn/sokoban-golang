package systems

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/qiekn/components"
	"github.com/qiekn/constants"
	"github.com/qiekn/managers"
)

func Render(screen *ebiten.Image) {
	// background
	screen.Fill(color.RGBA{56, 63, 88, 255})

	// offset (to center level)
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()

	layerWidth := managers.GetLevelManager().GetCurrentLevel().Width
	layerHeight := managers.GetLevelManager().GetCurrentLevel().Height

	offsetX := (screenWidth - layerWidth*constants.Tilesize) / 2
	offsetY := (screenHeight - layerHeight*constants.Tilesize) / 2

	em := managers.GetEntityManager()
	ids := em.GetEntitiesWithComponents("Texture", "Position")

	for _, id := range ids {
		position := em.GetComponent(id, "Position").(*components.Position)
		texture := em.GetComponent(id, "Texture").(*components.Texture)

		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(float64(position.X*constants.Tilesize), float64(position.Y*constants.Tilesize))
		opt.GeoM.Translate(float64(offsetX), float64(offsetY))
		screen.DrawImage(managers.GetTextureManager().GetTexture(texture.Name), opt)
	}
}
