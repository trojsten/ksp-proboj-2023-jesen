import sys
from typing import Tuple, Union

try:
    import ujson as json
except ImportError:
    print("fallback to default json", file=sys.stderr)
    import json
from abc import ABC, abstractmethod
from collections import deque

from ships import *


class Turn(ABC):
    @abstractmethod
    def __str__(self):
        raise NotImplementedError()


class MoveTurn(Turn):
    def __init__(self, ship_id: int, coords: XY):
        """
        Trieda, ktorou dávame lodi príkaz k presunu na súradnice.

        :param ship_id idčko lode, ktorej dávame príkaz
        :param coords kam sa má loď pohnúť
        """
        self.ship_id = ship_id
        self.coords = coords

    def __str__(self):
        return f"MOVE {self.ship_id} {self.coords.x} {self.coords.y}"

    def __repr__(self):
        return str(self)


class TradeTurn(Turn):
    def __init__(self, ship_id: int, resource: Union[ResourceEnum, int], amount: int):
        """
        Trieda, ktorou dávame lodi príkaz k obchodovaniu s prístavom.
        Loď musí byť na rovnakých súradniciach ako prístav.

        :param ship_id idčko lode, ktorej dávame príkaz
        :param resource akú surovinu máme predávať/nakupovať
        :param amount koľko suroviny ideme predávať/nakupovať. Kladné ak ideme nakupovať z prístavu, záporné ak predávame do prístavu
        """
        self.ship_id = ship_id
        self.resource = resource
        self.amount = amount

    def __str__(self):
        resource = self.resource
        if isinstance(resource, ResourceEnum):
            resource = resource.value
        return f"TRADE {self.ship_id} {resource} {self.amount}"

    def __repr__(self):
        return str(self)


class LootTurn(Turn):
    def __init__(self, ship_id: int, target: int):
        """
        Trieda, ktorou dávame lodi príkaz na ťaženie surovín z vraku lode.
        Loď musí byť na rovnakých súradniciach ako vrak.

        :param ship_id idčko lode, ktorej dávame príkaz
        :param target idčko lode, z ktorej chceme ťažiť suroviny
        """
        self.ship_id = ship_id
        self.target = target

    def __str__(self):
        return f"LOOT {self.ship_id} {self.target}"

    def __repr__(self):
        return str(self)


class ShootTurn(Turn):
    def __init__(self, ship_id: int, target: int):
        """
        Trieda, ktorou dávame lodi príkaz na strieľanie na inú loď.
        Loď musí mať dostrel na túto loď.

        :param ship_id idčko lode, ktorej dávame príkaz
        :param target idčko lode, na ktorú chceme strieľať.
        """
        self.ship_id = ship_id
        self.target = target

    def __str__(self):
        return f"SHOOT {self.ship_id} {self.target}"

    def __repr__(self):
        return str(self)


class BuyTurn(Turn):
    def __init__(self, ship: ShipsEnum):
        """
        Trieda, ktorou dávame príkaz na nákup novej lode.
        Na tento nákup musí mať hráč peniaze v základni (nestačí na lodiach).

        :param ship druh lode, ktorú chceme kúpiť
        """
        self.ship = ship

    def __str__(self):
        return f"BUY {self.ship.value}"

    def __repr__(self):
        return str(self)


class StoreTurn(Turn):
    def __init__(self, ship_id: int, amount: int):
        """
        Trieda, ktorou dávame príkaz na uloženie/vybratie zlata zo základne.
        Musíme mať loď v základni.

        :param ship_id idčko lode, ktorej dávame príkaz
        :param amount množstvo zlata, ktoré chceme uložiť/vybrať. Kladné ak ukladáma do základné, záporné ak vyberáme.
        """
        self.amount = amount
        self.ship_id = ship_id

    def __str__(self):
        return f"STORE {self.ship_id} {self.amount}"

    def __repr__(self):
        return str(self)


@dataclass
class Harbor:
    """
    Trieda, ktorá reprezentuje prístav v hre.
    * coords - poloha prístavu
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

    def resource_cost(self, resource: ResourceEnum):
        """
        :param resourse, ktore cenu chcem zistit
        :return cenu suroviny vo viditeľnom prístave.
        :return Ak je funkcia volaný pre neviditeľný prístav, vráti základnú cenu.
        """
        base_price = [1, 2, 5, 10, 3, 5, 2, 3, 1]
        if not self.visible:
            return base_price[resource.value]
        return int(
            min(100 / (self.storage[resource] + 3) + 1, 4) * base_price[resource.value]
        )

    def __str__(self):
        return f"Harbor({self.coords})"

    def __repr__(self):
        return ("" if self.visible else "invisible") + f"Harbor({self.coords})"


class TileEnum(enum.Enum):
    """
    Enum, ktorý hovorí o type políčka.
    """

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

    def __repr__(self):
        return (
            f"(Tile({self.type.name}, player:{self.index})"
            if self.index != -1
            else self.type.name[5:]
        )

    def __str__(self):
        tiles = ["~", "O", "H", "B"]
        return tiles[self.type.value]

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
    directions: Tuple[Tuple[int]] = ((1, 0), (-1, 0), (0, 1), (0, -1))

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
            "".join(str(tile) for tile in line) for line in self.tiles
        )

    def __repr__(self):
        return f"map {self.width}x{self.height}\n" + "\n".join(
            ",".join(repr(tile) for tile in line) for line in self.tiles
        )

    def tile_type_at(self, pos: XY) -> "TileEnum":
        """
        :param pos súradnice
        :return typ políčka na súradniciach `pos`
        """
        return self.tiles[pos.y][pos.x].type

    def neighbours(self, pos: XY) -> List[XY]:
        """
        :return pole políčok s ktorými `pos` susedí
        """
        out = []
        for dx, dy in self.directions:
            new_x, new_y = pos.x + dx, pos.y + dy
            if 0 <= new_x < self.height and 0 <= new_y < self.width:
                if self.tiles[new_y][new_x].type != TileEnum.TILE_GROUND:
                    out.append(XY(new_x, new_y))
        return out


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
        self.base = XY(0, 0)
        self._myself: int

    def mine_ships(self) -> List[Ship]:
        """
        :return: pole lodí, ktoré patria tebe
        """
        return [i for i in self.ships if i.mine]

    def is_occupied_by_ship(self, coord: XY) -> bool:
        """
        :param coord o akom políčku zisťujeme informáciu
        :return či je políčko `coord` obsadené loďou
        """
        for i in self.mine_ships():
            if coord == i.coords:
                return True
        return False

    @staticmethod
    def log(*args):
        """
        Vypíše dáta do logu. Syntax je rovnaká ako `print()`.
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
        self._read_myself(state["index"], state["gold"])
        if state["map"] is not None and state["map"]["tiles"] is not None:
            self.map = Map.read_map(state["map"])
            for i in range(self.map.width):
                for j in range(self.map.height):
                    if (
                        self.map.tile_type_at(XY(i, j)) == TileEnum.TILE_BASE
                        and self.map.tiles[j][i].index == self.myself.index
                    ):
                        self.base = XY(i, j)

        self.harbors = Harbor.read_harbors(state["harbors"])
        self.ships = Ship.read_ships(state["ships"])
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


class Utils:
    @classmethod
    def bfs_path(cls, start: XY, goal: XY, mapa: Map) -> Union[List[XY], None]:
        """
        :return pole políčok, cez ktoré treba ísť, alebo `None`, ak sa nedá dosiahnuť
        """
        q = deque()
        q.append(start)
        parent = {start: None}

        while q:
            curr = q.popleft()
            if curr == goal:
                path = []
                while curr != start:
                    path.append(curr)
                    curr = parent[curr]
                path.reverse()
                return path

            for node in mapa.neighbours(curr):
                if node not in parent:
                    q.append(node)
                    parent[node] = curr
