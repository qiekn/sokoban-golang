package systems

import (
	"fmt"

	"github.com/qiekn/components"
	"github.com/qiekn/managers"
)

func Win() bool {
	em := managers.GetEntityManager()
	targets := em.GetEntitiesWithComponents("Target")
	for _, target := range targets {
		// check if there are boxes at each target location
		if pos, ok := em.GetComponent(target, "Position").(*components.Position); ok {
			x, y := pos.X, pos.Y
			if !em.HasComponentsAt(x, y, "Box") {
				return false
			}
		}
	}
	fmt.Println("赢了了喵")
	return true
}
