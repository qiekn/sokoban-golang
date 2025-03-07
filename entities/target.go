package entities

import (
	"github.com/qiekn/components"
	"github.com/qiekn/managers"
)

func NewTarget(x, y int) managers.EntityId {
	em := managers.GetEntityManager()
	id := em.CreateEntity()
	em.AddComponent(id, "Position", &components.Position{X: x, Y: y})
	em.AddComponent(id, "Texture", &components.Texture{Name: "target", Order: 0})
	em.AddComponent(id, "Target", &components.Target{})
	return id
}
