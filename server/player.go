package main

type Player struct {
	Index int
	Name  string
	Gold  int
	game  *Game
}

func NewPlayer(game *Game, idx int, name string) Player {
	return Player{
		Index: idx,
		Name:  name,
		game:  game,
	}
}

func (p *Player) Ships() []*Ship {
	var playersShips []*Ship
	for _, ship := range p.game.Ships {
		if ship.PlayerIndex == p.Index {
			playersShips = append(playersShips, ship)
		}
	}
	return playersShips
}

func (p *Player) Base() *Base {
	for _, base := range p.game.Bases {
		if base.PlayerIndex == p.Index {
			return &base
		}
	}
	return nil
}
