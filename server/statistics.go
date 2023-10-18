package main

type Statistics struct {
	Kills               int
	SellsToHarbor       int
	PurchasesFromHarbor int
}

func (statistics Statistics) newKill(playerIndex int) {
	statistics.Kills += 1
	// TODO
}

func (statistics Statistics) addDamage(playerIndex int, damage int) {
	statistics.Kills += 1
	// TODO
}

func (statistics Statistics) newSell(resourceId int, amount int) {
	statistics.SellsToHarbor += 1
	// TODO
}

func (statistics Statistics) newPurchase(resourceId int, amount int) {
	statistics.PurchasesFromHarbor += 1
	// TODO
}
