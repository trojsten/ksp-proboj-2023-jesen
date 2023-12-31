package main

import (
	"encoding/json"
	"fmt"
	"github.com/trojsten/ksp-proboj/client"
	"image/png"
	"math/rand"
	"os"
)

type GameStats struct {
	Game
	ShipTypes map[int]string `json:"ship_types"`
}

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
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
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
	g.Map.HeatMap = make([][]int, g.Map.Height)
	for y := 0; y < g.Map.Height; y++ {
		g.Map.Tiles[y] = make([]Tile, g.Map.Width)
		g.Map.HeatMap[y] = make([]int, g.Map.Width)

		for x := 0; x < g.Map.Width; x++ {
			color := im.At(im.Bounds().Min.X+x, im.Bounds().Min.Y+y)
			red, green, blue, _ := color.RGBA()
			red &= 255
			green &= 255
			blue &= 255

			g.Map.HeatMap[y][x] = 0
			if red == 0 && green == 0 && blue == 255 {
				g.Map.Tiles[y][x] = Tile{Type: TILE_WATER, Index: -1}
			} else if red == 0 && green == 128 && blue == 0 {
				g.Map.Tiles[y][x] = Tile{Type: TILE_GROUND, Index: -1}
			} else if red == 255 && green == 0 && blue == 0 {
				g.Map.Tiles[y][x] = Tile{Type: TILE_HARBOR, Index: -1}
				prodLikely := []int{0, 1, 1, 1, -1, -1, -1}
				prodUnlikely := []int{0, 0, 0, 1, -1}
				h := Harbor{
					X: x,
					Y: y,
					Production: Resources{
						Wood:      (prodLikely[rand.Intn(len(prodLikely))] * (BASE_PRODUCTION[0] + rand.Intn(3))),
						Stone:     (prodLikely[rand.Intn(len(prodLikely))] * (BASE_PRODUCTION[1] + rand.Intn(3))),
						Iron:      (prodUnlikely[rand.Intn(len(prodUnlikely))] * (BASE_PRODUCTION[2] + rand.Intn(3))),
						Gem:       (prodUnlikely[rand.Intn(len(prodUnlikely))] * (BASE_PRODUCTION[3] + rand.Intn(3))),
						Wool:      (prodUnlikely[rand.Intn(len(prodUnlikely))] * (BASE_PRODUCTION[4] + rand.Intn(3))),
						Hide:      (prodUnlikely[rand.Intn(len(prodUnlikely))] * (BASE_PRODUCTION[5] + rand.Intn(3))),
						Wheat:     (prodUnlikely[rand.Intn(len(prodUnlikely))] * (BASE_PRODUCTION[6] + rand.Intn(3))),
						Pineapple: (prodUnlikely[rand.Intn(len(prodUnlikely))] * (BASE_PRODUCTION[7] + rand.Intn(3))),
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
				}
				if h.Production.Wood < 1 && h.Production.Stone < 1 && h.Production.Iron < 1 &&
					h.Production.Gem < 1 && h.Production.Wool < 1 && h.Production.Hide < 1 &&
					h.Production.Wheat < 1 && h.Production.Pineapple < 1 {

					r := ResourceType(rand.Intn(7))
					*h.Production.Resource(r) = rand.Intn(4) + 1
				}
				g.Harbors = append(g.Harbors, h)

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

type GlobalStatistics struct {
	Players map[string]Statistics `json:"players"`
	HeatMap [][]int               `json:"heatmap"`
}

func (g *Game) SaveStats() error {
	stats := GlobalStatistics{
		Players: map[string]Statistics{},
	}
	for _, player := range g.Players {
		stats.Players[player.Name] = player.Statistics
	}
	stats.HeatMap = g.Map.HeatMap

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
