#include "common.h"
#include "proboj.h"
using namespace std;

World world;

bool condition(POINT a,POINT b){
	return world.mapa.can_move(b);
}

Ship ship_by_id(int id){
	for(Ship i : world.ships)
		if(i.index == id)
			return i;
}

XY base = {-1,-1};
int target = -1;

int main(){
	char dot;
	while(1){
		cin >> world >> dot;
		vector<Turn> turns;
		cerr << "HARBOR STORAGE" << endl;
		cerr << world.harbors[0].storage.resources << endl;
		if(base.x == -1){
			for(int i = 0;i<world.mapa.height;i++){
				for(int j = 0;j<world.mapa.width;j++){
					Tile tile = world.mapa[{j,i}];
					if(tile.type == TileEnum::TILE_BASE && tile.index == world.index)
						base = {i,j};
				}
			}
		}
		if(target == -1){
			for(auto i : world.ships)
				if(!i.mine)
					target = i.index;
		}
		cerr << "TARGET " << target << endl;
		for(auto i : world.my_ships()){
			cerr << "SHIP ID " << i.index << " POSITION " << i.coords << endl;
			cerr << "SHIP CAPACITY " << i.stats.max_cargo << endl;
			cerr << i.resources.resources << endl;
			if(target == -1) continue;
			if(ship_by_id(target).is_wreck){
				turns.push_back(LootTurn(i.index,target));
				target = -1;
			}
			else if(i.can_attack(ship_by_id(target).coords)){
				turns.push_back(ShootTurn(i.index,target));
			}else{
				turns.push_back(MoveTurn(i.index,move_to(i.coords,world.ships[target].coords,condition,i.stats.max_move_range)));
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
