import json
import sys

from ships import *

_input = input


def lepsiInput():
    d = _input()
    print("citame:", d[:10], file=sys.stderr)
    return d


input = lepsiInput

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
        return f"MOVE {self.ship_id} {self.coords.x} {self.coords.y}"


class TradeTurn(Turn):

    def __init__(self, ship_id: int, resource: ResourceEnum, amount: int):
        self.ship_id = ship_id
        self.resource = resource
        self.amount = amount

    def __str__(self):
        return f"TRADE {self.ship_id} {self.resource} {self.amount}"


class LootTurn(Turn):

    def __init__(self, ship_id: int):
        self.ship_id = ship_id

    def __str__(self):
        return f"LOOT {self.ship_id}"


class ShootTurn(Turn):

    def __init__(self, ship_id: int, target: int):
        self.ship_id = ship_id
        self.target = target

    def __str__(self):
        return f"SHOOT {self.ship_id} {self.target}"


class BuyTurn(Turn):

    def __init__(self, ship_id: int):
        self.ship_id = ship_id

    def __str__(self):
        return f"BUY {self.ship_id}"


class StoreTurn(Turn):

    def __init__(self, amount: int):
        self.amount = amount

    def __str__(self):
        return f"STORE {self.amount}"


@dataclass
class Harbor:
    x: int
    y: int
    production: Resources
    storage: Resources
    visible: bool

    @classmethod
    def read_harbors(cls, state_harbors: dict) -> List["Harbor"]:
        harbors = []
        for h in state_harbors:
            harbors.append(Harbor(**h))
        return harbors


class TileEnum(enum.Enum):
    TILE_WATER = 0
    TILE_GROUND = 1
    TILE_HARBOR = 2
    TILE_BASE = 3


@dataclass
class Tile:
    type: TileEnum
    index: int


@dataclass
class Map:
    width: int
    height: int
    tiles: List[List[Tile]]

    @classmethod
    def read_map(cls, state_map) -> "Map":
        return Map(**state_map)


class Player:
    """
    Trieda, ktorá reprezentuje Teba v hre.
    * idx - tvoje idčko
    * gold - koľko peňazí máš
    """

    def __init__(self, index: int, gold: int):
        self.index: int = index
        self.gold: int = gold


class ProbojPlayer:
    """
    Táto trieda vykonáva ťahy v hre
    * map - objekt, ktorý reprezentuje svet
    * harbors - pole objektov, ktoré reprezentujú prístavy
    * ships - pole objektov, ktoré reprezentujú lode
    * myself - Ty
    * _myself - Tvoje id
    """

    def __init__(self):
        self.map: Map
        self.harbors: List[Harbor]
        self.ships: List[Ship]
        self.myself: Player
        self._myself: int

    @staticmethod
    def log(*args):
        """
        Vypíše dáta do logu. Syntax je rovnaká ako print().
        """
        print(*args, file=sys.stderr)

    def _read_myself(self, index: int, gold: int):
        """
        Načíta info o sebe
        """
        self.myself = Player(index, gold)
        self._myself = self.myself.index

    def _read_turn(self):
        """
        Načíta vstup pre hráča
        """
        state = json.loads(input())
        self.map = Map.read_map(state['map'])
        self.harbors = Harbor.read_harbors(state['harbors'])
        self.ships = Ship.read_ships(state['ships'])
        self._read_myself(state["index"], state["gold"])
        input()
        # input()

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
