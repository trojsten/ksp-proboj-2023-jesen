package main

import (
	"fmt"
	"math"
)

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
	if g.Map.Tiles[y][x].Type == TILE_WATER {
		targetShip := ShipAt(g, x, y)
		if targetShip != nil && !targetShip.IsWreck {
			return fmt.Errorf("there is non wreck ship on coordinates where ship %d wanted to move", shipId)
		}
	}

	ship.X = min(max(x, 0), g.Map.Width-1)
	ship.Y = min(max(y, 0), g.Map.Height-1)
	g.Ships[shipId] = ship
	return nil
}

func price(resourceType ResourceType, amount int) int {
	return min(100/(amount+3)+1, 4) * BASE_PRICE[resourceType]
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
	if resource == nil || resourceId == int(RESOURCE_GOLD) {
		return fmt.Errorf("player send commands to ship %d to trade resource %d, which is INVALID", shipId, resourceId)
	}

	var harbor *Harbor = nil

	for _, h := range g.Harbors {
		if h.X == g.Ships[shipId].X && h.Y == g.Ships[shipId].Y {
			harbor = &h
			break
		}
	}
	if harbor == nil {
		return fmt.Errorf("ship %d is not in the harbor", shipId)
	}

	if amount > 0 { // we take from harbor
		if *harbor.Production.Resource(ResourceType(resourceId)) <= 0 {
			return fmt.Errorf("cannot take resource %d from harbor, because production is not greater than 0", resourceId)
		}
		g.Runner.Log(fmt.Sprintf("(%s) try to TAKE %d pieces. Harbor storage: %d", p.Name, amount, *harbor.Storage.Resource(ResourceType(resourceId))))
		amount = min(amount, *harbor.Storage.Resource(ResourceType(resourceId)))
		g.Runner.Log(fmt.Sprintf(", so taking %d\n", amount))
		unitPrice := price(ResourceType(resourceId), *harbor.Storage.Resource(ResourceType(resourceId)))
		price := unitPrice * amount
		if price > g.Ships[shipId].Resources.Gold {
			return fmt.Errorf("ship %d don't have enough gold to trade", shipId)
		}
		if g.Ships[shipId].Resources.countResources()+amount > g.Ships[shipId].Type.Stats().MaxCargo {
			return fmt.Errorf("ship %d don't have enough cargo space to make a trade", shipId)
		}
		if amount == 0 {
			return fmt.Errorf("result amount for trade is 0, so not trading")
		}
		*g.Ships[shipId].Resources.Resource(ResourceType(resourceId)) += amount
		*harbor.Storage.Resource(ResourceType(resourceId)) -= amount
		g.Ships[shipId].Resources.Gold -= price
		p.Score.newPurchase(amount)
		p.Statistics.newPurchase(resourceId, amount)
	} else { // we give to harbor
		if *harbor.Production.Resource(ResourceType(resourceId)) >= 0 {
			return fmt.Errorf("cannot take resource %d from harbor, because production is not lesser than 0", resourceId)
		}
		g.Runner.Log(fmt.Sprintf("(%s) try to GIVE %d pieces. Ship storage: %d", p.Name, -1*amount, *g.Ships[shipId].Resources.Resource(ResourceType(resourceId))))
		amount = min(-1*amount, *g.Ships[shipId].Resources.Resource(ResourceType(resourceId)))
		g.Runner.Log(fmt.Sprintf(", so giving %d\n", amount))
		if amount == 0 {
			return fmt.Errorf("result amount for trade is 0, so not trading")
		}
		*g.Ships[shipId].Resources.Resource(ResourceType(resourceId)) -= amount
		*harbor.Storage.Resource(ResourceType(resourceId)) += amount
		unitPrice := price(ResourceType(resourceId), *harbor.Storage.Resource(ResourceType(resourceId)))
		price := unitPrice * amount
		g.Ships[shipId].Resources.Gold += price
		p.Score.newSell(amount)
		p.Score.newGoldEarned(price)
		p.Statistics.newSell(resourceId, amount)
	}
	return nil
}

func loot(g *Game, p *Player, line string, commandedShips map[int]bool) error {
	var shipId, targetShipId int
	_, err := fmt.Sscanf(line, "%d %d", &shipId, &targetShipId)
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

	wreckShip := g.Ships[targetShipId]
	if wreckShip != nil && wreckShip.IsWreck {
		remainingSpace := ship.Type.Stats().MaxCargo - ship.Resources.countResources()
		ship.Resources.Gold += min(int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Gold)), remainingSpace)

		remainingSpace = ship.Type.Stats().MaxCargo - ship.Resources.countResources()
		ship.Resources.Gem += min(int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Gem)), remainingSpace)

		remainingSpace = ship.Type.Stats().MaxCargo - ship.Resources.countResources()
		ship.Resources.Iron += min(int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Iron)), remainingSpace)

		remainingSpace = ship.Type.Stats().MaxCargo - ship.Resources.countResources()
		ship.Resources.Hide += min(int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Hide)), remainingSpace)

		remainingSpace = ship.Type.Stats().MaxCargo - ship.Resources.countResources()
		ship.Resources.Wool += min(int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Wool)), remainingSpace)

		remainingSpace = ship.Type.Stats().MaxCargo - ship.Resources.countResources()
		ship.Resources.Pineapple += min(int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Pineapple)), remainingSpace)

		remainingSpace = ship.Type.Stats().MaxCargo - ship.Resources.countResources()
		ship.Resources.Wheat += min(int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Wheat)), remainingSpace)

		remainingSpace = ship.Type.Stats().MaxCargo - ship.Resources.countResources()
		ship.Resources.Stone += min(int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Stone)), remainingSpace)

		remainingSpace = ship.Type.Stats().MaxCargo - ship.Resources.countResources()
		ship.Resources.Wood += min(int(ship.Type.Stats().Yield*float32(wreckShip.Resources.Wood)), remainingSpace)

		wreckShip.Resources = Resources{}
		p.Score.newGoldEarned(int(ship.Type.Stats().Yield * float32(wreckShip.Resources.Gold)))
	} else {
		return fmt.Errorf("ship %d try to loot ship %d which not exist or is not wreck", shipId, targetShipId)
	}

	return nil
}

func minDistanceToHarborAndBase(g *Game, x int, y int) int {
	res := math.MaxInt
	for _, harbor := range g.Harbors {
		res = min(res, dist(harbor.X, harbor.Y, x, y))
	}
	for _, base := range g.Bases {
		res = min(res, dist(base.X, base.Y, x, y))
	}
	return res
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

	if enemyShip == nil {
		return fmt.Errorf("enemy ship with id %d not exist", targetShipId)
	}

	distance := dist(ship.X, ship.Y, enemyShip.X, enemyShip.Y)
	if distance <= ship.Type.Stats().Range {
		if minDistanceToHarborAndBase(g, enemyShip.X, enemyShip.Y) > HARBOUR_DAMAGE_RADIUS/2 {
			enemyShip.Health -= g.Ships[shipId].Type.Stats().Damage
			if enemyShip.Health <= 0 {
				p.Score.newKill()
				p.Statistics.newKill(g.Players[enemyShip.PlayerIndex].Name)
			}
			p.Statistics.addDamage(g.Players[enemyShip.PlayerIndex].Name, g.Ships[shipId].Type.Stats().Damage)
			g.Ships[targetShipId] = enemyShip
		} else {
			return fmt.Errorf("refuse to shoot. Enemy ship is within range of some harbor")
		}
	} else {
		return fmt.Errorf("enemy ship with out of range (distance: %d, range: %d)", distance, ship.Type.Stats().Range)
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
		return fmt.Errorf("ShipTypeId out of range: %d", shipTypeId)
	}

	if p.Gold < ships[shipTypeId].Stats().Price {
		return fmt.Errorf("try to buy ship %d and dont have enough gold (price: %d, have %d)", shipTypeId, ships[shipTypeId].Stats().Price, p.Gold)
	}
	p.Gold -= ships[shipTypeId].Stats().Price
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
	var shipId, amount int
	_, err := fmt.Sscanf(line, "%d %d", &shipId, &amount)
	if err != nil {
		return fmt.Errorf("sscanf of command STORE failed: %w", err)
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

	base := p.Base()

	if base != nil {
		if ship.X == base.X && ship.Y == base.Y {
			if amount > 0 {
				g.Runner.Log(fmt.Sprintf("(%s) try to STORE %d golds. Ship storage: %d", p.Name, amount, ship.Resources.Gold))
				var goldToStore = min(amount, ship.Resources.Gold)
				p.Gold += goldToStore
				g.Ships[shipId].Resources.Gold -= goldToStore
			} else {
				g.Runner.Log(fmt.Sprintf("(%s) try to WITHDRAW %d golds. Player gold: %d.", p.Name, -1*amount, p.Gold))
				var goldToRemove = min(-1*amount, p.Gold)
				g.Ships[shipId].Resources.Gold += goldToRemove
				p.Gold -= goldToRemove
			}
		} else {
			return fmt.Errorf("ship %d is not in base", shipId)
		}
	} else {
		return fmt.Errorf("that should not be possible. Can't find base belonging to player %d. Please contact organizers", p.Index)
	}
	return nil
}
