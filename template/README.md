# Proboj Template

V tomto template sú všetky súbory, ktoré sú potrebné k rozbehaniu Proboja.

## Ako začať programovať?

Keď mám stiahnuté súbory templatu, otovorím tieto súbory v svojom obľúbenom editore a vyberiem si svoj obľúbenejší jazyk medzi c++ a python. Keď sa na tieto súbory pozriete, zistíte, že už majú zopár preprogramovaných funkcií, ktoré môžete používať.

Pre zvolený jazyk kóďte v pridelenom priečinku napr. pre python: `/player/python/`

## Ako si spustiť hru?
V template sa nachádzajú aj dva binárne súbory: server a runner a dva konfiguračné .json súbory: config a games.
Do config.json vyplním podľa templatu cesty k svojim hráčom a ich jazyk.
Potom v games.json vyberiem, ktorý z hráčov sa majú v hre objaviť a do args pridám cestu v k mape.

Následne už len spustím runner s configuračními súbormi. 

Pre linux: `./runner config.json games.json`.

Pre windows: `runner.exe config.json games.json`.

## Ako odovzdať bota?

Zazipujem súbory v mojom priečinku s hráčom.
Napr. pre `players/python/` to budú súbory: `player.py`, `proboj.py`, `ships.py`. Pre C++ používateľov - **NEZIPOVAŤ ZKOMPLIVANU BINARKU**

Potom na [stránku](http://proboj.ksp.sk/games/1/bots/) po svojho bota nahrám tento zip
a teším sa.
