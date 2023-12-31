package main

import (
	"fmt"
	"github.com/trojsten/ksp-proboj/client"
	"math/rand"
	"time"
)

var BuildTime = "unknown"

func main() {
	runner := client.NewRunner()
	runner.Log(fmt.Sprintf("Started server, built on %v", BuildTime))
	runner.Log(fmt.Sprintf("started"))
	seed := time.Now().UnixMilli()
	rand.Seed(seed)
	runner.Log(fmt.Sprintf("seed %d", seed))
	players, config := runner.ReadConfig()

	game := Game{Runner: runner}
	for i, player := range players {
		game.Players = append(game.Players, NewPlayer(&game, i, player))
	}
	game.Ships = map[int]*Ship{}

	err := game.LoadMap(config)
	if err != nil {
		runner.Log(err.Error())
		panic(err)
	}

	err = game.Run()
	if err != nil {
		runner.Log(err.Error())
		panic(err)
	}
	runner.Log("Game successfully ended")

	scores := map[string]int{}
	for _, player := range game.Players {
		scores[player.Name] = player.Score.FinalScore
	}
	runner.Scores(scores)

	err = game.SaveStats()
	if err != nil {
		game.Runner.Log(fmt.Sprintf("error while saving stats: %s", err.Error()))
	}

	runner.End()
}
