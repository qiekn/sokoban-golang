package systems

import (
	"image/color"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/qiekn/components"
	"github.com/qiekn/constants"
	"github.com/qiekn/managers"
)

type renderEntity struct {
	id       managers.EntityId
	position *components.Position
	texture  *components.Texture
}

func Render(screen *ebiten.Image) {
	// background
	screen.Fill(color.RGBA{56, 63, 88, 255})

	// offset (to center level)
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()

	levelWidth := managers.GetLevelManager().GetCurrentLevel().Width
	levelHeight := managers.GetLevelManager().GetCurrentLevel().Height

	offsetX := (screenWidth - levelWidth*constants.Tilesize) / 2
	offsetY := (screenHeight - levelHeight*constants.Tilesize) / 2

	// get components
	em := managers.GetEntityManager()
	ids := em.GetEntitiesWithComponents("Texture", "Position")

	// get render entity
	renderEntities := make([]renderEntity, 0, len(ids))
	for _, id := range ids {
		position := em.GetComponent(id, "Position").(*components.Position)
		texture := em.GetComponent(id, "Texture").(*components.Texture)
		renderEntities = append(renderEntities, renderEntity{
			id:       id,
			position: position,
			texture:  texture,
		})
	}

	// sort
	sort.Slice(renderEntities, func(i, j int) bool {
		return renderEntities[i].texture.Order < renderEntities[j].texture.Order
	})

	// render
	for _, re := range renderEntities {
		opt := &ebiten.DrawImageOptions{}
		opt.GeoM.Translate(
			float64(re.position.X*constants.Tilesize),
			float64(re.position.Y*constants.Tilesize),
		)
		opt.GeoM.Translate(float64(offsetX), float64(offsetY))
		screen.DrawImage(managers.GetTextureManager().GetTexture(re.texture.Name), opt)
	}
}
