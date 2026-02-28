// Package config handles configuration loading using Viper.
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Camera perspective constants
const (
	PerspectiveFirstPerson     = "first-person"
	PerspectiveOverTheShoulder = "over-the-shoulder"
)

// Config holds all application configuration.
type Config struct {
	Game   GameConfig   `mapstructure:"game"`
	Window WindowConfig `mapstructure:"window"`
	Server ServerConfig `mapstructure:"server"`
	Audio  AudioConfig  `mapstructure:"audio"`
	Debug  DebugConfig  `mapstructure:"debug"`
}

// GameConfig holds game-related settings.
type GameConfig struct {
	Seed       int64  `mapstructure:"seed"`
	Genre      string `mapstructure:"genre"`
	MaxPlayers int    `mapstructure:"max_players"`
	MapSize    int    `mapstructure:"map_size"`
}

// WindowConfig holds window/display settings.
type WindowConfig struct {
	Title       string `mapstructure:"title"`
	Width       int    `mapstructure:"width"`
	Height      int    `mapstructure:"height"`
	Fullscreen  bool   `mapstructure:"fullscreen"`
	VSync       bool   `mapstructure:"vsync"`
	Perspective string `mapstructure:"perspective"`
}

// ServerConfig holds server settings.
type ServerConfig struct {
	Address  string `mapstructure:"address"`
	Port     int    `mapstructure:"port"`
	TickRate int    `mapstructure:"tick_rate"`
}

// AudioConfig holds audio settings.
type AudioConfig struct {
	Enabled bool    `mapstructure:"enabled"`
	Volume  float64 `mapstructure:"volume"`
}

// DebugConfig holds debug settings.
type DebugConfig struct {
	Enabled bool `mapstructure:"enabled"`
	ShowFPS bool `mapstructure:"show_fps"`
}

// Load reads configuration from file, returning a Config.
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	setDefaults()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("reading config: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshaling config: %w", err)
	}

	return &cfg, nil
}

func setDefaults() {
	viper.SetDefault("game.seed", 0)
	viper.SetDefault("game.genre", "fantasy")
	viper.SetDefault("game.max_players", 12)
	viper.SetDefault("game.map_size", 512)

	viper.SetDefault("window.title", "Where")
	viper.SetDefault("window.width", 1280)
	viper.SetDefault("window.height", 720)
	viper.SetDefault("window.fullscreen", false)
	viper.SetDefault("window.vsync", true)
	viper.SetDefault("window.perspective", PerspectiveFirstPerson)

	viper.SetDefault("server.address", "localhost")
	viper.SetDefault("server.port", 7777)
	viper.SetDefault("server.tick_rate", 60)

	viper.SetDefault("audio.enabled", true)
	viper.SetDefault("audio.volume", 0.8)

	viper.SetDefault("debug.enabled", false)
	viper.SetDefault("debug.show_fps", false)
}
