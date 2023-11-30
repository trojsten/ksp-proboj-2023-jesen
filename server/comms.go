package main

import (
	"encoding/json"
	"fmt"
	"github.com/trojsten/ksp-proboj/client"
	"strings"
	"time"
)

func sendStateToPlayer(g *Game, p *Player, sendMap bool, round int, maxRounds int) error {
	state := StateForPlayer(g, p, sendMap)
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}

	resp := g.Runner.ToPlayer(p.Name, fmt.Sprintf("ROUND %d/%d", round, maxRounds), string(data))
	if resp != client.Ok {
		return fmt.Errorf("response from Runner: %v", resp)
	}
	return nil
}

func handlePlayer(g *Game, p *Player) error {
	start := time.Now()
	status, resp := g.Runner.ReadPlayer(p.Name)
	end := time.Now()
	if status != client.Ok {
		return fmt.Errorf("(%s) response from Runner: %v", p.Name, status)
	}

	responseTime := end.Sub(start).Microseconds()
	g.Runner.Log(fmt.Sprintf("(%s) player responded in %d us", p.Name, responseTime))
	p.Statistics.addTimeOfResponse(responseTime)
	for _, ship := range p.Ships() {
		p.Statistics.addTimeByShip(ship.Type)
	}

	var commandedShips = map[int]bool{}
	for _, line := range strings.Split(resp, "\n") {
		parts := strings.SplitN(line, " ", 2)
		g.Runner.Log(fmt.Sprintf("(%s) %s", p.Name, line))
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
			g.Runner.Log(fmt.Sprintf("(%s) player send INVALID command (%q): %s", p.Name, line, err))
		}
	}
	return nil
}
