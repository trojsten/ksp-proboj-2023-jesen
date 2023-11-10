#include "common.h"
using namespace std;

World world;

int main(){
	char dot;
	while(1){
		cin >> world;
		cin >> dot;
		cerr << world.mapa << endl;
		cout << "." << endl;
	}
}
