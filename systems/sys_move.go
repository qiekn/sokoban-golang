package systems

import (
	"github.com/qiekn/components"
	"github.com/qiekn/managers"
)

func MovePlayer(dx, dy int) {
	// find relevant position components
	em := managers.GetEntityManager()
	ids := em.GetEntitiesWithComponents("MoveInput", "Position", "Movable")
	for _, id := range ids {
		pos := em.GetComponent(id, "Position").(*components.Position)
		if pos == nil {
			return
		}

		newX := pos.X + dx
		newY := pos.Y + dy

		// 1. move without hitting collider
		if !em.HasComponentAt(newX, newY, "Collider") {
			pos.X, pos.Y = newX, newY
			return
		}
		// 2. can push obstacle away
		if PushEntity(em.GetEntityAt(newX, newY), dx, dy) {
			pos.X, pos.Y = newX, newY
		}
	}
}

// push without hitting any collider
func PushEntity(id managers.EntityId, dx, dy int) bool {
	em := managers.GetEntityManager()
	// this entity is movable
	if checker := em.HasComponents(id, "Collider", "Movable", "Position", "Texture"); checker {
		if pos, ok := em.GetComponent(id, "Position").(*components.Position); ok {
			// will hit any collider
			if em.HasComponentAt(pos.X+dx, pos.Y+dy, "Collider") {
				return false
			}
			pos.X += dx
			pos.Y += dy
			return true
		}
	}
	return false
}
