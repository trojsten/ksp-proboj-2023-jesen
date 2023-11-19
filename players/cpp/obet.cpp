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
			for(int i = 0;i<world.ships.size();i++){
				if(!world.ships[i].mine)
					target = i;
			}
		}
		cerr << "TARGET " << target << endl;
		for(auto i : world.my_ships()){
			cerr << "SHIP ID " << i.index << " POSITION " << i.coords << endl;
			cerr << "SHIP CAPACITY " << i.stats.max_cargo << endl;
			cerr << i.resources.resources << endl;
			if(i.resources[ResourceEnum::Gold] == 0)
				turns.push_back(StoreTurn(i.index,-50));
			else
				turns.push_back(MoveTurn(i.index,move_to(i.coords,{50,50},condition,i.stats.max_move_range)));
		}
		if(world.my_ships().size() < 2)
			turns.push_back(BuyTurn(ShipsEnum::LargeMerchantShip));
		for(auto i : turns){
			cout<<i<<"\n";
			cerr<<i<<"\n";
		}
		cout << "." << endl;
	}
}
