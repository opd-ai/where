// Package rendering provides the runtime sprite, tile, and particle pipeline.
package rendering

import (
	"image/color"

	"github.com/opd-ai/where/config"
)

// Camera perspective constants (re-exported from config for convenience)
const (
	PerspectiveFirstPerson     = config.PerspectiveFirstPerson
	PerspectiveOverTheShoulder = config.PerspectiveOverTheShoulder
)

// Palette defines a genre-specific color palette for rendering.
type Palette struct {
	Genre  string
	Colors []color.RGBA
}

// TileSynthesizer generates terrain tiles at runtime.
type TileSynthesizer struct {
	Palette *Palette
}

// NewTileSynthesizer creates a new TileSynthesizer with the given palette.
func NewTileSynthesizer(palette *Palette) *TileSynthesizer {
	return &TileSynthesizer{Palette: palette}
}

// Renderer manages the rendering pipeline.
type Renderer struct {
	Width       int
	Height      int
	perspective string // "first-person" or "over-the-shoulder"
}

// NewRenderer creates and returns a new Renderer.
func NewRenderer(width, height int) *Renderer {
	return &Renderer{
		Width:       width,
		Height:      height,
		perspective: PerspectiveFirstPerson,
	}
}

// SetPerspective sets the camera perspective.
// Valid values are PerspectiveFirstPerson or PerspectiveOverTheShoulder.
// Invalid values are ignored.
func (r *Renderer) SetPerspective(perspective string) {
	if perspective == PerspectiveFirstPerson || perspective == PerspectiveOverTheShoulder {
		r.perspective = perspective
	}
}

// GetPerspective returns the current camera perspective.
func (r *Renderer) GetPerspective() string {
	return r.perspective
}

// Draw performs a render pass.
func (r *Renderer) Draw() {
	// Skeleton: rendering draw logic
	// Rendering will differ based on r.perspective
}
