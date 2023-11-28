#include "common.h"
#include "proboj.h"
using namespace std;

World world;

bool condition(XY a,XY b){
	return world.mapa.can_move(b);
}

int main(){
	char dot;
	while(1){
		cin >> world >> dot;
		vector<Turn> turns;
		cerr << "HARBOR STORAGE" << endl;
		cerr << world.harbors[0].storage.resources << endl;
		for(auto i : world.my_ships()){
			cerr << "SHIP ID " << i.index << " POSITION " << i.coords << endl;
			cerr << "SHIP CAPACITY " << i.stats.max_cargo << endl;
			cerr << i.resources.resources << endl;
			if(i.resources[ResourceEnum::Gold] == 0)
				turns.push_back(StoreTurn(i.index,-50));
			else
				turns.push_back(MoveTurn(i.index,move_to(i,{50,50},condition)));
		}
		if(world.my_ships().size() < 2)
			turns.push_back(BuyTurn(ShipsEnum::Cln));
		for(auto i : turns){
			cout<<i<<"\n";
		}
		cerr << turns << endl;
		cout << "." << endl;
	}
}