// Command client launches the Where game client.
package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/opd-ai/where/config"
	"github.com/opd-ai/where/pkg/engine"
	"github.com/opd-ai/where/pkg/rendering"
)

// Game implements the ebiten.Game interface.
type Game struct {
	cfg      *config.Config
	world    *engine.World
	renderer *rendering.Renderer
}

// Update advances the game state by one tick.
func (g *Game) Update() error {
	// Toggle perspective with 'P' key
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		currentPerspective := g.renderer.GetPerspective()
		var newPerspective string
		if currentPerspective == config.PerspectiveFirstPerson {
			newPerspective = config.PerspectiveOverTheShoulder
		} else {
			newPerspective = config.PerspectiveFirstPerson
		}
		g.renderer.SetPerspective(newPerspective)
		g.cfg.Window.Perspective = newPerspective
	}
	return nil
}

// Draw renders the current game state.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 34, G: 34, B: 34, A: 255})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Where — %s [seed %d]\nPerspective: %s (Press 'P' to toggle)",
		g.cfg.Game.Genre, g.cfg.Game.Seed, g.renderer.GetPerspective()))
}

// Layout returns the logical screen dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.cfg.Window.Width, g.cfg.Window.Height
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	world := engine.NewWorld()
	renderer := rendering.NewRenderer(cfg.Window.Width, cfg.Window.Height)
	renderer.SetPerspective(cfg.Window.Perspective)

	game := &Game{
		cfg:      cfg,
		world:    world,
		renderer: renderer,
	}

	ebiten.SetWindowSize(cfg.Window.Width, cfg.Window.Height)
	ebiten.SetWindowTitle(cfg.Window.Title)
	ebiten.SetVsyncEnabled(cfg.Window.VSync)

	if err := ebiten.RunGame(game); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
