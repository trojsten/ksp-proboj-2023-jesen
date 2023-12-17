# ksp-proboj-2023-jesen
# Čo je to Proboj a ako funguje? 
Proboj, skratka pre progamátorský boj, je aktivita z KSP sústredení, kde hráči (vy) programujú vlastného bota, ktorý
súťazí v predom pripravenej hre. K hre je taktiež pripravený template bota, ktorý zvláda komunikáciu so serverom a nejaké
užitočné funkcie. Taktiež obsahuje veľmi jednoduchý príklad jednoduchého bota, ktorého môžete dalej upravovať.

# Pravidlá hry

## Krátky opis hry

Každý hráč riadi folitlu lodí. Cieľom hry je sa stať pomocou obchovania a zabíjania
ostatných hráčov tým najbohatším z nich.

Hra sa hrá na ťahy - pre každú loď môžete v jednom ťahu urobiť jeden úkon, ktorý pridáte do zoznamu úkonov, ktoré tvoria váš ťah.

## Ako vyzerá mapa ?

Mapa je rozdelená na dve časti - more a zem. Po mori sa, samozrejme, budete pohybovať loďami.
Na súši sa nachádzajú základne hráčov a prístavy.

Mapa je mriežka, po ktorej sa viete pohybovať v štyroch smeroch. 

## Herné mechaniky

Pre každú loď je možné vykonať jeden z týchto úkonov - pohyb, obchodovanie, strielanie, lootovanie. Okrem toho môžeme nakupovať nové lode.

### Pohyb

Každá loď sa može pohybovať po mriežke v štyroch smeroch. 

V jednom ťahu sa loď može pohnúť o najviac toľko
políčok, koľko je v jej dosahu (`MaxMoveRange`) - dosah sa ráta ako BFS vzdialenosť.

Loďou sa nemôžete pohnúť na políčko, kde už je iné loď (výnimkou je prístav). 

Možete však vašou loďou "preskakovať" ostatné lode alebo aj zem, všetko v rámci dosahu lode (`MaxMoveRange`) 

### Obchodovanie

Keď loď dorazí do prístavu, može s ním obchodovať. Vie nakupovať a predávať suroviny. 

Ceny týchto surovín sa menia na základe produkcie a dopytu. 

Prístav však nechce suroviny, ktoré mu netreba - dajú sa mu predať len tie, ktoré konzumuje. 

Keď nakupujem alebo predávam tak sa na celú transakciu aplikuje cena vzhľadom na stav sladku na jej začiatku.

### Boj

Lode medzi sebou možu bojovať - strielať po sebe. Možeme strielať na lubovoľnú loď v dosahu (`Range`) lode.
Ak na loď vystrelím a je v dosahu vždy sa trafím a vykonám poškodenie (`Damage`). Ak loď zníčím, zostane po nej vrak.

### Lootavanie

Zničené lode - vraky - možem lootvať. Lootavanie je však dosť náročné a teda nedostaneme celý obsah nákladného priestoru.

Koľko z neho naozaj dostaneme záleží na state `Yield%` lootujucej lode. Z každej suroviny na lodi dostanem do nákladného
priestoru lode, ktorá lootuje `počet suroviny*Yield%` pre každú surovinu, ktorú mala zničená loď na palube.

### Odkladanie bohatstva

Keďže lode môžu byť zničené, tak si zlato vieme odložiť aj do základne.

Zo základne si vieme uložené zlato znova vybrať.

### Nákup lodí

Môžeme kúpiť nové lode. Objavia sa v základni hráča s prázdnym nákladnym priestorom. 

Novo nakúpenej lodi môžem dávať príkazy až v nasledujúcom kole od nákupu.

## Herné objekty

Všetky objekty, s ktorými sa počas hry interaguje.

### Základňa

Základňa je miesto, kde sa spawnujú naše nakúpené lode. 

V základi môžem uložiť svoje goldy.

Taktiež sa mi v základni "opravuje" poškodená loď. Pri pobyte v základni sa lodi regeneruje 1 HP každé kolo.

Základňa striela na všetky nepriateľské lode, ktoré sa k nej priblížia. Toto robí vo vzdialenosti štyroch políčok od seba.
Základňa týmto lodiam dáva 1 damage.

### Prístav

Prístav vytvára suroviny a taktiež suroviny spotrebúva. Na to aby sme zistili aké suroviny predáva a aký je stav jeho skladu musíme byť priamo v prístave.

Vieme tu suroviny nakupovať a predávať.

V prístave nevieme útočiť na iné lode. 

Prístav je však mierová zóna teda na všetky útočné lode, ktoré
sa priblížia k prístavu bude prístav strielať. Toto robí vo vzdialenosti ôsmich políčok od seba.
Taktiež nie je možné strielať na žiadne lode, ktoré sú blízko prístavu. Platí to pre lode, ktoré sú bližšie ako štyry políčka k prístavu. Prístav týmto lodiam dáva 1 damage.

### Loď

Loď sa pohybuje po mriežke. Každá loď má svoj nákladný priestor, do ktorého ukladá suroviny. Lode sú rôznych
typov a majú rôzne staty:

+ **MaxHealth**: počet životov, ktoré má loď.
+ **Damage**: počet životov, ktorá loď uberá pri strielaní.
+ **Range**: vzdialenosť, na ktorú loď dovidí.
+ **MaxMoveRange**: maximálna vzdialenosť, ktorú može loď prejsť v jednom ťahu.
+ **MaxCargo**: kapacita nákladného priestoru.
+ **Price**: koľko zlata za loď zaplatím.
+ **Yield%**: aké percento surovín je pri lootovaní loď schopná extrahovať.
+ **Class**: akého typu je loď.

| Name               | MaxHealth | Damage | Range | MaxMoveRange | MaxCargo | Price | Yield% | Class  |
|--------------------|-----------|--------|-------|--------------|----------|-------|-------|--------|
| Cln                | 10        | 1      | 1     | 2            | 10       | 10    | 20 %  | Trade  |
| Plt                | 15        | 1      | 2     | 1            | 50       | 30    | 20 %  | Trade  |
| SmallMerchantShip  | 30        | 1      | 3     | 3            | 50       | 100   | 20 %  | Trade  |
| LargeMerchantShip  | 50        | 2      | 4     | 2            | 100      | 200   | 20 %  | Trade  |
| SomalianPirateShip | 10        | 3      | 2     | 3            | 5        | 15    | 50 %  | Attack |
| BlackPearl         | 50        | 5      | 4     | 2            | 30       | 50    | 50 %  | Attack |
| SniperAttackShip   | 5         | 8      | 3     | 1            | 10       | 30    | 50 %  | Attack |
| LooterScooter      | 5         | 0      | 5     | 4            | 30       | 50    | 80 %  | Loot   |

### Suroviny

Suroviny, ktoré nájdeme v prístavoch majú rôznu vzácnosť. Ich cena sa však mení podľa dopytu a produkcie.
Mení sa podľa tohto vzorca: **cena = min(100/(amount+3)+1, 4) * BASE_PRICE[resourceType]**

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
+ **Aktuálne zlato**: počet zlata, ktoré akutálne má v základni a lodiach.
+ **Počet zabití lodí**: počet zabitých lodí.
+ **Objem predajov**: počet surovín, ktoré predal.
+ **Objem nákupov**: počet surovín, ktoré kúpil.

### Vzorec hodnotenia
**skóre = Zarobené zlato + Aktuálne zlato / 3 + Počet zabití lodí * 500 + Objem predajov / 5 + Objem nákupov / 5**
