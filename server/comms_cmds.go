package main

import "fmt"

func move(g *Game, p *Player, args string, commandedShips map[int]bool) error {
	var shipId, x, y int
	_, err := fmt.Sscanf(args, "%d %d %d", &shipId, &x, &y)
	if err != nil {
		return fmt.Errorf("sscanf of command MOVE failed: %w", err)
	}
	_, exist := commandedShips[shipId]
	if exist {
		return fmt.Errorf("multiple commands to ship %d. ignoring", shipId)
	}
	ship, err := p.Ship(g, shipId)
	if err != nil {
		return err
	}
	commandedShips[shipId] = true
	if !IsReachableBfs(g, ship.X, ship.Y, x, y, ship.Type.Stats().MaxMoveRange) {
		return fmt.Errorf("wanted to move ship %d out of its range", shipId)
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
		return fmt.Errorf("sscanf of command TRADE failed: %w", err)
	}
	_, exist := commandedShips[shipId]
	if exist {
		return fmt.Errorf("multiple commands to ship %d. ignoring", shipId)
	}
	ship, err := p.Ship(g, shipId)
	if err != nil {
		return err
	}
	commandedShips[shipId] = true
	var resource = ship.Resources.Resource(ResourceType(resourceId))
	if resource == nil {
		g.runner.Log(fmt.Sprintf("player send commands to ship %d to trade resource %d, which is not valid", shipId, resourceId))
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
		return fmt.Errorf("sscanf of command LOOT failed: %w", err)
	}
	_, exist := commandedShips[shipId]
	if exist {
		return fmt.Errorf("multiple commands to ship %d. ignoring", shipId)
	}
	ship, err := p.Ship(g, shipId)
	if err != nil {
		return err
	}
	commandedShips[shipId] = true

	wreckShip := ShipAt(g, g.Ships[shipId].X, g.Ships[shipId].Y)
	if wreckShip != nil && wreckShip.IsWreck {
		ship.Resources = Resources{
			Wood:      ship.Resources.Wood + int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Wood)),
			Stone:     ship.Resources.Stone + int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Stone)),
			Iron:      ship.Resources.Iron + int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Iron)),
			Gem:       ship.Resources.Gem + int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Gem)),
			Wool:      ship.Resources.Wool + int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Wool)),
			Hide:      ship.Resources.Hide + int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Hide)),
			Wheat:     ship.Resources.Wheat + int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Wheat)),
			Pineapple: ship.Resources.Pineapple + int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Pineapple)),
			Gold:      ship.Resources.Gold + int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Gold)),
		}
		wreckShip.Resources = Resources{
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
		return fmt.Errorf("sscanf of command SHOOT failed: %w", err)
	}
	_, exist := commandedShips[shipId]
	if exist {
		return fmt.Errorf("multiple commands to ship %d. ignoring", shipId)
	}
	ship, err := p.Ship(g, shipId)
	if err != nil {
		return err
	}
	commandedShips[shipId] = true

	enemyShip := g.Ships[targetShipId]

	if dist(ship.X, ship.Y, enemyShip.X, enemyShip.Y) <= ship.Type.Stats().Range {
		enemyShip.Health -= g.Ships[shipId].Type.Stats().Damage
		g.Ships[targetShipId] = enemyShip
	}
	return nil
}

func buy(g *Game, p *Player, line string, commandedShips map[int]bool) error {
	var shipTypeId int
	_, err := fmt.Sscanf(line, "%d", &shipTypeId)
	if err != nil {
		return fmt.Errorf("sscanf of command BUY failed: %w", err)
	}

	if shipTypeId < 0 || shipTypeId >= len(ships) {
		return fmt.Errorf("sscanf of ShipTypeId out of range: %d", shipTypeId)
	}

	if p.Gold < ships[shipTypeId].Stats().Price {
		return fmt.Errorf("try to buy ship %d and dont have enough gold", shipTypeId)
	}
	p.Gold -= ships[shipTypeId].Stats().Price
	commandedShips[shipTypeId] = true
	base := p.Base()
	if base != nil {
		g.Ships[g.MaxShipId] = &Ship{
			Id:          g.MaxShipId,
			PlayerIndex: p.Index,
			Type:        ships[shipTypeId],
			X:           base.X,
			Y:           base.Y,
			Health:      ships[shipTypeId].Stats().MaxHealth,
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
		return fmt.Errorf("sscanf of command STORE failed: %w", err)
	}

	base := p.Base()

	if base != nil {
		for i, _ := range p.Ships() {
			if p.Ships()[i].X == base.X && p.Ships()[i].Y == base.Y {
				ship, err := p.Ship(g, p.Ships()[i].Id)
				if err != nil {
					return err
				}
				commandedShips[ship.Id] = true

				if amount > 0 {
					var goldToStore = min(amount, p.Ships()[i].Resources.Gold)
					p.Gold += goldToStore
					p.Ships()[i].Resources.Gold -= goldToStore
				} else {
					var goldToRemove = min(-1*amount, p.Gold)
					p.Ships()[i].Resources.Gold += goldToRemove
					p.Gold -= goldToRemove
				}
			}
		}
	}
	return nil
}
