package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/qiekn/components"
	"github.com/qiekn/managers"
)

func MoveInputUpdate() {
	dx, dy := 0, 0
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		dy = -1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		dy = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		dx = -1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		dx = 1
	}
	comps := managers.GetEntityManager().GetComponentsFromAll("MoveInput")
	for _, comp := range comps {
		if comp, ok := comp.(*components.MoveInput); ok {
			comp.Dx = dx
			comp.Dy = dy
		}
	}
	if dx != 0 || dy != 0 {
		MovePlayer(dx, dy)
	}
}
