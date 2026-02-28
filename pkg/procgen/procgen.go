// Package procgen provides the procedural generation interface and utilities.
package procgen

// GenerationParams holds parameters for procedural generation.
type GenerationParams struct {
	Genre string
	Size  int
}

// Generator is the interface for all procedural generators.
type Generator interface {
	Generate(seed int64, params GenerationParams) (interface{}, error)
	Validate() error
}
