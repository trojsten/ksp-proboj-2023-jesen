package main

type Score struct {
	GoldEarned          int `json:"gold_earned"`
	CurrentGold         int `json:"current_gold"`
	Kills               int `json:"kills"`
	SellsToHarbor       int `json:"sells_to_harbor"`
	PurchasesFromHarbor int `json:"purchases_from_harbor"`
	FinalScore          int `json:"final_score"`
}

func (score *Score) newKill() {
	score.Kills += 1
	score.updateFinalScore()
}

func (score *Score) newSell(amount int) {
	score.SellsToHarbor += amount
	score.updateFinalScore()
}

func (score *Score) newPurchase(amount int) {
	score.PurchasesFromHarbor += amount
	score.updateFinalScore()
}

func (score *Score) newGoldEarned(amount int) {
	score.GoldEarned += amount
	score.updateFinalScore()
}

func (score *Score) updateCurrentGold(amount int) {
	score.CurrentGold = amount
	score.updateFinalScore()
}

func (score *Score) updateFinalScore() {
	score.FinalScore = (score.GoldEarned + score.CurrentGold/3 + score.Kills*500 + score.SellsToHarbor/5 + score.PurchasesFromHarbor/5) - NEW_PLAYER_GOLD
}
