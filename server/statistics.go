package main

type Statistics struct {
	Kills           map[int]int `json:"kills"`
	Damage          map[int]int `json:"damage"`
	SellsByType     map[int]int `json:"sells_by_type"`
	PurchasesByType map[int]int `json:"purchases_by_type"`
	TimeOfResponses int64       `json:"time_of_responses"`
}

func (statistics *Statistics) newKill(playerIndex int) {
	_, ok := statistics.Kills[playerIndex]
	if ok {
		statistics.Kills[playerIndex] += 1
	} else {
		statistics.Kills[playerIndex] = 1
	}
}

func (statistics *Statistics) addDamage(playerIndex int, damage int) {
	_, ok := statistics.Damage[playerIndex]
	if ok {
		statistics.Damage[playerIndex] += damage
	} else {
		statistics.Damage[playerIndex] = damage
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

func (statistics *Statistics) addTimeOfResponse(time int64) {
	statistics.TimeOfResponses += time
}
