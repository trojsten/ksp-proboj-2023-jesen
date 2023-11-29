#ifndef PROBOJ_H
#define PROBOJ_H
#include "common.h"
#include <bits/stdc++.h>

// hack, aby sa dalo zistit, ci je struktura iterovatelna
namespace detail {
// To allow ADL with custom begin/end
using std::begin;
using std::end;

template <typename T>
auto is_iterable_impl(int)
    -> decltype(begin(std::declval<T&>()) != end(std::declval<T&>()),   // begin/end and operator !=
                void(),                                                 // Handle evil operator ,
                ++std::declval<decltype(begin(std::declval<T&>()))&>(), // operator ++
                void(*begin(std::declval<T&>())),                       // operator*
                std::true_type{});

template <typename T> std::false_type is_iterable_impl(...);

} // namespace detail

template <typename T> using is_iterable = decltype(detail::is_iterable_impl<T>(0));

// pomocou tohto vieme vypisat lubovolnu iterovatelnu strukturu
template <typename T> std::ostream& operator<<(std::ostream& os, const std::vector<T>& v) {
    for (int i = 0; i < v.size(); i++) {
        os << v[i];
        if (i != v.size() - 1 && !is_iterable<T>())
            os << " ";
    }
    if (!is_iterable<T>())
        os << "\n";
    return os;
}

template <typename L, typename R>
std::ostream& operator<<(std::ostream& os, const std::pair<L, R>& p) {
    os << "(" << p.first << "," << p.second << ")";
    return os;
}

// hashovanie pairov a XY
namespace std {
template <typename L, typename R> struct hash<pair<L, R>> {
    auto operator()(const pair<L, R>& m) { return 231095 * m.first + m.second; }
};

template <> struct hash<XY> {
    auto operator()(const XY& m) const { return hash<pair<int, int>>()({m.x, m.y}); }
};
} // namespace std

std::vector<XY> SMERY{{0, 1}, {0, -1}, {1, 0}, {-1, 0}};
std::vector<XY> ADJ{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}, {0, 1}, {0, -1}, {1, 0}, {-1, 0}};

/// @brief Zráta vzdialenosti od viacero bodov
/// @param start vector štartovacích bodov
/// @param condition funkcia, ktorá vráti true ak sa dá pohnúť z bodu a do bodu b
/// @param dist mapa vzdialeností a rodičov, ktorú táto funkcia naplní
/// @param transitions povolené smery pohybu
void bfs(std::vector<XY>& start, bool (*condition)(XY, XY),
         std::unordered_map<XY, std::pair<int, XY>>& dist, std::vector<XY>& transitions = SMERY) {
    std::queue<XY> q;
    for (auto i : start) {
        q.push(i);
        if (dist.find(i) == dist.end()) {
            dist[i] = {0, i};
        } else {
            dist[i].second = i;
        }
    }
    while (!q.empty()) {
        XY nv = q.front();
        q.pop();
        for (auto i : transitions) {
            if (condition(nv, nv + i) && dist.find(nv + i) == dist.end()) {
                q.push(nv + i);
                dist[nv + i] = {dist[nv].first + 1, nv};
            }
        }
    }
}

/// @brief Zráta vzdialenosti od jedného bodu
/// @param start štartovací bod
/// @param condition funkcia, ktorá vráti true ak sa dá pohnúť z bodu A do bodu B
/// @param dist mapa vzdialeností a rodičov, ktorú táto funkcia naplní
/// @param transitions povolené smery pohybu (default susedné hranou)
void bfs(XY start, bool (*condition)(XY, XY), std::unordered_map<XY, std::pair<int, XY>>& dist,
         std::vector<XY>& transitions = SMERY) {
    std::vector<XY> tmp{start};
    bfs(tmp, condition, dist, transitions);
}

/// @brief Vypočíta cestu z destinácie a mapy vzdialeností
/// @param destination koncový bod
/// @param dist mapa vzdialeností
/// @return vector bodov na ceste (od start do end)
std::vector<XY> recreate_path(XY destination, std::unordered_map<XY, std::pair<int, XY>>& dist) {
    std::vector<XY> out;
    XY cur = destination;
    if (dist.find(destination) == dist.end())
        return {};
    while (dist[cur].second != cur) {
        out.push_back(cur);
        cur = dist[cur].second;
    }
    out.push_back(cur);
    std::reverse(out.begin(), out.end());
    return out;
}

/// @brief Posuň sa smerom na bod end
/// @param ship loď, ktorou pohybujem
/// @param destination cieľ
/// @param condition funkcia, ktorá vráti true ak sa dá pohnúť z bodu a do bodu b
/// @param transitions povolené smery pohybu (default susedné hranou)
/// @return bod, kam sa mám posunúť tento ťah
XY move_to(Ship& ship, XY destination, bool (*condition)(XY, XY),
           std::vector<XY>& transitions = SMERY) {
    XY start = ship.coords;
    int range = ship.stats.max_move_range;
    std::unordered_map<XY, std::pair<int, XY>> dist;
    bfs(start, condition, dist, transitions);
    if (dist.find(destination) == dist.end())
        return ship.coords;
    std::vector<XY> path = recreate_path(destination, dist);
    return path[std::min((int)path.size() - 1, range)];
}
#endif
