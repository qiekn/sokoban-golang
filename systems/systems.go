package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/qiekn/components"
	"github.com/qiekn/constants"
)

func Move(entities []map[string]any, dx, dy int) {
	for _, entity := range entities {
		_, hasMovable := entity["Movable"].(*components.Movable)
		pos, hasPos := entity["Position"].(*components.Position)
		if hasPos && hasMovable {
			pos.X += dx
			pos.Y += dy
		}
	}
}

func Render(entities []map[string]any, screen *ebiten.Image) {
	/* TODO: Offset <2025-03-06 00:08, @qiekn> */
	offsetX := 0
	offsetY := 0

	for _, entity := range entities {
		pos, hasPos := entity["Position"].(*components.Position)
		sprite, hasSprite := entity["Sprite"].(*components.Sprite)

		if hasPos && hasSprite {
			opt := &ebiten.DrawImageOptions{}
			opt.GeoM.Translate(float64(pos.X*constants.Tilesize), float64(pos.Y*constants.Tilesize))
			opt.GeoM.Translate(float64(offsetX), float64(offsetY))
			screen.DrawImage(sprite.Texture, opt)
		}
	}
}

func Undo() {}

func Reset() {}

func History() {}
