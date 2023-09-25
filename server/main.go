package main

import (
	"fmt"
	"github.com/trojsten/ksp-proboj/client"
	"math/rand"
	"time"
)

func main() {
	runner := client.NewRunner()
	runner.Log(fmt.Sprintf("started"))
	seed := time.Now().UnixMilli()
	rand.Seed(seed)
	runner.Log(fmt.Sprintf("seed %d", seed))
	players, config := runner.ReadConfig()

	game := Game{runner: runner}
	for i, player := range players {
		game.Players = append(game.Players, NewPlayer(&game, i, player))
	}

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
}