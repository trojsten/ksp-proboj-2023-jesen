import sys
from typing import List

from ships import *


@dataclass
class XY:
    x: int
    y: int

    def dist(self, other: "XY"):
        return abs(self.x - other.x) + abs(self.y - other.y)

    def __hash__(self):
        return hash((self.x, self.y))

    def __str__(self):
        return f"({self.x} {self.y})"


class Turn(ABC):
    @abstractmethod
    def __str__(self):
        raise NotImplementedError()


class MoveTurn(Turn):

    def __init__(self, ship_id: int, coords: XY):
        self.ship_id = ship_id
        self.coords = coords

    def __str__(self):
        print(f"MOVE {self.ship_id} {self.coords.x} {self.coords.y}")


class TradeTurn(Turn):

    def __init__(self, ship_id: int, resource: ResourceEnum, amount: int):
        self.ship_id = ship_id
        self.resource = resource
        self.amount = amount

    def __str__(self):
        print(f"TRADE {self.ship_id} {self.resource} {self.amount}")


class LootTurn(Turn):

    def __init__(self, ship_id: int):
        self.ship_id = ship_id

    def __str__(self):
        print(f"LOOT {self.ship_id}")


class ShootTurn(Turn):

    def __init__(self, ship_id: int, target: int):
        self.ship_id = ship_id
        self.target = target

    def __str__(self):
        print(f"SHOOT {self.ship_id} {self.target}")


class BuyTurn(Turn):

    def __init__(self, ship_id: int):
        self.ship_id = ship_id

    def __str__(self):
        print(f"BUY {self.ship_id}")


class StoreTurn(Turn):

    def __init__(self, amount: int):
        self.amount = amount

    def __str__(self):
        print(f"STORE {self.amount}")


class TileEnum(enum.Enum):
    TILE_WATER = 0
    TILE_GROUND = 1
    TILE_HARBOR = 2
    TILE_BASE = 3


@dataclass
class Tile:
    type: TileEnum
    PlayerIndex: int


@dataclass
class Map:
    width: int
    height: int
    tiles: List[List[Tile]]

    @classmethod
    def read_map(cls) -> "Map":
        return Map(0, 0, [])


class Player:
    """
    Trieda, ktorá reprezentuje bežného hráča v hre.
    * idx - jeho idčko
    * ...
    """

    def __init__(self):
        self.id: int = 0

    @classmethod
    def read_player(cls) -> "Player":
        # TODO
        player = Player()
        return player

    def __eq__(self, other):
        return self.id == other.id

    def __hash__(self):
        return self.id


class MyPlayer(Player):
    """
    Trieda, ktorá reprezentuje Tvojho hráča v hre.
    * ...
    """

    def __init__(self):
        super().__init__()
        # TODO add

    @classmethod
    def read_myplayer(cls) -> "MyPlayer":
        myplayer = MyPlayer()
        return myplayer


class ProbojPlayer:
    """
    Táto trieda vykonáva ťahy v hre
    * world - objekt, ktorý reprezentuje svet
    * myself - Ty
    * _myself - Tvoje id
    * players - `dictionary` hráčov `{id: Player}`
    """

    def __init__(self):
        self.map: Map
        self.myself: MyPlayer
        self._myself: int
        self.players: dict[int: Player]

    @staticmethod
    def log(*args):
        """
        Vypíše dáta do logu. Syntax je rovnaká ako print().
        """
        print(*args, file=sys.stderr)

    def _read_myself(self):
        """
        Načíta info o sebe
        """
        self.myself = MyPlayer.read_myplayer()
        self._myself = self.myself.id

    def _read_players(self):
        """
        Načíta informácie o ostatných hráčoch
        """
        # TODO

    def _read_turn(self):
        """
        Načíta vstup pre hráča
        """
        self.map = Map.read_map()
        self._read_myself()
        self._read_players()
        input()
        input()

    def _send_turns(self, turns: List[Turn]):
        """
        Odošle ťah serveru.
        """
        for turn in turns:
            print(turn)
        print('.')

    def make_turn(self) -> List[Turn]:
        """
        Vykoná ťah.
        Funkcia vracia List objektov Turn
        """
        raise NotImplementedError()

    def run(self):
        """
        Hlavný cyklus hry.
        """
        while True:
            self._read_turn()
            turns = self.make_turn()
            self._send_turns(turns)
