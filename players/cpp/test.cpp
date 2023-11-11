#include "common.h"
using namespace std;

World world;

int main(){
	char dot;
	while(1){
		cin >> world >> dot;
		vector<Turn> turns;
		for(auto i : world.ships){
			if(i.mine){
				turns.push_back(MoveTurn(i.index,i.coords + XY(1,0)));
			}
		}
		turns.push_back(BuyTurn(Plt));
		for(auto i : turns){
			cout<<i<<"\n";
			cerr<<i<<"\n";
		}
		cout << "." << endl;
	}
}
