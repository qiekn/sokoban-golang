package main

type Entity struct {
	Id         int
	Position   *Position
	Renderable *Renderable
}

type Manager struct {
	entities   []*Entity
	nextFreeId int
}

func NewManager() *Manager {
	return &Manager{
		entities:   make([]*Entity, 0),
		nextFreeId: 0,
	}
}

func (m *Manager) AddEntity(e *Entity) {
	e.Id = m.nextFreeId // assign an Id
	m.nextFreeId++
	m.entities = append(m.entities, e)
}

func (m *Manager) GetEntitiesWithRenderable() []*Entity {
	var result []*Entity
	for _, e := range m.entities {
		if e.Renderable != nil {
			result = append(result, e)
		}
	}
	return result
}
