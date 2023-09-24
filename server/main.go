package main

import "github.com/trojsten/ksp-proboj/client"

func main() {
	runner := client.NewRunner()
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
