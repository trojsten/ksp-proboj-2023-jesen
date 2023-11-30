package main

type Statistics struct {
	Kills           map[string]int `json:"kills"`
	Damage          map[string]int `json:"damage"`
	SellsByType     map[int]int    `json:"sells_by_type"`
	PurchasesByType map[int]int    `json:"purchases_by_type"`
	TimeByShip      map[string]int `json:"time_by_ship"`
	TimeOfResponses int64          `json:"time_of_responses"`
}

func (statistics *Statistics) newKill(playerName string) {
	_, ok := statistics.Kills[playerName]
	if ok {
		statistics.Kills[playerName] += 1
	} else {
		statistics.Kills[playerName] = 1
	}
}

func (statistics *Statistics) addDamage(playerName string, damage int) {
	_, ok := statistics.Damage[playerName]
	if ok {
		statistics.Damage[playerName] += damage
	} else {
		statistics.Damage[playerName] = damage
	}
}

func (statistics *Statistics) newSell(resourceId int, amount int) {
	_, ok := statistics.SellsByType[resourceId]
	if ok {
		statistics.SellsByType[resourceId] += amount
	} else {
		statistics.SellsByType[resourceId] = amount
	}
}

func (statistics *Statistics) newPurchase(resourceId int, amount int) {
	_, ok := statistics.PurchasesByType[resourceId]
	if ok {
		statistics.PurchasesByType[resourceId] += amount
	} else {
		statistics.PurchasesByType[resourceId] = amount
	}
}

func (statistics *Statistics) addTimeByShip(shipType ShipType) {
	_, ok := statistics.TimeByShip[shipType.Name()]
	if ok {
		statistics.TimeByShip[shipType.Name()] += 1
	} else {
		statistics.TimeByShip[shipType.Name()] = 1
	}
}

func (statistics *Statistics) addTimeOfResponse(time int64) {
	statistics.TimeOfResponses += time
}
