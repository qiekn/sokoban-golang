package entities

import (
	"github.com/qiekn/components"
	"github.com/qiekn/managers"
)

func NewBox(x, y int) managers.EntityId {
	em := managers.GetEntityManager()
	id := em.CreateEntity()
	em.AddComponent(id, "Position", &components.Position{X: x, Y: y})
	em.AddComponent(id, "Texture", &components.Texture{Name: "box"})
	em.AddComponent(id, "Collider", &components.Collider{})
	em.AddComponent(id, "Movable", &components.Movable{})
	return id
}
