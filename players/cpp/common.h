#ifndef COMMON_H
#define COMMON_H
#include "json.hpp"
#include <bits/stdc++.h>

namespace common {
using namespace std;
using json = nlohmann::json;
enum class TileEnum : int { TILE_WATER, TILE_GROUND, TILE_HARBOR, TILE_BASE };
enum class ResourceEnum : int { Wood, Stone, Iron, Gem, Wool, Hide, Wheat, Pineapple, Gold };
enum class ShipsEnum : int {
    Cln,
    Plt,
    SmallMerchantShip,
    LargeMerchantShip,
    SomalianPirateShip,
    BlackPearl,
    SniperAttackShip,
    LooterScooter
};
enum class ShipClass : int { SHIP_TRADE = 0, SHIP_ATTACK = 1, SHIP_LOOT = 1 };
enum class TurnType : int { MOVE, TRADE, LOOT, SHOOT, BUY, STORE };

unordered_map<string, int> strToResource{{"wood", 0},  {"stone", 1},     {"iron", 2},
                                         {"gem", 3},   {"wool", 4},      {"hide", 5},
                                         {"wheat", 6}, {"pineapple", 7}, {"gold", 8}};

ostream& operator<<(ostream& os, const TurnType& t) {
    vector<string> mapping{"MOVE", "TRADE", "LOOT", "SHOOT", "BUY", "STORE"};
    os << mapping[static_cast<int>(t)];
    return os;
}

/// @brief Class reprezentujúca pozíciu s užitočnými operáciami (sčitovanie, odčitovanie,
/// vypisovanie)
struct XY {
    int x, y;
    XY() {}
    XY(int _x, int _y) : x(_x), y(_y) {}
    XY(pair<int, int> xy) : x(xy.first), y(xy.second) {}
    XY(json& j) {
        j.at("x").get_to(x);
        j.at("y").get_to(y);
    }
    XY operator-(XY other) { return XY(this->x - other.x, this->y - other.y); }
    XY operator+(XY other) { return XY(this->x + other.x, this->y + other.y); }
    bool operator==(const XY other) const { return this->x == other.x && this->y == other.y; }
    bool operator!=(const XY other) const { return this->x != other.x || this->y != other.y; }
    friend bool operator<(const XY a, const XY b) { return a.x != b.x ? a.x < b.x : a.y < b.y; }
    friend ostream& operator<<(ostream& os, const XY& a) {
        os << a.x << " " << a.y;
        return os;
    };
    friend int dist(const XY a, const XY b) { return abs(a.x - b.x) + abs(a.y - b.y); }
};

struct Turn {
    TurnType type;
    int ship_id;
    XY coords;
    int target, amount;
    int resource;
    int ship_to_buy;
    friend ostream& operator<<(ostream& os, const Turn& t) {
        ;
        os << t.type << " ";
        switch (t.type) {
        case TurnType::MOVE:
            os << t.ship_id << " " << t.coords;
            break;
        case TurnType::TRADE:
            os << t.ship_id << " " << t.resource << " " << t.amount;
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

struct MoveTurn : Turn {
    /// @brief Posunie loď
    /// @param ship_id id lode, ktorú chceme posunúť
    /// @param coords pozícia, na ktorú chceme loď posunúť (musí byť v dosahu)
    MoveTurn(int ship_id, XY coords) : Turn{TurnType::MOVE, ship_id, coords} {}
};

struct TradeTurn : Turn {
    /// @brief Obchodovanie
    /// @param ship_id id lode, ktorá obchoduje (musí byť v prístave)
    /// @param resource typ suroviny na obchodovanie
    /// @param amount množstvo suroviny na obchodovanie. Ak je záporné, tak predávame, ak kladné,
    /// tak kupujeme
    TradeTurn(int ship_id, ResourceEnum resource, int amount) : Turn{TurnType::TRADE, ship_id} {
        this->resource = static_cast<int>(resource);
        this->amount = amount;
    }
};

struct LootTurn : Turn {
    /// @brief Lootime vrak
    /// @param ship_id id mojej lode, ktorá lootí
    /// @param target id vraku na lootenie
    LootTurn(int ship_id, int target) : Turn{TurnType::LOOT, ship_id} { this->target = target; }
};

struct ShootTurn : Turn {
    /// @brief Strieľame na loď
    /// @param ship_id id mojej lode, ktorá strieľa
    /// @param target id
    ShootTurn(int ship_id, int target) : Turn{TurnType::SHOOT, ship_id} { this->target = target; }
};

struct BuyTurn : Turn {
    /// @brief Kúpime loď
    /// @param ship_to_buy typ lode, ktorú chceme kúpiť
    BuyTurn(ShipsEnum ship_to_buy) : Turn{TurnType::BUY} {
        this->ship_to_buy = static_cast<int>(ship_to_buy);
    }
};

struct StoreTurn : Turn {
    /// @brief Uložíme zlato do základne (musíme byť vo svojej základni)
    /// @param ship_id id lode, z ktorej berieme zlato
    /// @param amount počet zlatiek, ktoré berieme, ak záporný, tak berieme zo základne, inak
    /// ukladáme do základne
    StoreTurn(int ship_id, int amount) : Turn{TurnType::STORE, ship_id} { this->amount = amount; }
};

struct Resources {
    vector<int> resources = vector<int>(strToResource.size(), 0);
    Resources() {}
    Resources(json j) {
        for (auto& [key, value] : j.get<unordered_map<string, int>>()) {
            resources[strToResource[key]] = value;
        }
    }
    int& operator[](ResourceEnum key) { return resources[static_cast<int>(key)]; }
    friend ostream& operator<<(ostream& os, const Resources& r) {
        vector<string> mapping{"wood", "stone", "iron",      "gem", "wool",
                               "hide", "wheat", "pineapple", "gold"};
        os << "Resources:\n";
        for (int i = 0; i < r.resources.size(); i++) {
            os << "\t" << mapping[i] << " : " << r.resources[i] << "\n";
        }
        return os;
    }
};

struct Harbor {
    XY coords;
    Resources production, storage;
    bool visible;
    Harbor(json harbor_dict) {
        coords = XY(harbor_dict["x"].get<int>(), harbor_dict["y"].get<int>());
        production = Resources(harbor_dict["production"]);
        storage = Resources(harbor_dict["storage"]);
        visible = harbor_dict["visible"].get<bool>();
    }
    /// @brief Za akú cenu kúpim/predám jednu jednotku suroviny?
    /// @param surovina
    /// @return cena jednej jednotky suroviny (ak na harbor nevidím, tak base cost)
    int resource_cost(ResourceEnum what) {
        vector<int> resourceCost{1, 2, 5, 10, 3, 5, 2, 3};
        if (!visible)
            return resourceCost[static_cast<int>(what)];
        int ret = min(100 / (storage[what] + 3) + 1, 4) * resourceCost[static_cast<int>(what)];
        return ret;
    }
};

template <class T> vector<T> load(const json& j) {
    vector<T> ret;
    for (auto item : j.items()) {
        ret.push_back(T(item.value()));
    }
    return ret;
}

struct Tile {
    TileEnum type;
    int index;
    Tile(int type, int index) : type(static_cast<TileEnum>(type)), index(index) {}
    Tile() {}
    friend ostream& operator<<(ostream& os, const Tile& a) {
        os << "Tile(" << static_cast<int>(a.type) << "," << a.index << ")";
        return os;
    }
};

void from_json(const json& j, Tile& t) { t = Tile(j["type"], j["index"]); }

struct Map {
    int width, height;
    vector<vector<Tile>> tiles;
    Map() {}
    Map(json r) {
        width = r["width"].get<int>();
        height = r["height"].get<int>();
        tiles = r["tiles"].get<vector<vector<Tile>>>();
    }

    /// @brief Sú súradnice vnútry hracej plochy?
    /// @param coords pozícia na overenie
    /// @return true ak je pozícia vnútry plochy
    bool inside(XY coords) {
        return coords.x >= 0 && coords.x < width && coords.y >= 0 && coords.y < height;
    }

    /// @brief Viem sa pohnút na pozíciu coords?
    /// @param coords pozícia na overenie
    /// @return true ak viem ísť na pozíciu coords
    bool can_move(XY coords) {
        return inside(coords) && tiles[coords.y][coords.x].type != TileEnum::TILE_GROUND;
    }

    Tile& operator[](XY coords) { return tiles[coords.y][coords.x]; }
    vector<Tile>& operator[](int y) { return tiles[y]; }
    friend ostream& operator<<(ostream& os, const Map& a) {
        os << "Map " << a.width << " " << a.height << endl;
        for (auto i : a.tiles) {
            for (auto j : i) {
                switch (j.type) {
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

struct ShipStats {
    int max_health, damage, range, max_move_range, max_cargo, price;
    float yield_frac;
    ShipClass ship_class;
    ShipStats() {}
    ShipStats(const json& j) {
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

struct Ship {
    int index, player_index, health;
    XY coords;
    bool is_wreck, mine;
    Resources resources;
    ShipStats stats;
    Ship() {}
    Ship(const json& j) {
        j["index"].get_to(index);
        j["player_index"].get_to(player_index);
        coords = XY(j["x"].get<int>(), j["y"].get<int>());
        j["health"].get_to(health);
        j["is_wreck"].get_to(is_wreck);
        j["mine"].get_to(mine);
        resources = Resources(j["resources"]);
        stats = ShipStats(j["stats"]);
    }

    /// @brief je loď v dosahu tejto lode?
    /// @param target pozícia na overenie
    /// @return true ak vieme z tejto lode strieľať na target
    bool can_attack(XY target) { return stats.range >= dist(coords, target); }

    /// @brief vidíme na pozíciu target?
    /// @param target pozícia na overenie
    /// @return true ak vidíme, čo je na pozícii target
    bool can_see(XY target) { return stats.range * 8 >= dist(coords, target); }
};

/// @brief Class reprezentujúca stav hry
struct World {
    vector<Harbor> harbors;
    vector<Ship> ships;
    /// @brief Class reprezentujúca mapu. Indexovať viete normálne mapa[y][x] alebo pomocou XY
    Map mapa;
    XY my_base = {-1, -1};
    int gold, index;
    World() { cerr << "New world" << endl; };

    /// @brief Získa všetky moje lode
    /// @return vector lodí, ktoré sú moje, teda ich viem ovládať
    vector<Ship> my_ships() {
        vector<Ship> out;
        for (Ship i : ships) {
            if (i.mine)
                out.push_back(i);
        }
        return out;
    }
    /// @brief Získa referenciu na loď s daným id
    /// @param id id hľdanej lode
    /// @param s loď, do ktorej sa uloží výsledok
    /// @retval `true` loď existuje
    /// @retval `false` loď neexistuje
    bool ship_by_id(int id, Ship& s) {
        for (Ship ship : ships) {
            if (ship.index == id) {
                s = ship;
                return true;
            }
        }
        return false;
    }
    friend istream& operator>>(istream& is, World& world) {
        string inp;
        is >> inp;
        json data = json::parse(inp);
        world.harbors = load<Harbor>(data["harbors"]);
        world.ships = load<Ship>(data["ships"]);
        world.index = data["index"].get<int>();
        if (!data["map"]["tiles"].is_null()) {
            world.mapa = Map(data["map"]);
            for (int i = 0; i < world.mapa.height; i++) {
                for (int j = 0; j < world.mapa.width; j++) {
                    Tile tile = world.mapa[j][i];
                    if (tile.type == TileEnum::TILE_BASE && tile.index == world.index)
                        world.my_base = {i, j};
                }
            }
        }
        world.gold = data["gold"].get<int>();
        return is;
    }
};
} // namespace common
using namespace common;
#endif
