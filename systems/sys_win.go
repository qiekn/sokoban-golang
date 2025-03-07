package systems

import "github.com/qiekn/managers"

func Win() bool {
	em := managers.GetEntityManager()
	targetPoints := em.GetEntitiesWithComponents("Target")
	var _ = targetPoints
	return false
}
