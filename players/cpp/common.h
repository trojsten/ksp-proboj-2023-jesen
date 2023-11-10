#ifndef COMMON_H
#define COMMON_H
#include<bits/stdc++.h>
#include "RSJparser.tcc"

namespace
{
using namespace std;
enum TileEnum{TILE_WATER, TILE_GROUND, TILE_HARBOR, TILE_BASE};
enum ResourceEnum{Wood, Stone, Iron, Gem, Wool, Hide, Wheat, Pineapple, Gold};

struct XY{
	int x,y;
	XY(int _x,int _y) : x(_x),y(_y){}
	XY(){}
	XY operator-(XY &other){return XY(this->x - other.x,this->y - other.y);}
	XY operator+(XY &other){return XY(this->x + other.x,this->y + other.y);}
	string repr(){return to_string(x) + " " + to_string(y);}
	friend ostream& operator<<(ostream &os,const XY &a){os << a.x << " " << a.y; return os;};
};

struct Turn{
	virtual string repr(){return "";}
	friend ostream& operator<<(ostream &os,Turn &a){
		os << a.repr();
		return os;
	}
};

struct MoveTurn : Turn{
	int ship_id;
	XY coords;
	MoveTurn(int ship_id,XY coords) : ship_id(ship_id),coords(coords){};
	virtual string repr(){
		return "MOVE " + to_string(ship_id) + " " + coords.repr();
	}
};

struct TradeTurn : Turn{
	int ship_id,amount;
	ResourceEnum resource;
	TradeTurn(int ship_id,ResourceEnum resource,int amount) : ship_id(ship_id),resource(resource),amount(amount){};
	virtual string repr(){
		return "TRADE " + to_string(ship_id) + " " + to_string(resource) + " " + to_string(amount);
	}
};

struct LootTurn : Turn{
	int ship_id,target;
	LootTurn(int ship_id,int target) : ship_id(ship_id),target(target){};
	virtual string repr(){
		return "LOOT " + to_string(ship_id) + " " + to_string(target);
	}
};

struct ShootTurn : Turn{
	int ship_id,target;
	ShootTurn(int ship_id,int target) : ship_id(ship_id),target(target){};
	virtual string repr(){
		return "SHOOT " + to_string(ship_id) + " " + to_string(target);
	}
};

//TODO
struct BuyTurn : Turn{
	int ship_id;
	BuyTurn(int ship_id) : ship_id(ship_id){};
	virtual string repr(){
		return "";
	}
};

struct StoreTurn : Turn{
	int ship_id,amount;
	StoreTurn(int ship_id,int amount) : ship_id(ship_id),amount(amount){};
	virtual string repr(){
		return "STORE " + to_string(ship_id) + " " + to_string(amount);
	}
};

struct Harbor{
	XY coords;
	unordered_map<string,int> production, storage;
	bool visible;
	Harbor(RSJobject harbor_dict){
		coords = XY(harbor_dict["x"].as<int>(),harbor_dict["y"].as<int>());
		//TODO mapa enum -> int
		production = harbor_dict["production"].as_map<int>();
		storage = harbor_dict["storage"].as_map<int>();
		visible = harbor_dict["visible"].as<bool>();
	}
};

vector<Harbor> load_harbors(RSJresource harbors){
	vector<Harbor> ret;
	for(auto har : harbors.as_array()){
		ret.push_back(Harbor(har.as_object()));
	}
	return ret;
}

struct Tile{
	TileEnum type;
	int index;
	Tile(RSJresource r) : type(static_cast<TileEnum>(r["type"].as<int>())),index(r["index"].as<int>()){}
	Tile(){}
	friend ostream& operator<<(ostream &os,const Tile &a){os << "Tile(" << a.type << "," << a.index << ")";return os;}
};

struct Map{
	int width,height;
	vector<vector<Tile>> tiles;
	Map(){}
	Map(RSJresource r){
		width = r["width"].as<int>();
		height = r["height"].as<int>();
		for(auto row : r["tiles"].as_array()){
			tiles.push_back({});
			for(auto cell : row.as_array()){
				tiles.back().push_back(Tile(cell));
			}
		}
	}

	bool inside(XY coords){
		return coords.x >= 0 && coords.x < width && coords.y >= 0 && coords.y < height;
	}

	bool can_move(XY coords){
		return inside(coords) && tiles[coords.y][coords.x].type == TILE_WATER;
	}

	friend ostream& operator<<(ostream &os,const Map &a){
		os << "Map " << a.width << " " << a.height << endl;
		for(auto i : a.tiles){
			for(auto j : i){
				switch(j.type){
					case TILE_WATER:
						os << "~";
						break;
					case TILE_GROUND:
						os << "O";
						break;
					case TILE_HARBOR:
						os << "H";
						break;
					case TILE_BASE:
						os << "X";
						break;
				}
			}
			os << endl;
		}
		return os;
	}
};

struct World{
	vector<Harbor> harbors;
	//vector<Ship> ships;
	Map mapa;
	World(){cerr<<"New world"<<endl;};
	int gold,index;
	friend istream& operator>>(istream& is,World& world){
		string inp;is>>inp;
		RSJresource json(inp);
		world.harbors = load_harbors(json["harbors"]);
		cerr<<"loaded something"<<endl;
		if(false && json["map"]["tiles"].as_str() != "null")
			world.mapa = Map(json["map"]);
		else{
			cerr<<"Mapa je null"<<endl;
		}
		world.gold = json["gold"].as<int>();
		world.index = json["index"].as<int>();
		return is;
	}
};
};
#endif
