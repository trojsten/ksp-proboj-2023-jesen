#include "common.h"
#include "proboj.h"
using namespace std;

World world;

bool condition(XY a,XY b){
	return world.mapa.can_move(b);
}

int target = -1;

int main(){
	char dot;
	while(1){
		cin >> world >> dot;
		vector<Turn> turns;
		cerr << "HARBOR STORAGE" << endl;
		cerr << world.harbors[0].storage.resources << endl;
		if(target == -1){
			for(auto i : world.ships)
				if(!i.mine)
					target = i.index;
		}
		cerr << "TARGET " << target << endl;
		for(auto i : world.my_ships()){
			cerr << "SHIP ID " << i.index << " POSITION " << i.coords << endl;
			cerr << "SHIP CAPACITY " << i.stats.max_cargo << endl;
			cerr << i.resources << endl;
			if(target == -1) continue;
			Ship utocim;
			if(world.ship_by_id(target,utocim)){
				if(utocim.is_wreck){
					turns.push_back(LootTurn(i.index,target));
					target = -1;
				}
				else if(i.can_attack(utocim.coords)){
					turns.push_back(ShootTurn(i.index,target));
				}else{
					turns.push_back(MoveTurn(i.index,move_to(i,utocim.coords,condition)));
				}
			}
		}
		if(world.my_ships().size() < 2)
			turns.push_back(BuyTurn(ShipsEnum::BlackPearl));
		for(auto i : turns){
			cout<<i<<"\n";
			cerr<<i<<"\n";
		}
		cout << "." << endl;
	}
}
