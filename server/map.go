package main

type TileType int

const (
	TILE_WATER TileType = iota
	TILE_GROUND
	TILE_HARBOR
	TILE_BASE
)

type Tile struct {
	Type  TileType `json:"type"`
	Index int      `json:"index"`
}

type Base struct {
	X           int
	Y           int
	PlayerIndex int
}

type Harbor struct {
	X          int       `json:"x"`
	Y          int       `json:"y"`
	Production Resources `json:"production"`
	Storage    Resources `json:"storage"`
}

func (h *Harbor) tick() {
	h.Storage = Resources{
		Wood:      max(0, h.Storage.Wood+h.Production.Wood),
		Stone:     max(0, h.Storage.Stone+h.Production.Stone),
		Iron:      max(0, h.Storage.Iron+h.Production.Iron),
		Gem:       max(0, h.Storage.Gem+h.Production.Gem),
		Wool:      max(0, h.Storage.Wool+h.Production.Wool),
		Hide:      max(0, h.Storage.Hide+h.Production.Hide),
		Wheat:     max(0, h.Storage.Wheat+h.Production.Wheat),
		Pineapple: max(0, h.Storage.Pineapple+h.Production.Pineapple),
		Gold:      0,
	}
}

type Map struct {
	Tiles  [][]Tile `json:"tiles"`
	Width  int      `json:"width"`
	Height int      `json:"height"`
}

type XY struct {
	X int
	Y int
}

func Adjacent(x int, y int, g *Game) []XY {
	temp := []XY{{x, y - 1}, {x - 1, y}, {x + 1, y}, {x, y + 1}}
	m := g.Map

	var result []XY
	for _, v := range temp {
		if v.X < 0 || v.X >= m.Width || v.Y < 0 || v.Y >= m.Height {
			continue
		}

		if m.Tiles[v.Y][v.X].Type != TILE_WATER {
			continue
		}

		ship := ShipAt(g, v.X, v.Y)
		if ship != nil && !ship.IsWreck {
			continue
		}

		result = append(result, v)
	}
	return result
}

func IsReachableBfs(g *Game, x1 int, y1 int, x2 int, y2 int, maxDist int) bool {
	var queue []XY
	queue = append(queue, XY{X: x1, Y: y1})

	visited := map[XY]bool{}
	for len(queue) != 0 {
		pos := queue[0]
		queue = queue[1:]

		visited[pos] = true

		if pos.X == x2 && pos.Y == y2 {
			return true
		}

		if dist(pos.X, pos.Y, x1, y1) >= maxDist {
			continue
		}

		for _, next := range Adjacent(pos.X, pos.Y, g) {
			if visited[next] {
				continue
			}

			queue = append(queue, next)
		}
	}
	return false
}
