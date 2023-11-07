package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/trojsten/ksp-proboj/client"
)

func (g *Game) Run() error {
	for round := 0; round < MAX_ROUNDS; round++ {
		g.Runner.Log(fmt.Sprintf("started round %d", round))
		playerOrder := rand.Perm(len(g.Players))
		for _, i := range playerOrder {
			player := &g.Players[i]
			err := sendStateToPlayer(g, player, round == 0, round, MAX_ROUNDS)
			if err != nil {
				g.Runner.Log(fmt.Sprintf("error while communicating with player %s: %v", player.Name, err))
				markShipsAsWrecks(player)
				continue
			}

			err = handlePlayer(g, player)
			if err != nil {
				g.Runner.Log(fmt.Sprintf("error while communicating with player %s: %v", player.Name, err))
				markShipsAsWrecks(player)
				continue
			}
			player.Score.updateCurrentGold(player.CurrentGold())
		}

		for i, _ := range g.Harbors {
			g.Harbors[i].tick()
		}

		// apply damage to ships near harbours
		for _, harbor := range g.Harbors {
			for _, ship := range g.Ships {
				if ship.Type.Stats().Class == SHIP_ATTACK {
					if dist(harbor.X, harbor.Y, ship.X, ship.Y) < HARBOUR_DAMAGE_RADIUS {
						ship.Health -= HARBOUR_DAMAGE
						g.Runner.Log(fmt.Sprintf("attack ship %d was near harbour, so applying HARBOUR_DAMAGE", ship.Id))
					}
				}
			}
		}
		// apply damage to ships near bases
		for _, base := range g.Bases {
			for _, ship := range g.Ships {
				if ship.PlayerIndex != base.PlayerIndex {
					if dist(base.X, base.Y, ship.X, ship.Y) < BASE_DAMAGE_RADIUS {
						ship.Health -= BASE_DAMAGE
						g.Runner.Log(fmt.Sprintf("ship %d from player \"%s\" was near base of player \"%s\", so applying BASE_DAMAGE", ship.Id, g.Players[ship.PlayerIndex].Name, g.Players[base.PlayerIndex].Name))
					}
				}
			}
		}

		// heal ships on harbors and bases
		for _, ship := range g.Ships {
			tileType := g.Map.Tiles[ship.Y][ship.X].Type
			if tileType == TILE_HARBOR || tileType == TILE_BASE {
				ship.Health = min(ship.Health+HARBOR_BASE_HEAL, ship.Type.Stats().MaxHealth)
				g.Runner.Log(fmt.Sprintf("ship %d was on base or harbour so applying HARBOR_BASE_HEAL", ship.Id))
			}
		}

		// remove ships
		for i, _ := range g.Ships {
			if g.Ships[i].Health <= 0 && !g.Ships[i].IsWreck {
				g.Ships[i].IsWreck = true
				g.Ships[i].Health = 0
			}
			if g.Ships[i].IsWreck && g.Ships[i].Resources.empty() || g.Ships[i].IsWreck && g.Ships[i].Health <= WRECK_REMOVE_DAMAGE {
				delete(g.Ships, i)
			}
		}

		var gameToMarshall = GameStats{
			Game: Game{
				Map:       nil,
				Players:   g.Players,
				Ships:     g.Ships,
				MaxShipId: g.MaxShipId,
				Harbors:   g.Harbors,
				Bases:     g.Bases,
				Runner:    g.Runner,
			},
			ShipTypes: []string{},
		}
		for _, ship := range g.Ships {
			stateShip := StateShip{Ship: *ship}
			gameToMarshall.ShipTypes = append(gameToMarshall.ShipTypes, stateShip.Type.Name())
		}

		if round == 0 {
			gameToMarshall.Map = g.Map
		}
		data, err := json.Marshal(gameToMarshall)
		if err != nil {
			g.Runner.Log(fmt.Sprintf("could not marshal JSON for observer: %s", err.Error()))
		}
		resp := g.Runner.ToObserver(string(data) + "\n")
		if resp != client.Ok {
			g.Runner.Log(fmt.Sprintf("error while sending data to observer"))
		}
	}
	return nil
}

func markShipsAsWrecks(player *Player) {
	for _, ship := range player.Ships() {
		ship.IsWreck = true
	}
}
