import enum
from abc import ABC, abstractmethod
from dataclasses import dataclass
from typing import List


class ResourceEnum(enum.Enum):
    Wood = 0
    Stone = 1
    Iron = 2
    Gem = 3
    Wool = 4
    Hide = 5
    Wheat = 6
    Pineapple = 7
    Gold = 8


class Resources:

    def __init__(self, resources: dict[str: int]):
        assert len(resources) == len(ResourceEnum)
        self.resources = list(resources.values())

    def __getitem__(self, key: ResourceEnum):
        return self.resources[key.value]

    def __str__(self):
        return str(self.resources)
    
    def __repr__(self):
        return str(self.resources)

class ShipClass(enum.Enum):
    SHIP_TRADE = 0
    SHIP_ATTACK = 1
    SHIP_LOOT = 1


@dataclass
class ShipStats:
    max_health: int
    damage: int
    range: int
    max_move_range: int
    max_cargo: int
    price: int
    yield_frac: float
    ship_class: ShipClass


class ShipsEnum(enum.Enum):
    Cln = 0
    Plt = 1
    SmallMerchantShip = 2
    LargeMerchantShip = 3
    SomalianPirateShip = 4
    BlackPearl = 5
    SniperAttackShip = 6
    LooterScooter = 7

@dataclass
class Ship:
    index: int
    player_index: int
    x: int
    y: int
    health: int
    is_wreck: bool
    resources: Resources
    stats: ShipStats
    mine: bool

    @classmethod
    def read_ships(cls, state_ships: dict) -> List["Ship"]:
        ships = []
        for s in state_ships:
            ships.append(Ship(**s))
            ships[-1].resources = Resources(s["resources"])
            ships[-1].stats = ShipStats(**s["stats"])
            ships[-1].stats.ship_class = ShipClass(s["stats"]["ship_class"])
        return ships
