package main

type StateShip struct {
	Ship
	Mine  bool      `json:"mine"`
	Stats ShipStats `json:"stats"`
}

type StateHarbor struct {
	Harbor
	Visible bool `json:"visible"`
}

type State struct {
	Map     Map           `json:"map"`
	Harbors []StateHarbor `json:"harbors"`
	Ships   []StateShip   `json:"ships"`
	Gold    int           `json:"gold"`
	Index   int           `json:"index"`
}

func StateForPlayer(g *Game, p *Player, sendMap bool) (state State) {
	if sendMap {
		state.Map = *g.Map
	}
	state.Gold = p.Gold
	state.Index = p.Index
	state.Ships = []StateShip{}
	for _, harbor := range g.Harbors {
		var stateHarbor StateHarbor

		var harborVisible = false
		for _, ship := range p.Ships() {
			if dist(ship.X, ship.Y, harbor.X, harbor.Y) < ship.Type.Stats().Range {
				harborVisible = true
				break
			}
		}
		stateHarbor.Harbor = harbor
		stateHarbor.Visible = harborVisible
		if !harborVisible {
			stateHarbor.Storage = Resources{}
			stateHarbor.Production = Resources{}
		}

		state.Harbors = append(state.Harbors, stateHarbor)
	}

	// Add visible ships
	for _, ship := range g.Ships {
		visible := false
		for _, myShip := range p.Ships() {
			if float64(dist(ship.X, ship.Y, myShip.X, myShip.Y)) <= float64(myShip.Type.Stats().Range)*SHIP_SEE_RANGE {
				visible = true
				break
			}
		}

		stateShip := StateShip{Ship: *ship}

		if !visible {
			continue
		}
		stateShip.Stats = stateShip.Type.Stats()
		stateShip.Type = nil

		if stateShip.PlayerIndex != p.Index {
			stateShip.Mine = false
			stateShip.Resources = Resources{}
			stateShip.Health = 0
		} else {
			stateShip.Mine = true
		}
		state.Ships = append(state.Ships, stateShip)
	}

	return
}
