import enum
from abc import ABC, abstractmethod
from dataclasses import dataclass


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

    def __init__(self, resources):
        assert len(resources) == len(ResourceEnum)
        self.resources = resources[:]

    @staticmethod
    def read_stat_levels():
        resources = []
        for idx, lvl in enumerate(map(int, input().split())):
            resources.append(lvl)
        return Resources(resources)

    @staticmethod
    def read_stat_values():
        resources = []
        for idx, lvl in enumerate(map(float, input().split())):
            resources.append(lvl)
        return Resources(resources)

    def __getitem__(self, key: ResourceEnum):
        return self.resources[key.value]


class ShipClass(enum.Enum):
    SHIP_TRADE = 0
    SHIP_ATTACK = 1
    SHIP_LOOT = 1


@dataclass
class ShipStats:
    MaxHealth: int
    Damage: int
    Range: int
    MaxMoveRange: int
    MaxCargo: int
    Price: int
    Yield: float
    Class: ShipClass


@dataclass
class ShipType(ABC):

    @staticmethod
    def get_all_ships():
        return [
            Cln(),
            Plt(),
            SmallMerchantShip(),
            LargeMerchantShip(),
            SomalianPirateShip(),
            BlackPearl(),
            SniperAttackShip(),
            LooterScooter(),
        ]

    @staticmethod
    def get_ship(ship_id: int):
        for ship in ShipType.get_all_ships():
            if ship.ship_id == ship_id:
                return ship
        raise RuntimeError(f"Ship w/ {ship_id} not found")

    @abstractmethod
    def ship_id(self) -> int:
        pass

    @abstractmethod
    def stats_values(self) -> ShipStats:
        pass


@dataclass
class Cln(ShipType):
    ship_id = 0
    stats_values = ShipStats(0, 0, 0, 0, 0, 0, 0, ShipClass.SHIP_TRADE)


@dataclass
class Plt(ShipType):
    ship_id = 1
    stats_values = ShipStats(0, 0, 0, 0, 0, 0, 0, ShipClass.SHIP_TRADE)


@dataclass
class SmallMerchantShip(ShipType):
    ship_id = 2
    stats_values = ShipStats(0, 0, 0, 0, 0, 0, 0, ShipClass.SHIP_TRADE)


@dataclass
class LargeMerchantShip(ShipType):
    ship_id = 3
    stats_values = ShipStats(0, 0, 0, 0, 0, 0, 0, ShipClass.SHIP_TRADE)


@dataclass
class SomalianPirateShip(ShipType):
    ship_id = 4
    stats_values = ShipStats(0, 0, 0, 0, 0, 0, 0, ShipClass.SHIP_ATTACK)


@dataclass
class BlackPearl(ShipType):
    ship_id = 5
    stats_values = ShipStats(0, 0, 0, 0, 0, 0, 0, ShipClass.SHIP_ATTACK)


@dataclass
class SniperAttackShip(ShipType):
    ship_id = 6
    stats_values = ShipStats(0, 0, 0, 0, 0, 0, 0, ShipClass.SHIP_ATTACK)


@dataclass
class LooterScooter(ShipType):
    ship_id = 7
    stats_values = ShipStats(0, 0, 0, 0, 0, 0, 0, ShipClass.SHIP_LOOT)


class Ship:
    Id: int
    PlayerIndex: int
    Type: ShipType
    X: int
    Y: int
    Health: int
    IsWreck: bool
    Resources: Resources
