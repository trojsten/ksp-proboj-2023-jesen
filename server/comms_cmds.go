package main

import "fmt"

func move(g *Game, p *Player, args string, commandedShips map[int]bool) error {
	var shipId, x, y int
	_, err := fmt.Sscanf(args, "%d %d %d", &shipId, &x, &y)
	if err != nil {
		return fmt.Errorf("(%s) sscanf of command MOVE failed: %w", p.Name, err)
	}
	_, exist := commandedShips[shipId]
	if exist {
		g.runner.Log(fmt.Sprintf("(%s) player send multiple commands to ship %d", p.Index, shipId))
		return nil
	}
	ship := g.Ships[shipId]
	if ship.PlayerIndex != p.Index {
		return fmt.Errorf("(%s) try to command ship %d which dont belong to him", p.Name, shipId)
	}
	commandedShips[shipId] = true
	if !IsReachableBfs(g, ship.X, ship.Y, x, y, ship.Type.Stats().MaxMoveRange) {
		return fmt.Errorf("player wanted to move ship %d out of its range", shipId)
	}

	ship.X = min(max(x, 0), g.Map.Width-1)
	ship.Y = min(max(y, 0), g.Map.Height-1)
	g.Ships[shipId] = ship
	return nil
}

func trade(g *Game, p *Player, line string, commandedShips map[int]bool) error {
	var shipId, resourceId, amount int
	_, err := fmt.Sscanf(line, "%d %d %d", &shipId, &resourceId, &amount)
	if err != nil {
		return fmt.Errorf("(%s) sscanf of command TRADE failed: %w", p.Name, err)
	}
	_, exist := commandedShips[shipId]
	if exist {
		g.runner.Log(fmt.Sprintf("(%s) player send multiple commands to ship %d", p.Index, shipId))
		return nil
	}
	if g.Ships[shipId].PlayerIndex != p.Index {
		return fmt.Errorf("(%s) try to command ship %d which dont belong to him", p.Name, shipId)
	}
	commandedShips[shipId] = true
	if g.Ships[shipId].IsWreck {
		g.runner.Log(fmt.Sprintf("(%s) player send commands to ship %d, which is already wreck", p.Index, shipId))
		return nil
	}
	var resource = g.Ships[shipId].Resources.Resource(ResourceType(resourceId))
	if resource == nil {
		g.runner.Log(fmt.Sprintf("(%s) player send commands to ship %d to trade resource %d, which is not valid", p.Index, shipId, resourceId))
		return nil
	}

	var harbor Harbor

	for _, h := range g.Harbors {
		if h.X == g.Ships[shipId].X && h.Y == g.Ships[shipId].Y {
			harbor = h
			break
		}
	}

	if amount > 0 { // we take from harbor
		amount = min(amount, *harbor.Storage.Resource(ResourceType(resourceId)))
		*g.Ships[shipId].Resources.Resource(ResourceType(resourceId)) += amount
		*harbor.Storage.Resource(ResourceType(resourceId)) -= amount
		g.Ships[shipId].Resources.Gold -= amount // TODO vzorec
	} else { // we take give to harbor
		amount = min(-1*amount, *g.Ships[shipId].Resources.Resource(ResourceType(resourceId)))
		*g.Ships[shipId].Resources.Resource(ResourceType(resourceId)) -= amount
		*harbor.Storage.Resource(ResourceType(resourceId)) += amount
		g.Ships[shipId].Resources.Gold += amount // TODO vzorec
	}
	return nil
}

func loot(g *Game, p *Player, line string, commandedShips map[int]bool) error {
	var shipId int
	_, err := fmt.Sscanf(line, "%d", &shipId)
	if err != nil {
		return fmt.Errorf("(%s) sscanf of command LOOT failed: %w", p.Name, err)
	}
	_, exist := commandedShips[shipId]
	if exist {
		g.runner.Log(fmt.Sprintf("(%s) player send multiple commands to ship %d", p.Name, shipId))
		return nil
	}
	if g.Ships[shipId].PlayerIndex != p.Index {
		return fmt.Errorf("(%s) try to command ship %d which dont belong to him", p.Name, shipId)
	}
	commandedShips[shipId] = true
	if g.Ships[shipId].IsWreck {
		g.runner.Log(fmt.Sprintf("(%s) player send commands to ship %d, which is already wreck", p.Index, shipId))
		return nil
	}

	ship := ShipAt(g, g.Ships[shipId].X, g.Ships[shipId].Y)
	if ship != nil && ship.IsWreck {
		ship.Resources = Resources{
			Wood:      ship.Resources.Wood + int(ship.Type.Stats().Yield*float32(g.Ships[shipId].Resources.Wood)),
			Stone:     ship.Resources.Stone + int(ship.Type.Stats().Yield*float32(g.Ships[shipId].Resources.Stone)),
			Iron:      ship.Resources.Iron + int(ship.Type.Stats().Yield*float32(g.Ships[shipId].Resources.Iron)),
			Gem:       ship.Resources.Gem + int(ship.Type.Stats().Yield*float32(g.Ships[shipId].Resources.Gem)),
			Wool:      ship.Resources.Wool + int(ship.Type.Stats().Yield*float32(g.Ships[shipId].Resources.Wool)),
			Hide:      ship.Resources.Hide + int(ship.Type.Stats().Yield*float32(g.Ships[shipId].Resources.Hide)),
			Wheat:     ship.Resources.Wheat + int(ship.Type.Stats().Yield*float32(g.Ships[shipId].Resources.Wheat)),
			Pineapple: ship.Resources.Pineapple + int(ship.Type.Stats().Yield*float32(g.Ships[shipId].Resources.Pineapple)),
			Gold:      ship.Resources.Gold + int(ship.Type.Stats().Yield*float32(g.Ships[shipId].Resources.Gold)),
		}
		ship.Resources = Resources{
			Wood:      0,
			Stone:     0,
			Iron:      0,
			Gem:       0,
			Wool:      0,
			Hide:      0,
			Wheat:     0,
			Pineapple: 0,
			Gold:      0,
		}
	}

	return nil
}

func shoot(g *Game, p *Player, line string, commandedShips map[int]bool) error {
	var shipId, targetShipId int
	_, err := fmt.Sscanf(line, "%d %d", &shipId, &targetShipId)
	if err != nil {
		return fmt.Errorf("(%s) sscanf of command SHOOT failed: %w", p.Name, err)
	}
	_, exist := commandedShips[shipId]
	if exist {
		g.runner.Log(fmt.Sprintf("(%s) player send multiple commands to ship %d", p.Name, shipId))
		return nil
	}
	if g.Ships[shipId].PlayerIndex != p.Index {
		return fmt.Errorf("(%s) try to command ship %d which dont belong to him", p.Name, shipId)
	}
	commandedShips[shipId] = true
	if g.Ships[shipId].IsWreck {
		g.runner.Log(fmt.Sprintf("(%s) player send commands to ship %d, which is already wreck", p.Name, shipId))
		return nil
	}

	if dist(g.Ships[shipId].X, g.Ships[shipId].Y, g.Ships[targetShipId].X, g.Ships[targetShipId].Y) <= g.Ships[shipId].Type.Stats().Range {
		ship := g.Ships[targetShipId]
		ship.Health -= g.Ships[shipId].Type.Stats().Damage
		g.Ships[targetShipId] = ship
	}
	return nil
}

func buy(g *Game, p *Player, line string, commandedShips map[int]bool) error {
	var shipId int
	_, err := fmt.Sscanf(line, "%d", &shipId)
	if err != nil {
		return fmt.Errorf("(%s) sscanf of command BUY failed: %w", p.Name, err)
	}

	if shipId < 0 || shipId >= len(ships) {
		return fmt.Errorf("(%s) sscanf of ShipTypeId out of range: %d", p.Name, shipId)
	}

	if g.Ships[shipId].PlayerIndex != p.Index {
		return fmt.Errorf("(%s) try to command ship %d which dont belong to him", p.Name, shipId)
	}
	if p.Gold < ships[shipId].Stats().Price {
		return fmt.Errorf("(%s) try to buy ship %d and dont have enough gold", p.Name, shipId)
	}
	p.Gold -= ships[shipId].Stats().Price
	commandedShips[shipId] = true
	base := p.Base()
	if base != nil {
		g.Ships[g.MaxShipId] = &Ship{
			Id:          g.MaxShipId,
			PlayerIndex: p.Index,
			Type:        ships[shipId],
			X:           base.X,
			Y:           base.Y,
			Health:      ships[shipId].Stats().MaxHealth,
			IsWreck:     false,
			Resources:   Resources{},
		}
		g.MaxShipId++
	}

	return nil
}

func store(g *Game, p *Player, line string, commandedShips map[int]bool) error {
	var amount int
	_, err := fmt.Sscanf(line, "%d", &amount)
	if err != nil {
		return fmt.Errorf("(%s) sscanf of command STORE failed: %w", p.Name, err)
	}

	base := p.Base()

	if base != nil {
		for _, ship := range p.Ships() {
			if ship.X == base.X && ship.Y == base.Y {
				_, exist := commandedShips[ship.Id]
				if exist {
					g.runner.Log(fmt.Sprintf("(%s) player send multiple commands to ship %d", p.Name, ship.Id))
					return nil
				}
				commandedShips[ship.Id] = true

				if amount > 0 {
					var goldToStore = min(amount, ship.Resources.Gold)
					p.Gold += goldToStore
					ship.Resources.Gold -= goldToStore
				} else {
					var goldToRemove = min(-1*amount, p.Gold)
					ship.Resources.Gold += goldToRemove
					p.Gold -= goldToRemove
				}
			}
		}
	}
	return nil
}
