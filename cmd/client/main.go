// Command client launches the Where game client.
package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/opd-ai/where/config"
	"github.com/opd-ai/where/pkg/engine"
)

// Game implements the ebiten.Game interface.
type Game struct {
	cfg   *config.Config
	world *engine.World
}

// Update advances the game state by one tick.
func (g *Game) Update() error {
	return nil
}

// Draw renders the current game state.
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 34, G: 34, B: 34, A: 255})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Where — %s [seed %d]", g.cfg.Game.Genre, g.cfg.Game.Seed))
}

// Layout returns the logical screen dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.cfg.Window.Width, g.cfg.Window.Height
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Printf("config: %v, using defaults", err)
		cfg = &config.Config{
			Game:   config.GameConfig{Genre: "fantasy", MaxPlayers: 12, MapSize: 512},
			Window: config.WindowConfig{Title: "Where", Width: 1280, Height: 720},
		}
	}

	world := engine.NewWorld()

	game := &Game{
		cfg:   cfg,
		world: world,
	}

	ebiten.SetWindowSize(cfg.Window.Width, cfg.Window.Height)
	ebiten.SetWindowTitle(cfg.Window.Title)
	ebiten.SetVsyncEnabled(cfg.Window.VSync)

	if err := ebiten.RunGame(game); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
