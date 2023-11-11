#!/bin/env python3
import random

from proboj import *


class MyPlayer(ProbojPlayer):
    def make_turn(self):
        self.log(self.harbors)
        self.log(self.ships)
        self.log(self.myself, self.myself.gold)
        # self.log(self.map)
        self.log(self.is_occupied_by_ship(XY(0,0)))

        moves = [
            BuyTurn(ShipsEnum.Cln),
            StoreTurn(0, 1)
        ]

        for ship in self.mine_ships():
            options = []
            for coord in self.map.neighbours(ship.coords):
                options.append(MoveTurn(ship.index, coord))
            moves.append(random.choice(options))
            
        self.log(moves)

        return moves


if __name__ == "__main__":
    p = MyPlayer()
    p.run()
