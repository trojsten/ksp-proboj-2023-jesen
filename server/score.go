package main

type Score struct {
	Kills               int `json:"kills"`
	SellsToHarbor       int `json:"sells_to_harbor"`
	PurchasesFromHarbor int `json:"purchases_from_harbor"`
	FinalScore          int `json:"final_score"`
}

func (score Score) newKill() {
	score.Kills += 1
	score.updateFinalScore()
}

func (score Score) newSell() {
	score.SellsToHarbor += 1
	score.updateFinalScore()
}

func (score Score) newPurchase() {
	score.PurchasesFromHarbor += 1
	score.updateFinalScore()
}

func (score Score) updateFinalScore() {
	score.FinalScore = score.Kills + score.SellsToHarbor + score.PurchasesFromHarbor
}
