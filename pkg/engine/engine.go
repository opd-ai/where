// Package engine provides the ECS registry and game loop.
package engine

// Entity is a unique identifier for an ECS entity.
type Entity uint64

// World is the ECS registry that manages entities and components.
type World struct {
	nextID      Entity
	Positions   map[Entity]*Position
	Hungers     map[Entity]*Hunger
	Inventories map[Entity]*Inventory
	Shelters    map[Entity]*Shelter
	Statuses    map[Entity]*StatusEffects
	Cameras     map[Entity]*Camera
}

// Position represents an entity's location in the world.
type Position struct {
	X, Y float64
}

// Hunger tracks an entity's hunger level.
type Hunger struct {
	Value float64
	Max   float64
}

// Inventory holds items for an entity.
type Inventory struct {
	Items    []string
	Capacity int
}

// Shelter represents an entity's shelter state.
type Shelter struct {
	Type      string
	Materials []string
	Integrity float64
}

// StatusEffects tracks active effects on an entity.
type StatusEffects struct {
	Effects map[string]float64
}

// Camera represents the camera perspective for an entity.
type Camera struct {
	Perspective string  // "first-person" or "over-the-shoulder"
	Distance    float64 // Distance from entity (used for over-the-shoulder)
	Angle       float64 // Camera angle in radians
}

// NewWorld creates and returns a new ECS World.
func NewWorld() *World {
	return &World{
		nextID:      1,
		Positions:   make(map[Entity]*Position),
		Hungers:     make(map[Entity]*Hunger),
		Inventories: make(map[Entity]*Inventory),
		Shelters:    make(map[Entity]*Shelter),
		Statuses:    make(map[Entity]*StatusEffects),
		Cameras:     make(map[Entity]*Camera),
	}
}

// NewEntity creates a new entity and returns its ID.
func (w *World) NewEntity() Entity {
	id := w.nextID
	w.nextID++
	return id
}

// System is the interface all ECS systems implement.
type System interface {
	Update(world *World, dt float64)
}
