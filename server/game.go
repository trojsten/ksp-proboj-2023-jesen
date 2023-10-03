package main

import (
	"fmt"
	"github.com/trojsten/ksp-proboj/client"
	"image/png"
	"math/rand"
	"os"
	"path"
)

type Game struct {
	Map       Map           `json:"map"`
	Players   []Player      `json:"players"`
	Ships     map[int]*Ship `json:"ships"`
	MaxShipId int
	Harbors   []Harbor `json:"harbors"`
	Bases     []Base   `json:"bases"`
	runner    client.Runner
}

func (g *Game) LoadMap(filename string) error {
	f, err := os.OpenFile(path.Join("../../maps", filename), os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	im, err := png.Decode(f)
	if err != nil {
		return err
	}
	size := im.Bounds().Size()
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
			}
			if red == 0 && green == 255 && blue == 0 {
				g.Map.Tiles[y][x] = Tile{Type: TILE_GROUND, Index: -1}
			}
			if red == 255 && green == 0 && blue == 0 {
				g.Map.Tiles[y][x] = Tile{Type: TILE_HARBOR, Index: -1}
				prod := []int{0, 0, 0, 1, -1}
				g.Harbors = append(g.Harbors, Harbor{
					X: y,
					Y: x,
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
			}
			if red == 255 && green == 255 && blue == 255 {
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
			}
		}
	}
	if playerIdx != len(playersOrder) {
		return fmt.Errorf("too many players, not enough bases")
	}

	return nil
}
