#include "common.h"
#include "utils.h"
using namespace std;

World world;

// funkcia na pohyb moze vyzerat napriklad takto
// z a do b sa da dostat ak b je volne policko
bool condition(XY a, XY b) { return world.mapa.can_move(b); }

vector<Turn> do_turn() {
    // sem ide vas kod
    vector<Turn> turns;
    for (Ship ship : world.my_ships()) {
        if (ship.coords == world.my_base && ship.resources[ResourceEnum::Gold] == 0)
            turns.push_back(StoreTurn(ship.index, -5));
        else
            // takto sa pohybuje smerom na suradnice 50 50
            turns.push_back(MoveTurn(ship.index, move_to(ship, {50, 50}, condition)));
    }
    if (world.my_ships().size() < 3)
        // takto mozes kupit lod
        turns.push_back(BuyTurn(ShipsEnum::Cln));
    cerr << "Takto mozete vypisovat do logov" << endl;
    cerr << turns << endl;
    return turns;
}

int main() {
    char dot;
    while (1) {
        cin >> world >> dot;
        for (Turn turn : do_turn())
            cout << turn << "\n";
        cout << dot << endl;
    }
}
