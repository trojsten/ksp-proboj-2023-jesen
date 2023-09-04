# ksp-proboj-2023-jesen


## Pravidla
Kazdy hrac ovlada flotilu lodi. Lode sa pohybuju po mape, kde je more(volno) a ostrovy(stena). Na ostrovoch su pristavy -- specialne more, kde moze byt viacero lodi naraz. Hrac ma zakladnu. Na zaciatku ma hrac zlatky a nejake resources (na spravenie lode). Kazdy pristav ma surovinu, ktoru vytvara a surovinu, ktoru spotrebuva a kapacitu na obe a rate vytvarania a spotreby. Existuje funkcia, ktora urci cenu suroviny podla naplnenosti skladu. 

## Lode
Existuje viacero typov lodi.
| |Cln| Plt| Lod 1| Lod 2| Lod 3| ...|..|..|..|
|----|----|----|----|----|----|----|--|--|--|
|HP| 10|25|15|50|10|50|15|5|
|DMG| 1|1|1|1|3|5|10|1|
|RANGE|1|1|1|1|1|1|3|1|
|SPEED|1|2|5|1|2|1|1|5|
|STOR|5|20|10|50|10|30|5|20|
|PRICE|10|50|60|100|10|50|50|30|
|YIELD|20%|20%|20%|20%|50%|50%|50%|80%|
|TYPE|trade|trade|trade|trade|attack|attack|attack|loot|

Lode sa daju v baseke vylepsovat - statny sa prenasobia nejakym cislom.

## Pristavy
Kazdy ostrov ma pristav. 
- V pristave sa nemoze strielat. 
- Ked sa attack lode priblizia k pristavu dostavaju damage. 
- poskodene lode sa mozu v pristave healovat za zlatky

## Base
Ked do nejak donesies zlatky uchovaju sa na do tvojeho "trezoru" - inventar.
Mozes tu kupovat a vylepsovat nove lode tiez za zlatky.

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
|Apple|         3|        3|
|----|-----------|---------|
|Gold|          1|        0|
