# Where

Procedurally generated wilderness survival game built with Go and Ebitengine.

## Directory Structure

```
where/
├── cmd/
│   ├── client/       # Game client entry point (Ebitengine window)
│   └── server/       # Dedicated server entry point
├── config/           # Configuration loading (Viper)
├── pkg/
│   ├── audio/        # Synthesized audio pipeline
│   ├── crafting/     # Procedural crafting system
│   ├── engine/       # ECS registry and game loop
│   ├── network/      # TCP server and client networking
│   ├── procgen/      # Procedural generation interface
│   ├── rendering/    # Runtime sprite/tile/particle rendering
│   ├── social/       # Tribal council and alliance mechanics
│   ├── survival/     # Hunger, thirst, temperature, fatigue systems
│   └── world/        # Biome, climate, erosion, weather, ecosystem
├── config.yaml       # Default configuration file
├── go.mod
├── ROADMAP.md
└── LICENSE
```

## Build

```
go build ./cmd/client
go build ./cmd/server
```

Ebitengine requires system libraries on Linux:

```
sudo apt-get install libx11-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev libxxf86vm-dev libgl1-mesa-dev libasound2-dev
```

## Run

```
# Client (opens game window)
./client

# Dedicated server
./server
```

## Configuration

Edit `config.yaml` in the working directory. Defaults are used if the file is absent.

| Section  | Fields                                  |
|----------|-----------------------------------------|
| game     | seed, genre, max_players, map_size      |
| window   | title, width, height, fullscreen, vsync |
| server   | address, port, tick_rate                |
| audio    | enabled, volume                         |
| debug    | enabled, show_fps                       |

## Dependencies

- [Go](https://go.dev/) 1.24+
- [Ebitengine](https://ebitengine.org/) v2 — game engine
- [Viper](https://github.com/spf13/viper) — configuration management
