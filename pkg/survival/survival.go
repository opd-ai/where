// Package survival provides hunger, thirst, temperature, and fatigue systems.
package survival

import "github.com/opd-ai/where/pkg/engine"

// Stats holds the survival statistics for an entity.
type Stats struct {
	Hunger      float64
	Thirst      float64
	Temperature float64
	Fatigue     float64
}

// SurvivalSystem updates survival stats each tick.
type SurvivalSystem struct{}

// Update processes survival stat decay for all entities.
func (s *SurvivalSystem) Update(world *engine.World, dt float64) {
	// Skeleton: tick-based hunger, thirst, temperature, fatigue decay
}

// ForageSystem handles foraging actions.
type ForageSystem struct{}

// Update processes foraging for all entities.
func (s *ForageSystem) Update(world *engine.World, dt float64) {
	// Skeleton: forage system logic
}

// HuntSystem handles hunting actions.
type HuntSystem struct{}

// Update processes hunting for all entities.
func (s *HuntSystem) Update(world *engine.World, dt float64) {
	// Skeleton: hunt system logic
}

// ShelterSystem handles shelter building and maintenance.
type ShelterSystem struct{}

// Update processes shelter logic for all entities.
func (s *ShelterSystem) Update(world *engine.World, dt float64) {
	// Skeleton: shelter system logic
}
