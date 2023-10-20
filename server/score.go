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

func (score *Score) newSell() {
	score.SellsToHarbor += 1
	score.updateFinalScore()
}

func (score *Score) newPurchase() {
	score.PurchasesFromHarbor += 1
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
	// TODO real score formula
	score.FinalScore = score.GoldEarned + score.CurrentGold + score.Kills + score.SellsToHarbor + score.PurchasesFromHarbor
}
