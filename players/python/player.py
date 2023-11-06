#!/bin/env python3
import random

from proboj import *


class MyPlayer(ProbojPlayer):
    def make_turn(self):
        self.log(self.harbors)
        self.log(self.ships)
        self.log(self.myself, self.myself.gold)
        # self.log(self.map)
        # self.log(self.map.tiles[0][0])

        moves = [
            BuyTurn(ShipsEnum.Cln),
            # ShootTurn(0, 1),
            LootTurn(0, 1),
        ]

        for ship in self.ships:
            moves.append(random.choice([MoveTurn(ship.index, coord) for coord in ship.coords.nbhs()]))

        self.log(moves)

        return moves


if __name__ == "__main__":
    p = MyPlayer()
    p.run()
