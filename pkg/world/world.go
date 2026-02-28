// Package world provides biome, climate, and erosion simulation.
package world

import (
	"fmt"

	"github.com/opd-ai/where/pkg/procgen"
)

// BiomeType identifies a biome.
type BiomeType string

const (
	BiomeForest    BiomeType = "forest"
	BiomeDesert    BiomeType = "desert"
	BiomeTundra    BiomeType = "tundra"
	BiomeSwamp     BiomeType = "swamp"
	BiomeMountain  BiomeType = "mountain"
)

// Tile represents a single map tile.
type Tile struct {
	Height      float64
	Temperature float64
	Humidity    float64
	Biome       BiomeType
}

// Map represents the game world map.
type Map struct {
	Width  int
	Height int
	Tiles  [][]Tile
}

// BiomeGenerator generates biome data from a seed.
type BiomeGenerator struct{}

// Generate produces a world map from the given seed and params.
func (g *BiomeGenerator) Generate(seed int64, params procgen.GenerationParams) (interface{}, error) {
	if params.Size <= 0 {
		return nil, fmt.Errorf("invalid map size: %d (must be > 0)", params.Size)
	}
	// Skeleton: biome generation logic
	m := &Map{
		Width:  params.Size,
		Height: params.Size,
		Tiles:  make([][]Tile, params.Size),
	}
	for i := range m.Tiles {
		m.Tiles[i] = make([]Tile, params.Size)
	}
	return m, nil
}

// Validate checks the generator configuration.
func (g *BiomeGenerator) Validate() error {
	return nil
}

// WeatherState represents a weather condition.
type WeatherState string

const (
	WeatherClear WeatherState = "clear"
	WeatherRain  WeatherState = "rain"
	WeatherStorm WeatherState = "storm"
	WeatherFog   WeatherState = "fog"
	WeatherSnow  WeatherState = "snow"
)

// WeatherGenerator generates weather state transitions.
type WeatherGenerator struct{}

// Generate produces a weather sequence from the given seed and params.
func (g *WeatherGenerator) Generate(seed int64, params procgen.GenerationParams) (interface{}, error) {
	// Skeleton: weather generation logic
	return WeatherClear, nil
}

// Validate checks the generator configuration.
func (g *WeatherGenerator) Validate() error {
	return nil
}

// EcosystemGenerator generates predator/prey population tables.
type EcosystemGenerator struct{}

// Generate produces ecosystem data from the given seed and params.
func (g *EcosystemGenerator) Generate(seed int64, params procgen.GenerationParams) (interface{}, error) {
	// Skeleton: ecosystem generation logic
	return nil, nil
}

// Validate checks the generator configuration.
func (g *EcosystemGenerator) Validate() error {
	return nil
}
