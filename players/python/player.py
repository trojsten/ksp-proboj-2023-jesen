#!/bin/env python3
from proboj import * 

class MyPlayer(ProbojPlayer):
    def make_turn(self):
        return []

if __name__ == "__main__":
    p = MyPlayer()
    p.run()
