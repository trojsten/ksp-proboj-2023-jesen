#ifndef PROBOJ_H
#define PROBOJ_H
#include<bits/stdc++.h>
#define POINT std::pair<int,int>

POINT operator+(POINT &a,POINT &b){
    return {a.first + b.first,a.second + b.second};
}

namespace detail
{
    // To allow ADL with custom begin/end
    using std::begin;
    using std::end;

    template <typename T>
    auto is_iterable_impl(int)
    -> decltype (
        begin(std::declval<T&>()) != end(std::declval<T&>()), // begin/end and operator !=
        void(), // Handle evil operator ,
        ++std::declval<decltype(begin(std::declval<T&>()))&>(), // operator ++
        void(*begin(std::declval<T&>())), // operator*
        std::true_type{});

    template <typename T>
    std::false_type is_iterable_impl(...);

}

template <typename T>
using is_iterable = decltype(detail::is_iterable_impl<T>(0));

template<typename T>
std::ostream& operator<< (std::ostream &os,const std::vector<T>&v){
    for(int i = 0;i < v.size();i++){
        os<< v[i];
        if(i != v.size()-1 && !is_iterable<T>()) os<<" ";
    }
    if(!is_iterable<T>()) os<<"\n";
    return os;
}

std::vector<POINT> SMERY{{0,1},{0,-1},{1,0},{-1,0}};
std::vector<POINT> ADJ{{1,1},{-1,1},{1,-1},{-1,-1},{0,1},{0,-1},{1,0},{-1,0}};

/// @brief Zráta vzdialenosti od viacero bodov
/// @param start vector štartovacích bodov
/// @param condition funkcia, ktorá vráti true ak sa dá pohnúť z bodu a do bodu b
/// @param dist mapa vzdialeností a rodičov, ktorú táto funkcia naplní
/// @param transitions povolené smery pohybu
void bfs(std::vector<POINT> &start, bool (*condition)(POINT,POINT), std::map<POINT,std::pair<int,POINT>> &dist,std::vector<POINT> &transitions = SMERY){
    std::queue<POINT> q;
    for(auto i : start){
        q.push(i);
        if(dist.find(i) == dist.end()){
            dist[i] = {0,i};
        }else{
            dist[i].second = i;
        }
    }
    while(!q.empty()){
        POINT nv = q.front();q.pop();
        for(auto i : transitions){
            if(condition(nv,nv + i) && dist.find(nv + i) == dist.end()){
                q.push(nv+i);
                dist[nv+i] = {dist[nv].first + 1,nv};
            }
        }
    }
}

/// @brief Zráta vzdialenosti od jedného bodu
/// @param start štartovací bod
/// @param condition funkcia, ktorá vráti true ak sa dá pohnúť z bodu a do bodu b
/// @param dist mapa vzdialeností a rodičov, ktorú táto funkcia naplní
/// @param transitions povolené smery pohybu
void bfs(POINT start, bool (*condition)(POINT,POINT), std::map<POINT,std::pair<int,POINT>> &dist,std::vector<POINT> &transitions = SMERY){
    std::vector<POINT> tmp{start};
    bfs(tmp,condition,dist,transitions);
}

void dijkstra(std::vector<POINT> &start,int (*cost)(POINT,POINT),std::map<POINT,std::pair<int,POINT>> &dist,std::vector<POINT> &transitions = SMERY){
    std::priority_queue<std::pair<int,POINT>> q;
    for(auto i : start){
        if(dist.find(i) == dist.end()){
            dist[i] = {0,i};
            q.push({0,i});
        }else{
            dist[i].second = i;
            q.push({-dist[i].first,i});
        }
    }
    while(!q.empty()){
        std::pair<int,POINT> nv = q.top();q.pop();
        if(dist.find(nv.second) != dist.end()) continue;
        for(auto i : transitions){
            if(cost(nv.second,nv.second + i) != INT_MAX && dist.find(nv.second + i) == dist.end()){
                q.push({nv.first - cost(nv.second,nv.second + i),nv.second + i});
                dist[nv.second+i] = {-nv.first + cost(nv.second,nv.second + i),nv.second};
            }
        }
    }
}

void dijkstra(POINT start,int (*cost)(POINT,POINT),std::map<POINT,std::pair<int,POINT>> &dist,std::vector<POINT> &transitions = SMERY){
    std::vector<POINT> tmp{start};
    dijkstra(tmp,cost,dist,transitions);
}

/// @brief Zráta vzdialenosti medzi dvoma bodmi
/// @param a prvý bod
/// @param b druhý bod
/// @return vzdialenosť medzi a a b
int dist(POINT a,POINT b){
    return std::abs(a.first - b.first) + std::abs(a.second - b.second);
}

/// @brief Vypočíta cestu z destinácie a mapy vzdialeností
/// @param end koncový bod
/// @param dist mapa vzdialeností
/// @return vector bodov na ceste
std::vector<POINT> recreate_path(POINT end,std::map<POINT,std::pair<int,POINT>> &dist){
    std::vector<POINT> out;
    POINT cur = end;
    while(dist[cur].second != cur){
        out.push_back(cur);
        cur = dist[cur].second;
    }
    out.push_back(cur);
    std::reverse(out.begin(),out.end());
    return out;
}


/// @brief Posuň sa smerom na bod end
/// @param start začiatočný bod
/// @param end destinácia
/// @param condition funkcia, ktorá vráti true ak sa dá pohnúť z bodu a do bodu b
/// @param range range lode, alebo o akú vzdialenosť sa viem posunúť naraz
/// @param transitions povolené smery pohybu
/// @return bod, kam sa mám posunúť tento ťah
POINT move_to(POINT start,POINT end, bool (*condition)(POINT,POINT),int range = 1,std::vector<POINT> &transitions = SMERY){
    std::map<POINT,std::pair<int,POINT>> dist;
    bfs(start,condition,dist,transitions);
    std::vector<POINT> path = recreate_path(end,dist);
	return path[std::min((int)path.size()-1,range)];
}
#endif
