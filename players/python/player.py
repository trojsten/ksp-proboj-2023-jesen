#!/bin/env python3
import random

from proboj import *


class MyPlayer(ProbojPlayer):
    def make_turn(self):
        self.log(self.harbors)
        self.log(self.ships)
        self.log(self.myself, self.myself.gold)
        # self.log(self.map)

        moves = [
            BuyTurn(ShipsEnum.Cln),
        ]

        for ship in self.mine_ships():
            options = []
            for coord in Utils.neighbours(ship.coords, self.map):
                options.append(MoveTurn(ship.index, coord))
            moves.append(random.choice(options))

        self.log(moves)

        return moves


if __name__ == "__main__":
    p = MyPlayer()
    p.run()
