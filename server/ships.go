package main

var ships = []ShipType{
	Cln{},
	Plt{},
	SmallMerchantShip{},
	LargeMerchantShip{},
	SomalianPirateShip{},
	BlackPearl{},
	SniperAttackShip{},
	LooterScooter{},
}

type Cln struct {
}

func (c Cln) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    10,
		Damage:       1,
		Range:        1,
		MaxMoveRange: 1,
		MaxCargo:     5,
		Price:        10,
		Yield:        0.2,
		Class:        SHIP_TRADE,
	}
}

type Plt struct {
}

func (c Plt) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    25,
		Damage:       1,
		Range:        1,
		MaxMoveRange: 2,
		MaxCargo:     20,
		Price:        50,
		Yield:        0.2,
		Class:        SHIP_TRADE,
	}
}

type SmallMerchantShip struct {
}

func (c SmallMerchantShip) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    15,
		Damage:       1,
		Range:        1,
		MaxMoveRange: 3,
		MaxCargo:     10,
		Price:        60,
		Yield:        0.2,
		Class:        SHIP_TRADE,
	}
}

type LargeMerchantShip struct {
}

func (c LargeMerchantShip) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    50,
		Damage:       1,
		Range:        1,
		MaxMoveRange: 1,
		MaxCargo:     50,
		Price:        100,
		Yield:        0.2,
		Class:        SHIP_TRADE,
	}
}

type SomalianPirateShip struct {
}

func (c SomalianPirateShip) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    10,
		Damage:       3,
		Range:        1,
		MaxMoveRange: 2,
		MaxCargo:     10,
		Price:        10,
		Yield:        0.5,
		Class:        SHIP_ATTACK,
	}
}

type BlackPearl struct {
}

func (c BlackPearl) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    50,
		Damage:       5,
		Range:        1,
		MaxMoveRange: 4,
		MaxCargo:     30,
		Price:        50,
		Yield:        0.5,
		Class:        SHIP_ATTACK,
	}
}

type SniperAttackShip struct {
}

func (c SniperAttackShip) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    15,
		Damage:       10,
		Range:        5,
		MaxMoveRange: 2,
		MaxCargo:     5,
		Price:        50,
		Yield:        0.5,
		Class:        SHIP_ATTACK,
	}
}

type LooterScooter struct {
}

func (c LooterScooter) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    5,
		Damage:       1,
		Range:        1,
		MaxMoveRange: 4,
		MaxCargo:     20,
		Price:        30,
		Yield:        0.8,
		Class:        SHIP_LOOT,
	}
}
