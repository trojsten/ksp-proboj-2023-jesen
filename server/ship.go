package main

type ShipType interface {
	Stats() ShipStats
	Name() string
}

type Ship struct {
	Id          int       `json:"index"`
	PlayerIndex int       `json:"player_index"`
	Type        ShipType  `json:"type,omitempty"`
	X           int       `json:"x"`
	Y           int       `json:"y"`
	Health      int       `json:"health"`
	IsWreck     bool      `json:"is_wreck"`
	Resources   Resources `json:"resources"`
}

type ShipClass int

const (
	SHIP_TRADE ShipClass = iota
	SHIP_ATTACK
	SHIP_LOOT
)

type ShipStats struct {
	MaxHealth    int       `json:"max_health"`
	Damage       int       `json:"damage"`
	Range        int       `json:"range"`
	MaxMoveRange int       `json:"max_move_range"`
	MaxCargo     int       `json:"max_cargo"`
	Price        int       `json:"price"`
	Yield        float32   `json:"yield_frac"`
	Class        ShipClass `json:"ship_class"`
}

type Resources struct {
	Wood      int `json:"wood"`
	Stone     int `json:"stone"`
	Iron      int `json:"iron"`
	Gem       int `json:"gem"`
	Wool      int `json:"wool"`
	Hide      int `json:"hide"`
	Wheat     int `json:"wheat"`
	Pineapple int `json:"pineapple"`
	Gold      int `json:"gold"`
}

type ResourceType int

const (
	RESOURCE_WOOD ResourceType = iota
	RESOURCE_STONE
	RESOURCE_IRON
	RESOURCE_GEM
	RESOURCE_WOOL
	RESOURCE_HIDE
	RESOURCE_WHEAT
	RESOURCE_PINEAPPLE
	RESOURCE_GOLD
)

func (r *Resources) Resource(id ResourceType) *int {
	switch id {
	case RESOURCE_WOOD:
		return &r.Wood
	case RESOURCE_STONE:
		return &r.Stone
	case RESOURCE_IRON:
		return &r.Iron
	case RESOURCE_GEM:
		return &r.Gem
	case RESOURCE_WOOL:
		return &r.Wool
	case RESOURCE_HIDE:
		return &r.Hide
	case RESOURCE_WHEAT:
		return &r.Wheat
	case RESOURCE_PINEAPPLE:
		return &r.Pineapple
	case RESOURCE_GOLD:
		return &r.Gold
	default:
		return nil
	}
}

func (r *Resources) countResources() int {
	return r.Wood + r.Stone + r.Iron + r.Gem + r.Wool + r.Hide + r.Wheat + r.Pineapple
}

func (r *Resources) empty() bool {
	return r.Wood+r.Stone+r.Iron+r.Gem+r.Wool+r.Hide+r.Wheat+r.Pineapple+r.Gold == 0
}

func ShipAt(g *Game, x int, y int) *Ship {
	for i, ship := range g.Ships {
		if ship.X == x && ship.Y == y {
			return g.Ships[i]
		}
	}
	return nil
}
