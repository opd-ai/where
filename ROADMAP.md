# Where — Implementation Roadmap

## 1. Project Overview

**Where** is a 100% procedurally generated wilderness survival sim built in Go 1.24+ with the Ebiten v2 engine. Inspired by *Alone*, *Survivor*, and *Extracted*, it drops 2–12 players into a procedurally generated wilderness where they must forage, hunt, craft, build shelter, and outlast rivals through a tribal-council elimination mechanic. Every asset — terrain, flora, fauna, audio, visuals — is generated at runtime from a deterministic seed; no files are bundled. Where extends the V-Series (`opd-ai/venture`, `opd-ai/vania`, `opd-ai/violence`, `opd-ai/velocity`) with biome/climate simulation, ecosystem food chains, dynamic day/night and weather cycles, and complex social gameplay (alliances, betrayals, vote-outs) across five distinct genre skins: **fantasy, sci-fi, horror, cyberpunk, post-apocalyptic**.

---

## 2. Core Architecture

**ECS:** Entities are `uint64` IDs. Components are plain structs (`Position`, `Hunger`, `Inventory`, `Shelter`, `StatusEffects`). Systems own all logic and iterate component sets each tick.

**V-Series reuse:**
- `pkg/procgen` — `Generator` interface (`Generate(seed int64, params GenerationParams) (interface{}, error)` + `Validate()`); extend with biome, flora, fauna, weather generators.
- `pkg/audio` — oscillator/envelope pipeline; add biome ambience, weather layers, animal-call synthesis.
- `pkg/rendering` — runtime sprite/tile/particle pipeline; add terrain tile blending, post-processing palettes per genre.
- `pkg/network` — TCP authoritative server, client-side prediction, delta compression; extend for open-world chunk sync and elimination protocol.
- `pkg/engine` — game loop, ECS registry, input; reuse unchanged.

**New packages:** `pkg/world` (biome/climate/erosion), `pkg/survival` (hunger/thirst/temp/fatigue), `pkg/crafting`, `pkg/social` (tribal council, alliances).

---

## 3. Implementation Phases

### Phase 1 — Foundation & World Generation (6 weeks)
1. Scaffold Go module, `cmd/client`, `cmd/server`, package skeletons matching V-Series layout.
   - **AC:** `go build ./...` succeeds; all packages importable.
2. Implement `BiomeGenerator`: climate-lattice → temperature/humidity → biome assignment, Diamond-Square erosion for terrain heightmap.
   - **AC:** 1000 seeds produce valid, varied heightmaps; same seed always identical.
3. Implement `WeatherGenerator`: Markov-chain day/night + weather state machine (clear/rain/storm/fog/snow).
   - **AC:** 24-hour simulated cycle completes without state deadlock; all five weather states reachable.
4. Implement genre-keyed visual palettes and terrain tile synthesizer; all five `GenreID` values produce visually distinct output.
   - **AC:** Screenshot diff tool confirms ≥40% pixel divergence between any two genre renders of the same seed.
5. Implement `EcosystemGenerator`: predator/prey population tables, food-chain rules, per-biome spawn rates.
   - **AC:** Ecosystem reaches stable equilibrium within 100 simulated in-game days; no species drives another to zero within 30 days.

### Phase 2 — Survival Systems (5 weeks)
6. Implement `SurvivalSystem`: hunger, thirst, body temperature, fatigue components; tick-based decay; death/incapacitation thresholds.
   - **AC:** Unit tests verify each stat reaches critical threshold within expected tick range under default parameters.
7. Implement `ForageSystem` and `HuntSystem`: genre-themed flora/fauna items procedurally generated; `CraftingSystem` with recipe graph derived from seed+genre.
   - **AC:** All five genres produce distinct item/recipe names; no two genres share >20% of recipe identifiers.
8. Implement `ShelterSystem`: procedurally generated shelter blueprints (material lists, stat bonuses) per genre/biome.
   - **AC:** Each genre×biome pair yields a distinct shelter type name and ≥1 unique material.
9. Implement genre-keyed audio: biome ambiences, weather layers, animal calls, crafting SFX — all synthesized via `pkg/audio` oscillator pipeline with per-genre pitch/timbre modifiers from V-Series spec.
   - **AC:** Audio output for horror genre has pitch factor ≤0.7× baseline; sci-fi ≥1.3×; cyberpunk exhibits hard-clipping waveform distortion.

### Phase 3 — Multiplayer Core (5 weeks)
10. Authoritative server: chunk-based world-state delta compression; inventory/crafting authority on server.
    - **AC:** Two clients on 200ms simulated latency maintain <50ms positional divergence.
11. Client-side prediction + server reconciliation for player movement and action inputs.
    - **AC:** On 5000ms simulated latency, client remains playable (no freeze); reconciliation corrects position within 3 ticks of packet receipt.
12. Spectator mode: eliminated players receive a read-only stream of world state.
    - **AC:** Spectator client connects post-elimination and renders live game state without write authority.

### Phase 4 — Social & Elimination Mechanics (4 weeks)
13. Tribal council protocol: server collects votes from all active players; majority vote eliminates one player per council round; tie-breaker by survival score.
    - **AC:** Automated integration test: 6 bots vote; correct player eliminated; spectator stream activates for eliminated player within 500ms.
14. Alliance/betrayal system: players form named alliances (server-authoritative); alliance membership affects council vote weighting.
    - **AC:** Alliance formation and dissolution round-trip latency <200ms on localhost; betrayal logs persisted in match record.
15. Audience/spectator voting: spectators cast votes that influence in-game events (weather shift, item drop, hazard spawn).
    - **AC:** Spectator vote tally processed server-side within one 5-second voting window; triggered event visible to all active clients.

### Phase 5 — Polish, Balancing & Hardening (3 weeks)
16. Determinism audit: run 100 seeds × 5 genres; assert identical output for identical inputs across OS/arch.
    - **AC:** CI matrix (linux/amd64, darwin/arm64, windows/amd64) produces byte-identical world state checksums.
17. Performance profiling: world generation <2s for 512×512 map; server tick <16ms for 12 players; audio synthesis <5ms per frame.
    - **AC:** `go test -bench` results meet above thresholds on reference hardware (4-core, 8 GB RAM).
18. Single-binary release: `go build -tags production ./cmd/server` and `./cmd/client` produce standalone executables with zero runtime asset dependencies.
    - **AC:** Binary runs on clean VM with no additional files; `ldd` shows no missing shared libs (or static build confirmed).

---

## 4. PCG Systems Inventory

| Generator | Algorithm/Approach | Fantasy | Sci-Fi | Horror | Cyberpunk | Post-Apoc |
|---|---|---|---|---|---|---|
| `BiomeGenerator` | Climate lattice + Diamond-Square erosion | Enchanted forest/moor | Alien mesa/crater | Fog marsh/dead forest | Toxic sprawl/ruin | Dust plain/rubble field |
| `WeatherGenerator` | Markov-chain state machine | Magical storms, fairy mist | Ion storms, zero-g dust | Bone rain, necrotic fog | Acid rain, neon smog | Fallout squalls, ash blizzard |
| `FloraGenerator` | L-system + genre palette | Glowing mushrooms, elder oak | Xenofern, bioluminescent lichen | Corpse flower, thorn vine | Mutant kudzu, rust weed | Scrub brush, irradiated cactus |
| `FaunaGenerator` | Population + behavior tables | Dire wolf, fae deer, gryphon | Xenobeast, drone swarm | Undead boar, wraith stag | Rad-rat, cyber hound | Feral hog, dust vulture |
| `CraftingGenerator` | Recipe graph seeded by biome+genre | Enchanted bow, rune trap | Plasma torch, force barrier | Bone knife, ward fetish | Circuit snare, EMP grenade | Scrap spear, salvage net |
| `ShelterGenerator` | Blueprint graph, genre-keyed materials | Wattle longhouse, stone circle | Prefab module, force-field dome | Bone-frame lean-to, crypt alcove | Corrugated bunker, hack-shack | Salvage shanty, buried dugout |
| `EcosystemGenerator` | Lotka-Volterra + food-chain rules | Magical predator/prey balance | Xenoecology cycles | Undead resurgence mechanic | Scavenger dominance | Radiation mutation rate |
| `AudioGenerator` | Oscillators + envelopes + motifs | Orchestral strings, nature SFX | Synth pads +30% pitch | Minor dissonance −30% pitch + vibrato | Hard-clipped synth +40% pitch | Detuned acoustic, wind noise |
| `PaletteGenerator` | HSL shift + post-processing LUT | Warm greens/golds | Cold blues/teals | Desaturated grey/green | Neon pink/cyan on black | Sepia/ochre + film grain |
| `EventGenerator` | Weighted random + narrative templates | Fae bargain, monster ambush | Alien signal, equipment malfunction | Haunting, curse spread | Corp drone strike, data breach | Raider attack, rad surge |

---

## 5. Multiplayer Design

**Authority model:** Server owns all world state; clients send inputs only. Crafting, foraging, and council votes are validated server-side before applying.

**Chunk sync:** 32×32-tile chunks; server sends delta-compressed chunk updates only for chunks within each client's view radius. Initial load: full chunk; subsequent ticks: changed tiles only.

**Prediction & reconciliation:** Client predicts movement locally; server reconciles every authoritative tick. On divergence, client snaps to server position and replays buffered inputs (V-Series pattern extended for 2D wilderness).

**Lag compensation:** Server maintains 1-second input history. Melee/ranged actions rewound to sender's local time for hit detection.

**Elimination protocol:** On council vote finalization, server broadcasts `EliminationEvent{PlayerID, Reason, SurvivalScore}`. Eliminated player's authority stripped; client transitions to spectator stream. Match ends when one player or one alliance reaches configurable win threshold.

**Spectator feed:** Separate lower-bandwidth stream (position + animation only, no inventory). Spectators receive `AudienceVoteWindow` events to cast influence votes.

**Tor/onion tolerance:** All protocol timers and vote windows parameterized for 200–5000ms RTT. No real-time-critical gating; council votes use a configurable 30-second window.

---

## 6. Genre Differentiation Matrix

| Aspect | Fantasy | Sci-Fi | Horror | Cyberpunk | Post-Apoc |
|---|---|---|---|---|---|
| Biome theme | Enchanted wilderness | Alien exoplanet | Cursed dead lands | Toxic urban ruin | Irradiated wasteland |
| Primary hazard | Magical creatures, fae curses | Hostile xenolife, equipment failure | Undead, psychological terror | Corp drones, toxic exposure | Raiders, radiation sickness |
| Crafting materials | Herb, bone, runestone, hide | Alloy scrap, polymer, crystal cell | Bone, obsidian, cursed relic | Wire, circuit board, chem canister | Scrap metal, rubber, salvage |
| Shelter style | Wattle longhouse, stone circle | Force-field dome, prefab pod | Bone-frame lean-to, barrow | Corrugated bunker, hack-shack | Salvage shanty, buried dugout |
| Creature behavior | Territorial, magically enhanced | Pack tactics, hive mind | Nocturnal, fear-mechanic driven | Patrol routes, sensor detection | Scavenger packs, ambush |
| Elimination flavor | Banishment ritual | Ejection pod | Exorcism ceremony | Signal blackout | Exile beyond the wall |
| Ambient audio | Folk strings, wind chimes | Synth drones, radio static | Dissonant drones, distant screams | Glitchy beats, neon hum | Wind howl, distant machinery |
| Visual post-process | Warm bloom, vignette | Scanlines, cold LUT | Desaturation, green grain | Chromatic aberration, neon bloom | Sepia wash, film grain |

---

## 7. Success Criteria

| Indicator | Target | Measurement |
|---|---|---|
| World generation determinism | 100% seed reproducibility | CI: 100 seeds × 5 genres × 3 OS — byte-identical checksums |
| World generation speed | <2s for 512×512 map | `go test -bench=BenchmarkWorldGen` on reference hardware |
| Server tick budget | <16ms for 12 players | `go test -bench=BenchmarkServerTick` |
| Latency tolerance | Playable at 5000ms RTT | Integration test with simulated latency shaper |
| Genre distinctiveness | <20% shared recipe IDs across genres | Unit test asserting recipe-set intersection |
| Zero asset files | No binary/image/audio files in repo | CI: `find . -name '*.png' -o -name '*.mp3' -o -name '*.wav'` returns empty |
| Single-binary build | Runs on clean VM | CI artifact test on fresh Docker image |
| Elimination protocol | Vote processed <500ms on LAN | Integration test: 6-bot council vote |

---

## 8. Risk Assessment

| Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|
| PCG determinism breaks across OS/arch | Medium | High | Use seeded PRNG (`math/rand/v2`) exclusively; ban `map` iteration in generators; CI matrix |
| Ebiten API instability between v2 minor versions | Low | Medium | Pin exact Ebiten version in `go.mod`; maintain minimal rendering abstraction layer |
| 5000ms latency makes social mechanics unusable | Medium | High | All vote/crafting windows ≥30s; async design; no real-time blocking calls |
| Ecosystem simulation diverges (extinction/explosion) | Medium | Medium | Cap population deltas per tick; add re-seeding guard when species <5% of initial count |
| Genre audio differentiation insufficient | Low | Medium | Automated spectral-centroid test per genre; assert ≥15% centroid divergence |
| Server memory growth with large worlds | Medium | High | Chunk eviction LRU cache; unload chunks outside all player view radii |
| Single binary size exceeds 50 MB | Low | Low | Profile with `go build -v`; strip debug symbols for release builds |
