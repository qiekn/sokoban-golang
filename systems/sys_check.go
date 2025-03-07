package systems

import (
	"github.com/qiekn/managers"
)

func IsMovableAt(x, y int) bool {
	em := managers.GetEntityManager()

	// no collider
	if !em.HasComponentAt(x, y, "Collider") {
		return true
	}
	if em.HasComponentAt(x, y, "Movable") {
		return true
	}
	return false
}

func IsBoxAt(x, y int) bool {
	em := managers.GetEntityManager()
	id := em.GetFirstEntityAt(x, y)
	if id == 0 {
		return false
	}
	res := em.HasComponents(id, "Box")
	return res

}

func IsEntityBox(id managers.EntityId) bool {
	hasBox := managers.GetEntityManager().GetComponent(id, "Box")
	if hasBox != nil {
		return true
	}
	return false
}
