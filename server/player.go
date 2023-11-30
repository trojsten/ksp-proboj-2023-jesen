package main

import "fmt"

type Player struct {
	Index      int    `json:"index"`
	Name       string `json:"name"`
	Gold       int    `json:"gold"`
	Score      Score  `json:"score"`
	Statistics Statistics
	game       *Game
}

func NewPlayer(game *Game, idx int, name string) Player {
	return Player{
		Index: idx,
		Name:  name,
		game:  game,
		Gold:  NEW_PLAYER_GOLD,
		Statistics: Statistics{
			Kills:           map[string]int{},
			Damage:          map[string]int{},
			SellsByType:     map[int]int{},
			PurchasesByType: map[int]int{},
			TimeByShip:      map[string]int{},
			TimeOfResponses: 0,
		},
	}
}

func (p *Player) Ships() []*Ship {
	var playersShips []*Ship
	for _, ship := range p.game.Ships {
		if ship.PlayerIndex == p.Index {
			playersShips = append(playersShips, ship)
		}
	}
	return playersShips
}

func (p *Player) Ship(g *Game, shipId int) (*Ship, error) {
	ship := g.Ships[shipId]
	if ship == nil {
		return nil, fmt.Errorf("ship %d doesn't exists", shipId)
	}
	if ship.PlayerIndex != p.Index {
		return nil, fmt.Errorf("try to command ship %d which dont belong to him", shipId)
	}
	if g.Ships[shipId].IsWreck {
		return nil, fmt.Errorf("ship %d is already wreck", shipId)
	}
	return ship, nil
}

func (p *Player) Base() *Base {
	for _, base := range p.game.Bases {
		if base.PlayerIndex == p.Index {
			return &base
		}
	}
	return nil
}

func (p *Player) CurrentGold() int {
	res := p.Gold
	for _, ship := range p.Ships() {
		res += ship.Resources.Gold
	}
	return res
}
