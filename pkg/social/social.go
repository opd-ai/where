// Package social provides tribal council and alliance mechanics.
package social

import "github.com/opd-ai/where/pkg/engine"

// Alliance represents a group of allied players.
type Alliance struct {
	Name    string
	Members []engine.Entity
}

// Vote represents a tribal council vote.
type Vote struct {
	Voter  engine.Entity
	Target engine.Entity
}

// CouncilSystem manages tribal council voting and elimination.
type CouncilSystem struct {
	Alliances []Alliance
	Votes     []Vote
}

// NewCouncilSystem creates a new CouncilSystem.
func NewCouncilSystem() *CouncilSystem {
	return &CouncilSystem{}
}

// Update processes social mechanics each tick.
func (s *CouncilSystem) Update(world *engine.World, dt float64) {
	// Skeleton: tribal council and alliance logic
}

// CastVote records a vote from one player against another.
func (s *CouncilSystem) CastVote(voter, target engine.Entity) {
	s.Votes = append(s.Votes, Vote{Voter: voter, Target: target})
}

// TallyVotes counts votes and returns the entity with the most votes.
// Returns (0, false) when there are no votes.
func (s *CouncilSystem) TallyVotes() (engine.Entity, bool) {
	if len(s.Votes) == 0 {
		return 0, false
	}
	counts := make(map[engine.Entity]int)
	for _, v := range s.Votes {
		counts[v.Target]++
	}
	var maxEntity engine.Entity
	maxCount := 0
	for entity, count := range counts {
		if count > maxCount || (count == maxCount && entity < maxEntity) {
			maxCount = count
			maxEntity = entity
		}
	}
	return maxEntity, true
}
