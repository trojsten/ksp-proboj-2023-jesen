package main

import (
	"encoding/json"
	"fmt"
	"github.com/trojsten/ksp-proboj/client"
	"strings"
)

func sendStateToPlayer(g *Game, p *Player) error {
	state := StateForPlayer(g, p)
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}

	resp := g.runner.ToPlayer(p.Name, "", string(data))
	if resp != client.Ok {
		return fmt.Errorf("response from runner: %v", resp)
	}
	return nil
}

func handlePlayer(g *Game, p *Player) error {
	status, resp := g.runner.ReadPlayer(p.Name)

	if status != client.Ok {
		return fmt.Errorf("(%s) response from runner: %v", p.Name, status)
	}

	var commandedShips = map[int]bool{}
	for _, line := range strings.Split(resp, "\n") {
		parts := strings.SplitN(line, " ", 2)
		command := parts[0]
		args := ""
		if len(parts) == 2 {
			args = parts[1]
		}

		switch command {
		case MOVE:
			move(g, p, args, commandedShips)
			break
		case TRADE:
			trade(g, p, args, commandedShips)
			break
		case LOOT:
			loot(g, p, args, commandedShips)
			break
		case SHOOT:
			shoot(g, p, args, commandedShips)
			break
		case BUY:
			buy(g, p, args, commandedShips)
			break
		case STORE:
			store(g, p, args, commandedShips)
			break
		}
	}
	return nil
}
