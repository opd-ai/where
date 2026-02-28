// Package rendering provides the runtime sprite, tile, and particle pipeline.
package rendering

import "image/color"

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
	Width  int
	Height int
}

// NewRenderer creates and returns a new Renderer.
func NewRenderer(width, height int) *Renderer {
	return &Renderer{
		Width:  width,
		Height: height,
	}
}

// Draw performs a render pass.
func (r *Renderer) Draw() {
	// Skeleton: rendering draw logic
}
