package entities

import "github.com/qiekn/components"

type Entity struct {
	Id         int
	Position   *components.Position
	Renderable *components.Renderable
}

type EntityManager struct {
	entities   []*Entity
	nextFreeId int
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities:   make([]*Entity, 0),
		nextFreeId: 0,
	}
}

func (m *EntityManager) AddEntity(e *Entity) {
	e.Id = m.nextFreeId // assign an Id
	m.nextFreeId++
	m.entities = append(m.entities, e)
}

func (m *EntityManager) GetAllEntities() []*Entity {
	return m.entities
}

func (m *EntityManager) GetEntitiesWithRenderable() []*Entity {
	var result []*Entity
	for _, e := range m.entities {
		if e.Renderable != nil {
			result = append(result, e)
		}
	}
	return result
}
