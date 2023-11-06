import sys
try:
    import ujson as json
except ImportError:
    print("fallback to default json", file=sys.stderr)
    import json
from abc import ABC, abstractmethod

from ships import *


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

    def __repr__(self):
        return str(self)


class TradeTurn(Turn):
    def __init__(self, ship_id: int, resource: ResourceEnum, amount: int):
        self.ship_id = ship_id
        self.resource = resource
        self.amount = amount

    def __str__(self):
        return f"TRADE {self.ship_id} {self.resource} {self.amount}"

    def __repr__(self):
        return str(self)


class LootTurn(Turn):
    def __init__(self, ship_id: int, target: int):
        self.ship_id = ship_id
        self.target = target

    def __str__(self):
        return f"LOOT {self.ship_id} {self.target}"

    def __repr__(self):
        return str(self)


class ShootTurn(Turn):
    def __init__(self, ship_id: int, target: int):
        self.ship_id = ship_id
        self.target = target

    def __str__(self):
        return f"SHOOT {self.ship_id} {self.target}"

    def __repr__(self):
        return str(self)


class BuyTurn(Turn):
    def __init__(self, ship: ShipsEnum):
        self.ship = ship

    def __str__(self):
        return f"BUY {self.ship.value}"

    def __repr__(self):
        return str(self)


class StoreTurn(Turn):
    def __init__(self, amount: int):
        self.amount = amount

    def __str__(self):
        return f"STORE {self.amount}"

    def __repr__(self):
        return str(self)


@dataclass
class Harbor:
    """
    Trieda, ktorá reprezentuje prístav v hre.
    * XY - poloha prístavu
    * production - koľko surovín prístav vygeneruje každý ťah.
    * storage - koľko surovín má prístav na sklade.
    * visible - či vidíš informácie o tomto prístave. Ak `False`, sú tam 0, alebo náhodné čísla.
    """

    coords: XY
    production: Resources
    storage: Resources
    visible: bool

    @classmethod
    def read_harbors(cls, state_harbors: dict) -> List["Harbor"]:
        harbors = []
        for h in state_harbors:
            harbor = Harbor(
                XY(h["x"], h["y"]),
                Resources(h["production"]),
                Resources(h["storage"]),
                h["visible"],
            )
            harbors.append(harbor)
        return harbors

    def __str__(self):
        return f"Harbor({self.coords})"

    def __repr__(self):
        return "" if self.visible else "invisible" + f"Harbor({self.coords})"


class TileEnum(enum.Enum):
    TILE_WATER = 0
    TILE_GROUND = 1
    TILE_HARBOR = 2
    TILE_BASE = 3


@dataclass
class Tile:
    """
    Trieda, ktorá reprezentuje mapu v hre.
    * type - druh dlaždice. Jeden z `TileEnum`
    * index - id hráča, ktorému patrí, inak -1
    """

    type: TileEnum
    index: int

    def __str__(self):
        return (
            f"(Tile({self.type.name}, player:{self.index})"
            if self.index != -1
            else self.type.name[5:]
        )

    @classmethod
    def read_tile(cls, tile):
        return Tile(TileEnum(tile["type"]), tile["index"])


@dataclass
class Map:
    """
    Trieda, ktorá reprezentuje mapu v hre.
    * width - šírka mapy
    * height - výška mapy
    * tiles - 2D pole `Tiles`, teda samotná mapa
    """

    width: int
    height: int
    tiles: List[List[Tile]]

    @classmethod
    def read_map(cls, state_map) -> "Map":
        tiles = []
        for line in state_map["tiles"]:
            tiles.append([])
            for cell in line:
                tiles[-1].append(Tile.read_tile(cell))
        return Map(state_map["width"], state_map["height"], tiles)

    def __str__(self):
        return f"map {self.width}x{self.height}\n" + "\n".join(
            ",".join(str(tile) for tile in line) for line in self.tiles
        )

    def tile_type_at(self, pos: XY) -> "Tile":
        return self.tiles[pos.y][pos.x].type


class Player:
    """
    Trieda, ktorá reprezentuje Teba v hre.
    * idx - tvoje idčko
    * gold - koľko peňazí máš
    """

    def __init__(self, index: int, gold: int):
        self.index: int = index
        self.gold: int = gold

    def __str__(self):
        return f"Player({self.index})"


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
        print(*args, file=sys.stderr, flush=True)

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
        inp = input()
        print("size of json", len(inp), file=sys.stderr, flush=True)
        state = json.loads(inp)
        if state["map"] is not None and state["map"]["tiles"] is not None:
            self.map = Map.read_map(state["map"])
        self.harbors = Harbor.read_harbors(state["harbors"])
        self.ships = Ship.read_ships(state["ships"])
        self._read_myself(state["index"], state["gold"])
        input()
        # input()

    def _send_turns(self, turns: List[Turn]):
        """
        Odošle ťah serveru.
        """
        for turn in turns:
            print(turn)
        print(".")

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
