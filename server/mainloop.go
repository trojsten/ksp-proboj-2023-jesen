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
			err := sendStateToPlayer(g, player)
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

		for _, harbor := range g.Harbors {
			harbor.tick()
		}

		// apply damage to ships near harbours
		for _, harbor := range g.Harbors {
			for _, ship := range g.Ships {
				if ship.Type.Stats().Class == SHIP_ATTACK {
					if dist(harbor.X, harbor.Y, ship.X, ship.Y) < HARBOUR_DAMAGE_RADIUS {
						ship.Health -= HARBOUR_DAMAGE
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
					}
				}
			}
		}

		// remove ships
		for i, ship := range g.Ships {
			if ship.IsWreck && ship.Resources.empty() {
				delete(g.Ships, i)
			} else if ship.Health <= 0 {
				g.Ships[i].IsWreck = true
			}
		}

		var gameToMarshall = Game{
			Map:       nil,
			Players:   g.Players,
			Ships:     g.Ships,
			MaxShipId: g.MaxShipId,
			Harbors:   g.Harbors,
			Bases:     g.Bases,
			Runner:    g.Runner,
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
