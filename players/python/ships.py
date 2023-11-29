import enum
from dataclasses import dataclass
from typing import List


@dataclass
class XY:
    """
    Trieda, ktorá reprezentuje súradnicu.
    :attr x súradnica
    :attr y súradnica
    """

    x: int
    y: int

    def dist(self, other: "XY"):
        """
        Vypočíta vzdielanosť k inej súradnici
        :param other druhá súradnica
        :return vzdialenosť
        """
        return abs(self.x - other.x) + abs(self.y - other.y)

    def __hash__(self):
        return hash((self.x, self.y))

    def __str__(self):
        return f"({self.x} {self.y})"

    def __repr__(self):
        return f"({self.x} {self.y})"


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
    def __init__(self, resources: dict[str, int]):
        assert len(resources) == len(ResourceEnum)
        self.resources = [
            resources["wood"],
            resources["stone"],
            resources["iron"],
            resources["gem"],
            resources["wool"],
            resources["hide"],
            resources["wheat"],
            resources["pineapple"],
            resources["gold"],
        ]

    def __getitem__(self, key: ResourceEnum):
        return self.resources[key.value]

    def __str__(self):
        return str(self.resources)

    def __repr__(self):
        return str(self.resources)


class ShipClass(enum.Enum):
    SHIP_TRADE = 0
    SHIP_ATTACK = 1
    SHIP_LOOT = 2


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
    """
    Trieda, ktorá reprezentuje loď.
    :attr index idčko lode
    :attr player_index komu loď patrí
    :attr coords kde sa loď nachádza
    :attr healt koľko má života
    :attr is_wreck či je to vrak
    :attr resources aké resourcies má loď. Ak ju nevidíš, tak nulové.
    :attr resources aké stats má loď. Ak ju nevidíš, tak nulové.
    :attr mine či je loď tvoja
    """

    index: int
    player_index: int
    coords: XY
    health: int
    is_wreck: bool
    resources: Resources
    stats: ShipStats
    mine: bool

    def __lt__(self, other):
        return self.index.__lt__(other.index)

    @classmethod
    def read_ships(cls, state_ships: dict) -> List["Ship"]:
        ships = []
        for s in state_ships:
            coords = XY(s["x"], s["y"])
            del s["x"], s["y"]
            s["coords"] = coords
            ships.append(Ship(**s))
            ships[-1].resources = Resources(s["resources"])
            ships[-1].stats = ShipStats(**s["stats"])
            ships[-1].stats.ship_class = ShipClass(s["stats"]["ship_class"])
        ships.sort()
        return ships
