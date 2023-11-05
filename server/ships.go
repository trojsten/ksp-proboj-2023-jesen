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
		MaxMoveRange: 2,
		MaxCargo:     10,
		Price:        10,
		Yield:        0.2,
		Class:        SHIP_TRADE,
	}
}
func (c Cln) Name() string {
	return "Cln"
}

type Plt struct {
}

func (c Plt) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    15,
		Damage:       1,
		Range:        2,
		MaxMoveRange: 1,
		MaxCargo:     50,
		Price:        30,
		Yield:        0.2,
		Class:        SHIP_TRADE,
	}
}

func (p Plt) Name() string {
	return "Plt"
}

type SmallMerchantShip struct {
}

func (c SmallMerchantShip) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    30,
		Damage:       1,
		Range:        3,
		MaxMoveRange: 3,
		MaxCargo:     50,
		Price:        100,
		Yield:        0.2,
		Class:        SHIP_TRADE,
	}
}

func (c SmallMerchantShip) Name() string {
	return "SmallMerchantShip"
}

type LargeMerchantShip struct {
}

func (c LargeMerchantShip) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    50,
		Damage:       2,
		Range:        4,
		MaxMoveRange: 2,
		MaxCargo:     100,
		Price:        200,
		Yield:        0.2,
		Class:        SHIP_TRADE,
	}
}

func (c LargeMerchantShip) Name() string {
	return "LargeMerchantShip"
}

type SomalianPirateShip struct {
}

func (c SomalianPirateShip) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    10,
		Damage:       3,
		Range:        2,
		MaxMoveRange: 3,
		MaxCargo:     5,
		Price:        15,
		Yield:        0.5,
		Class:        SHIP_ATTACK,
	}
}

func (c SomalianPirateShip) Name() string {
	return "SomalianPirateShip"
}

type BlackPearl struct {
}

func (c BlackPearl) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    50,
		Damage:       5,
		Range:        4,
		MaxMoveRange: 2,
		MaxCargo:     30,
		Price:        50,
		Yield:        0.5,
		Class:        SHIP_ATTACK,
	}
}

func (c BlackPearl) Name() string {
	return "BlackPearl"
}

type SniperAttackShip struct {
}

func (c SniperAttackShip) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    5,
		Damage:       8,
		Range:        3,
		MaxMoveRange: 1,
		MaxCargo:     10,
		Price:        30,
		Yield:        0.5,
		Class:        SHIP_ATTACK,
	}
}

func (c SniperAttackShip) Name() string {
	return "SniperAttackShip"
}

type LooterScooter struct {
}

func (c LooterScooter) Stats() ShipStats {
	return ShipStats{
		MaxHealth:    5,
		Damage:       0,
		Range:        5,
		MaxMoveRange: 4,
		MaxCargo:     30,
		Price:        50,
		Yield:        0.8,
		Class:        SHIP_LOOT,
	}
}

func (c LooterScooter) Name() string {
	return "LooterScooter"
}
