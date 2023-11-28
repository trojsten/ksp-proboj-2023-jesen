# ksp-proboj-2023-jesen

# Pravidlá

## Krátky opis hry

Každý hráč riadi folitlu lodí. Cieľom hry je sa stať pomocou obchovania a zabíjania
ostatných hráčov tým najbohatším z nich.

## Ako vyzerá mapa ?

Mapa je rozdelená na dve časti - more a zem. Po mori sa, samozrejme, budete pohybovať loďami.
Na súši sa nachádzajú základne hráčov a prístavy.

Mapa je mriežka.

## Herné mechaniky

Pre každú loď je možné vykonať jeden z týchto úkonov.

### Pohyb

Každá loď sa može pohybovať po mriežke v štyroch smeroch. V jednom ťahu sa loď može pohnúť o najviac toľko
políčok, koľko je jej dosahu (`MaxMoveRange`).

### Obchodovanie

Keď loď dorazí prístavu može s ním obochodovať. Vieme nakupovať a predávať suroviny. Ceny týchto surovín sa menia
na základe produkcie a dopytu.

### Nákup lodí

Môžeme kúpiť nové lode a objavia sa v základni.

### Boj

Lode medzi sebou možu bojovať - strielať po sebe. Možeme strielať na lubovoľnú loď v dosahu (`Range`) lode.
Ak na loď vystrelím a je v dosahu vždy sa trafím a vykonám poškodenie (`Damage`). Ak loď zníčím zostane po nej vrak.

### Lootavanie

Zničené lode - vraky možem lootvať. Lootavanie je však dosť náročné a teda nedostaneme celý obsah nákladného priestoru.
Koľko z neho naozaj dostaneme záleží na state `Yield` lootujucej lode.

### Odkladanie bohatstva

Keďže lode môžu byť zničené, tak si zlato vieme odloziť aj do základne.

## Herné objekty

Všetky objekty, s ktorými sa počas hry interaguje.

### Základňa

Základňa je miesto, kde sa spawnujú naše nakúpené lode. V základi môžem uložiť svoje goldy.
Taktiež sa mi v základni "opravuje" poškodená loď. Základňa striela na všetky nepriateľské lode,
ktoré sa k nej priblížia.

### Prístav

V prístave vytvára suroviny a taktiež suroviny spotrebúva. Vieme tu suroviny nakupovať a predávať suroviny.
V prístave neviem útočiť na iné lode. Prístav je však mierová zóna teda na všetky útočné lode, ktoré
sa priblížia k prístavu strielať. Taktiež nie je možné strielať na žiadne lode, ktoré sú blízko prístavu.

### Loď

Loď sa pohybuje po mriežke. Každá loď má svoj nákladný priestor, do ktorého ukladá suroviny. Lode sú rôznych
typov a majú rôzne štastiky:

+ **MaxHealth**: počet životov, ktoré má loď.
+ **Damage**: počet životov, ktorá loď uberá pri strielaní.
+ **Range**: vzdialenosť, na ktorú loď dovidí.
+ **MaxMoveRange**: maximálna vzdialenosť, ktorú može loď prejsť v jednom ťahu.
+ **MaxCargo**: kapacita nákladného priestoru.
+ **Price**: koľko zlata za loď zaplatím.
+ **Yield**: aké percento surovín je pri lootovaní loď schopná extrahovať.
+ **Class**: akého typu je loď.

| Name               | MaxHealth | Damage | Range | MaxMoveRange | MaxCargo | Price | Yield | Class  |
|--------------------|-----------|--------|-------|--------------|----------|-------|-------|--------|
| Cln                | 10        | 1      | 1     | 2            | 10       | 10    | 20    | Trade  |
| Plt                | 15        | 1      | 2     | 1            | 50       | 30    | 20    | Trade  |
| SmallMerchantShip  | 30        | 1      | 3     | 3            | 50       | 100   | 20    | Trade  |
| LargeMerchantShip  | 50        | 2      | 4     | 2            | 100      | 200   | 20    | Trade  |
| SomalianPirateShip | 10        | 3      | 2     | 3            | 5        | 15    | 50    | Attack |
| BlackPearl         | 50        | 5      | 4     | 2            | 30       | 50    | 50    | Attack |
| SniperAttackShip   | 5         | 8      | 3     | 1            | 10       | 30    | 50    | Attack |
| LooterScooter      | 5         | 0      | 5     | 4            | 30       | 50    | 80    | Loot   |

### Suroviny

Suroviny, ktoré nájdeme v prístavoch majú rôznu vzácnosť. Ich cena sa však mení podľa dopytu a produkcie.
Mení sa podľa tohto vzorca: **VZOREC**

| Name      | Base price |
|-----------|------------|
| Wood      | 1          |
| Stone     | 2          |
| Wheat     | 2          |
| Pineapple | 3          |
| Wool      | 3          |
| Iron      | 5          |
| Hide      | 5          |
| Gem       | 10         |
|           |            |
| Gold      | 1          |

## Hodnotenie

Vaši boti budú hodnotení na základe týchto kritérií:

+ **Zarobené zlato**: počet zlata, získaného počas hry.
+ **Aktuálne zlato**: počet zlata, ktoré akutálne má.
+ **Počet zabití lodí**: počet zabitých lodí.
+ **Objem predajov**: počet surovín, ktoré predal.
+ **Objem nákupov**: počet surovín, ktoré kúpil.