# ksp-proboj-2023-jesen

## Running info
  - maps folder musi byt o jedno vyssie ako je folder repa
  - na observer: pred buildom treba doinstalovat package `npm install -f package.json` a potom `npm run build` a potom `npm run dev`
  - build a move servera: `go build . && mv server ../srv`
  - runner: `./runner config.json games.json `


## Pravidla
Kazdy hrac ovlada flotilu lodi. Lode sa pohybuju po mape, kde je more(volno) a ostrovy(stena). Na ostrovoch su pristavy -- specialne more, kde moze byt viacero lodi naraz. Hrac ma zakladnu. Na zaciatku ma hrac zlatky a nejake resources (na spravenie lode). Kazdy pristav ma surovinu, ktoru vytvara a surovinu, ktoru spotrebuva a kapacitu na obe a rate vytvarania a spotreby. Existuje funkcia, ktora urci cenu suroviny podla naplnenosti skladu. 
# Objects

## Harbors
Kazdy ostrov ma pristav. 
- V pristave sa nemoze strielat. 
- Ked sa attack lode priblizia k pristavu dostavaju damage. 
- poskodene lode sa mozu v pristave healovat za zlatky

## Base
Ked do nejak donesies zlatky uchovaju sa na do tvojeho "trezoru" - inventar.
Mozes tu kupovat a vylepsovat nove lode tiez za zlatky.

## Ships
Existuje viacero typov lodi.
| |Cln| Plt| Lod 1| Lod 2| Lod 3| ...|..|..|..|
|----|----|----|----|----|----|----|--|--|--|
|HP| 10|25|15|50|10|50|15|5|
|DMG| 1|1|1|1|3|5|10|1|
|RANGE|1|1|1|1|1|1|3|1|
|MOVE CD|4|2|1|4|2|4|4|1|
|CARGO|5|20|10|50|10|30|5|20|
|PRICE|10|50|60|100|10|50|50|30|
|YIELD|20%|20%|20%|20%|50%|50%|50%|80%|
|TYPE|trade|trade|trade|trade|attack|attack|attack|loot|

Lode sa daju v baseke vylepsovat - statny sa prenasobia nejakym cislom.

## Resources
|Name| Base price|Prod rate|
|----|-----------|---------|
|Wood|          1|       10|
|Stone|         2|        5|
|Iron|          5|        2|
|Fancy rock|   10|        1|
|Wool|          3|        3|
|Hide|          5|        2|
|Wheat|         2|        5|
|Pineapple|         3|        3|
|----|-----------|---------|
|Gold|          1|        0|

# Mechanics
## Movement
Kazda lod sa moze hybat o jedno policko v casovom intervale podla `MOVE CD`.
Na kazdom policku mapy (okrem pristavu) sa moze nachadzat len jedna lod. 
## Combat
Ked je ina lod v dosahu tvojej `attack` lode mozes na nu v danom kole vystrelit. Po tom, ako je lod znicena zostane na policku vrak, ktory obsahuje jej cargo.
## Looting
Z vraku potopenej lode mozes vyextrahovat nejaky ten zvysny loot. Tu je zaujimavy `YIELD` atribut tvojej lode, ktory urcuje kolko najviac je schopna z vraku vytiahnut. Specialna class lodi - `loot`
## Trading
Pri kazdej zastavke v pristave mozes tradovat s nim.
Ked si na mori mozes vymienat tovary medzi vlastymi lodami ak si 1 od seba. 
V base mozes ulozit svoje hard earned zlataky.

# Goal
To be discussed...
## Metrics
### Gold earned
It's a trading game - more is more. Total gold you've earned troughout the game count.
### Gold left at bank
How much gold did you manage to safe by the end of the game.
### Number of kills
Are you a manace at heart ? You also get rewarded by the amout of ships you've managed to sunk.
### Trades made
Does commerce consume your mind easyily, but you aren't very good with the numbers? Don't worry! We also reward you for trying..
## Score formula
TBA
