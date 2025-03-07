package systems

import (
	"github.com/qiekn/components"
	"github.com/qiekn/managers"
)

func MovePlayer(dx, dy int) {
	em := managers.GetEntityManager()
	ids := em.GetEntitiesWithComponents("MoveInput", "Position", "Movable")
	for _, id := range ids {
		pos := em.GetComponent(id, "Position").(*components.Position)
		if pos == nil {
			return
		}

		newX := pos.X + dx
		newY := pos.Y + dy

		// 1. without hitting collider, just go
		if !em.HasComponentsAt(newX, newY, "Collider") {
			pos.X, pos.Y = newX, newY
			return
		}
		// 2. hitting collider, try push
		if pushCollider(getColliderAt(newX, newY), dx, dy) {
			pos.X, pos.Y = newX, newY
		}
	}
}

////////////////////////////////////////////////////////////////////////
//                          Helper Functions                          //
////////////////////////////////////////////////////////////////////////

func getColliderAt(x, y int) managers.EntityId {
	em := managers.GetEntityManager()
	ids := em.GetEntitiesAt(x, y)
	for _, id := range ids {
		if em.HasComponents(id, "Collider") {
			return id
		}
	}
	return 0
}

func pushCollider(id managers.EntityId, dx, dy int) bool {
	em := managers.GetEntityManager()
	// collider is movable
	if checker := em.HasComponents(id, "Collider", "Movable"); checker {
		if pos, ok := em.GetComponent(id, "Position").(*components.Position); ok {
			// hit another collider (player can only push 1 box at once)
			if em.HasComponentsAt(pos.X+dx, pos.Y+dy, "Collider") {
				return false
			}
			pos.X += dx
			pos.Y += dy
			return true
		}
	}
	return false
}
