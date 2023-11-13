#include "common.h"
#include "proboj.h"
using namespace std;

World world;

bool condition(POINT a,POINT b){
	return world.mapa.can_move(b);
}

XY base = {-1,-1};

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
		for(auto i : world.ships){
			if(i.mine){
				cerr << "SHIP ID " << i.index << " POSITION " << i.coords << endl;
				cerr << "SHIP CAPACITY " << i.stats.max_cargo << endl;
				cerr << i.resources.resources << endl;
				if(i.resources[ResourceEnum::Gold] == 0)
					turns.push_back(StoreTurn(i.index,-10));
				else if(i.coords == world.harbors[0].coords){
					turns.push_back(TradeTurn(i.index,ResourceEnum::Wood,5));
				}else
					turns.push_back(MoveTurn(i.index,move_to(i.coords,world.harbors[0].coords,condition)));
			}
		}
		XY a = POINT{3,2};
		turns.push_back(BuyTurn(ShipsEnum::Cln));
		for(auto i : turns){
			cout<<i<<"\n";
			cerr<<i<<"\n";
		}
		cout << "." << endl;
	}
}
