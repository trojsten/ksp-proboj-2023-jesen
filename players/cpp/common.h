#ifndef COMMON_H
#define COMMON_H
#include<bits/stdc++.h>
#include "json.hpp"

using namespace std;
using json = nlohmann::json;
enum class TileEnum : int {TILE_WATER, TILE_GROUND, TILE_HARBOR, TILE_BASE};
enum class ResourceEnum : int {Wood, Stone, Iron, Gem, Wool, Hide, Wheat, Pineapple, Gold};
enum class ShipsEnum : int {Cln,Plt,SmallMerchantShip,LargeMerchantShip,SomalianPirateShip,BlackPearl,SniperAttackShip,LooterScooter};
enum class ShipClass : int {SHIP_TRADE = 0,SHIP_ATTACK = 1,SHIP_LOOT = 1};
enum class TurnType : int {MOVE, TRADE, LOOT, SHOOT, BUY, STORE};

unordered_map<string,int> strToResource{
	{"wood" , 0},
	{"stone", 1},
	{"iron", 2},
	{"gem", 3},
	{"wool", 4},
	{"hide", 5},
	{"wheat", 6},
	{"pineapple", 7},
	{"gold", 8}
};

ostream& operator<<(ostream &os,const TurnType &t){
	vector<string> mapping{
		"MOVE",
		"TRADE",
		"LOOT",
		"SHOOT",
		"BUY",
		"STORE"
	};
	os << mapping[static_cast<int>(t)];
	return os;
}

struct XY{
	int x,y;
	XY(int _x,int _y) : x(_x),y(_y){}
	XY(){}
	XY(json &j){
		j.at("x").get_to(x);
		j.at("y").get_to(y);
	}
	XY operator-(XY other){return XY(this->x - other.x,this->y - other.y);}
	XY operator+(XY other){return XY(this->x + other.x,this->y + other.y);}
	friend ostream& operator<<(ostream &os,const XY &a){os << a.x << " " << a.y; return os;};
};

struct Turn{
	TurnType type;
	int ship_id;
	XY coords;
	int target,amount;
	int resource;
	int ship_to_buy;
	friend ostream& operator<<(ostream &os,const Turn &t){;
		os << t.type << " ";
		switch(t.type){
			case TurnType::MOVE:
				os << t.ship_id << " " << t.coords;
				break;
			case TurnType::TRADE:
				os << t.ship_id << " " << t.resource;
				break;
			case TurnType::LOOT:
			case TurnType::SHOOT:
				os << t.ship_id << " " << t.target;
				break;
			case TurnType::BUY:
				os << t.ship_to_buy;
				break;
			case TurnType::STORE:
				os << t.ship_id << " " << t.amount;
				break;
		}
		return os;
	}
};

struct MoveTurn : Turn{
	MoveTurn(int ship_id,XY coords) : Turn{TurnType::MOVE,ship_id,coords}{
	}
};

struct TradeTurn : Turn{
	TradeTurn(int ship_id,ResourceEnum resource,int amount) : Turn{TurnType::TRADE,ship_id}{
		this->resource = static_cast<int>(resource);
		this->amount = amount;
	}
};

struct LootTurn : Turn{
	LootTurn(int ship_id,int target) : Turn{TurnType::LOOT,ship_id}{
		this->target = target;
	}
};

struct ShootTurn : Turn{
	ShootTurn(int ship_id,int target) : Turn{TurnType::SHOOT,ship_id}{
		this->target = target;
	}
};

struct BuyTurn : Turn{
	BuyTurn(ShipsEnum ship_to_buy) : Turn{TurnType::BUY}{
		this->ship_to_buy = static_cast<int>(ship_to_buy);
	}
};

struct StoreTurn : Turn{
	StoreTurn(int ship_id,int amount) : Turn{TurnType::STORE,ship_id}{
		this->amount = amount;
	}
};

struct Resources{
	vector<int> resources = vector<int>(strToResource.size(),0);
	Resources(){}
	Resources(json j){
		for(auto &[key,value] : j.get<unordered_map<string,int>>()){
			resources[strToResource[key]] = value;
		}
	}
	int& operator[](TileEnum key){
		return resources[static_cast<int>(key)];
	}
};

struct Harbor{
	XY coords;
	Resources production, storage;
	bool visible;
	Harbor(json harbor_dict){
		coords = XY(harbor_dict["x"].get<int>(),harbor_dict["y"].get<int>());
		production = Resources(harbor_dict["production"]);
		storage = Resources(harbor_dict["storage"]);
		visible = harbor_dict["visible"].get<bool>();
	}
};

template<class T>
vector<T> load(const json &j){
	vector<T> ret;
	for(auto item : j.items()){
		ret.push_back(T(item.value()));
	}
	return ret;
}

struct Tile{
	TileEnum type;
	int index;
	Tile(int type, int index) : type(static_cast<TileEnum>(type)),index(index){}
	Tile(){}
	friend ostream& operator<<(ostream &os,const Tile &a){os << "Tile(" << static_cast<int>(a.type) << "," << a.index << ")";return os;}
};

void from_json(const json& j,Tile &t){
	t = Tile(j["type"],j["index"]);
}

struct Map{
	int width,height;
	vector<vector<Tile>> tiles;
	Map(){}
	Map(json r){
		width = r["width"].get<int>();
		height = r["height"].get<int>();
		tiles = r["tiles"].get<vector<vector<Tile>>>();
	}

	bool inside(XY coords){
		return coords.x >= 0 && coords.x < width && coords.y >= 0 && coords.y < height;
	}

	bool can_move(XY coords){
		return inside(coords) && tiles[coords.y][coords.x].type == TileEnum::TILE_WATER;
	}

	friend ostream& operator<<(ostream &os,const Map &a){
		os << "Map " << a.width << " " << a.height << endl;
		for(auto i : a.tiles){
			for(auto j : i){
				switch(j.type){
					case TileEnum::TILE_WATER:
						os << "~";
						break;
					case TileEnum::TILE_GROUND:
						os << "O";
						break;
					case TileEnum::TILE_HARBOR:
						os << "H";
						break;
					case TileEnum::TILE_BASE:
						os << "X";
						break;
				}
			}
			os << endl;
		}
		return os;
	}
};

struct ShipStats{
	int max_health,damage,range,max_move_range,max_cargo,price;
	float yield_frac;
	ShipClass ship_class;
	ShipStats(){}
	ShipStats(const json &j){
		j["max_health"].get_to(max_health);
		j["damage"].get_to(damage);
		j["range"].get_to(range);
		j["max_move_range"].get_to(max_move_range);
		j["max_cargo"].get_to(max_cargo);
		j["price"].get_to(price);
		j["yield_frac"].get_to(yield_frac);
		ship_class = static_cast<ShipClass>(j["ship_class"].get<int>());
	}
};

struct Ship{
	int index,player_index,health;
	XY coords;
	bool is_wreck,mine;
	Resources resources;
	ShipStats stats;
	Ship(){}
	Ship(const json &j){
		j["index"].get_to(index);
		j["player_index"].get_to(player_index);
		coords = XY(j["x"].get<int>(),j["y"].get<int>());
		j["health"].get_to(health);
		j["is_wreck"].get_to(is_wreck);
		resources = Resources(j["resources"]);
		stats = ShipStats(j["stats"]);
	}
};


struct World{
	vector<Harbor> harbors;
	vector<Ship> ships;
	Map mapa;
	World(){cerr<<"New world"<<endl;};
	int gold,index;
	friend istream& operator>>(istream& is,World& world){
		string inp;is>>inp;
		json data = json::parse(inp);
		world.harbors = load<Harbor>(data["harbors"]);
		world.ships = load<Ship>(data["ships"]);
		if(!data["map"]["tiles"].is_null())
			world.mapa = Map(data["map"]);
		else
			cerr<<"Mapa je null"<<endl;
		world.gold = data["gold"].get<int>();
		world.index = data["index"].get<int>();
		return is;
	}
};
#endif
