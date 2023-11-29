#!/bin/env python3
import random

from proboj import *


class MyPlayer(ProbojPlayer):
    def make_turn(self) -> List[Turn]:
        """
        Sem patrí váš kod. Táto metóda by mala vrátiť pole vašich ťahov.
        """
        self.log(self.harbors)
        self.log(self.ships)
        self.log(self.myself, self.myself.gold)
        # self.log(self.map)
        self.log(self.is_occupied_by_ship(XY(0, 0)))

        moves = [BuyTurn(ShipsEnum.Cln)]
        self.log(self.base)
        for ship in self.mine_ships():
            path = Utils.bfs_path(ship.coords, self.harbors[0].coords, self.map)
            if path:
                moves.append(MoveTurn(ship.index, path[0]))
        self.log(moves)

        return moves


if __name__ == "__main__":
    p = MyPlayer()
    p.run()
