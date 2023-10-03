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
		g.runner.Log(fmt.Sprintf("(%s) %s", p.Name, line))
		command := parts[0]
		args := ""
		if len(parts) == 2 {
			args = parts[1]
		}
		var err error
		switch command {
		case MOVE:
			err = move(g, p, args, commandedShips)
			break
		case TRADE:
			err = trade(g, p, args, commandedShips)
			break
		case LOOT:
			err = loot(g, p, args, commandedShips)
			break
		case SHOOT:
			err = shoot(g, p, args, commandedShips)
			break
		case BUY:
			err = buy(g, p, args, commandedShips)
			break
		case STORE:
			err = store(g, p, args, commandedShips)
			break
		default:
			err = fmt.Errorf("unkown command")
		}
		if err != nil {
			g.runner.Log(fmt.Sprintf("(%s) player send INVALID command (%q): %s", p.Name, line, err))
		}
	}
	return nil
}
