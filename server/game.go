package main

import (
	"encoding/json"
	"fmt"
	"github.com/trojsten/ksp-proboj/client"
	"image/png"
	"math/rand"
	"os"
	"path"
)

type Game struct {
	Map       *Map          `json:"map"`
	Players   []Player      `json:"players"`
	Ships     map[int]*Ship `json:"ships"`
	MaxShipId int
	Harbors   []Harbor `json:"harbors"`
	Bases     []Base   `json:"bases"`
	Runner    client.Runner
}

func (g *Game) LoadMap(filename string) error {
	f, err := os.OpenFile(path.Join("../maps", filename), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	im, err := png.Decode(f)
	if err != nil {
		return err
	}
	size := im.Bounds().Size()
	g.Map = &Map{}
	g.Map.Width = size.X
	g.Map.Height = size.Y

	playersOrder := rand.Perm(len(g.Players))
	playerIdx := 0

	g.Map.Tiles = make([][]Tile, g.Map.Height)
	for y := 0; y < g.Map.Height; y++ {
		g.Map.Tiles[y] = make([]Tile, g.Map.Width)

		for x := 0; x < g.Map.Width; x++ {
			color := im.At(im.Bounds().Min.X+x, im.Bounds().Min.Y+y)
			red, green, blue, _ := color.RGBA()
			red &= 255
			green &= 255
			blue &= 255

			if red == 0 && green == 0 && blue == 255 {
				g.Map.Tiles[y][x] = Tile{Type: TILE_WATER, Index: -1}
			} else if red == 0 && green == 128 && blue == 0 {
				g.Map.Tiles[y][x] = Tile{Type: TILE_GROUND, Index: -1}
			} else if red == 255 && green == 0 && blue == 0 {
				g.Map.Tiles[y][x] = Tile{Type: TILE_HARBOR, Index: -1}
				prod := []int{0, 0, 0, 1, -1}
				g.Harbors = append(g.Harbors, Harbor{
					X: x,
					Y: y,
					Production: Resources{
						Wood:      prod[rand.Intn(len(prod))] * rand.Intn(5),
						Stone:     prod[rand.Intn(len(prod))] * rand.Intn(5),
						Iron:      prod[rand.Intn(len(prod))] * rand.Intn(5),
						Gem:       prod[rand.Intn(len(prod))] * rand.Intn(5),
						Wool:      prod[rand.Intn(len(prod))] * rand.Intn(5),
						Hide:      prod[rand.Intn(len(prod))] * rand.Intn(5),
						Wheat:     prod[rand.Intn(len(prod))] * rand.Intn(5),
						Pineapple: prod[rand.Intn(len(prod))] * rand.Intn(5),
						Gold:      0,
					},
					Storage: Resources{
						Wood:      0,
						Stone:     0,
						Iron:      0,
						Gem:       0,
						Wool:      0,
						Hide:      0,
						Wheat:     0,
						Pineapple: 0,
						Gold:      0,
					},
				})
			} else if red == 255 && green == 255 && blue == 255 {
				if playerIdx >= len(playersOrder) {
					continue
				}

				g.Map.Tiles[y][x] = Tile{Type: TILE_BASE, Index: playersOrder[playerIdx]}
				g.Bases = append(g.Bases, Base{
					X:           x,
					Y:           y,
					PlayerIndex: playersOrder[playerIdx],
				})
				playerIdx++
			} else {
				g.Runner.Log(fmt.Sprintf("unkown color in map: %d %d %d", red, green, blue))
			}
		}
	}
	if playerIdx != len(playersOrder) {
		return fmt.Errorf("too many players, not enough bases")
	}

	return nil
}

func (g *Game) SaveStats() error {
	stats := map[string]Statistics{}
	for _, player := range g.Players {
		stats[player.Name] = player.Statistics
	}

	f, err := os.Create("stats.json")
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := json.Marshal(stats)
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	return err
}
