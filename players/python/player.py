#!/bin/env python3
from proboj import * 

class MyPlayer(ProbojPlayer):
    def make_turn(self):
        self.log(self.harbors)
        self.log(self.map)
        self.log(self.map.tiles[0][0])
        return [
            MoveTurn(0, XY(0, 0)),
            MoveTurn(0, XY(0, 0)),
            MoveTurn(0, XY(0, 0)),
            None,
            12,
        ]

if __name__ == "__main__":
    p = MyPlayer()
    p.run()
