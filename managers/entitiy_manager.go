package managers

import (
	"sync"

	"github.com/qiekn/components"
)

type EntityId int

type EntityManager struct {
	entities     map[EntityId]bool
	emMap        map[EntityId]map[string]any
	nextEntityId EntityId
}

var (
	entittyManagerInstance *EntityManager
	entityManagerOnce      sync.Once
)

func GetEntityManager() *EntityManager {
	entityManagerOnce.Do(func() {
		entittyManagerInstance = &EntityManager{
			entities:     make(map[EntityId]bool),
			emMap:        make(map[EntityId]map[string]any),
			nextEntityId: 1, // 0 as an invalid id
		}
	})
	return entittyManagerInstance
}

////////////////////////////////////////////////////////////////////////
//                           Single Entity                            //
////////////////////////////////////////////////////////////////////////

func (em *EntityManager) CreateEntity() EntityId {
	id := em.nextEntityId
	em.entities[id] = true
	em.emMap[id] = make(map[string]any)
	em.nextEntityId++
	return id
}

func (em *EntityManager) DestroyEntity(entityId EntityId) {
	delete(em.entities, entityId)
	delete(em.emMap, entityId)
}

func (em *EntityManager) HasEntity(entityId EntityId) bool {
	// check if entityId exists
	_, exists := em.emMap[entityId]
	return exists
}

func (em *EntityManager) GetFirstEntityAt(x, y int) EntityId {
	ids := em.GetEntitiesWithComponents("Position")
	for _, id := range ids {
		pos, ok := em.GetComponent(id, "Position").(*components.Position)
		if ok && pos.X == x && pos.Y == y {
			return id
		}
	}
	return 0
}

func (em *EntityManager) GetEntitiesAt(x, y int) []EntityId {
	var res []EntityId
	ids := em.GetEntitiesWithComponents("Position")
	for _, id := range ids {
		pos, ok := em.GetComponent(id, "Position").(*components.Position)
		if ok && pos.X == x && pos.Y == y {
			res = append(res, id)
		}
	}
	return res
}

////////////////////////////////////////////////////////////////////////
//                     Single Entity Component(s)                     //
////////////////////////////////////////////////////////////////////////

func (em *EntityManager) AddComponent(entityId EntityId, name string, component any) {

	_, exists := em.entities[entityId]
	if !exists {
		return
	}

	em.emMap[entityId][name] = component
}

func (em *EntityManager) HasComponents(entityId EntityId, names ...string) bool {
	comps, exists := em.emMap[entityId]
	if !exists {
		return false
	}
	for _, name := range names {
		_, hasComp := comps[name]
		if !hasComp {
			return false
		}
	}
	return true
}

func (em *EntityManager) GetComponent(entityId EntityId, name string) any {
	if em.HasEntity(entityId) {
		comp := em.emMap[entityId][name]
		return comp
	}
	return nil
}

func (em *EntityManager) HasComponentAt(x, y int, name string) bool {
	ids := em.GetEntitiesAt(x, y)
	res := false
	for _, id := range ids {
		if em.HasComponents(id, name) {
			res = true
		}
	}
	return res
}

////////////////////////////////////////////////////////////////////////
//                     All Entities Component(s)                      //
////////////////////////////////////////////////////////////////////////

func (em *EntityManager) GetComponentsFromAll(name string) []any {
	var res []any
	for _, components := range em.emMap {
		component, exists := components[name]
		if exists {
			res = append(res, component)
		}
	}
	return res
}

func (em *EntityManager) GetEntitiesWithComponents(names ...string) []EntityId {

	var entities []EntityId
	for id, comps := range em.emMap {
		match := true
		for _, name := range names {
			if _, exists := comps[name]; !exists {
				match = false
				break
			}
		}
		if match {
			entities = append(entities, id)
		}
	}
	return entities
}

func (em *EntityManager) Clear() {
	em.entities = make(map[EntityId]bool)
	em.emMap = make(map[EntityId]map[string]any)
	em.nextEntityId = 1
}
