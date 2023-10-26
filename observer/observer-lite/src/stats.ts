import { Base, Harbor, Ship, Turn } from "./observer";
import Playback from "./playback";


export default class Stats {
    static stats: HTMLDivElement;
    static currentShown: Ship | Harbor | Base | null = null;

    constructor() {
        Stats.stats = document.getElementById('stats') as HTMLDivElement;
    }

    static ShowShipStats(ship: Ship | null) {
        if (!ship) {
            Stats.stats.innerHTML = '';
            Stats.currentShown = null;
            return;
        }

        Stats.currentShown = ship;
        Stats.stats.innerHTML = `
            <div class="stats">
                <button class="close" id="closeBtn">
                    <span class="material-icons-round">close</span>
                </button>
                <h1>Ship ${ship.index} (${Playback.turn.players[ship.player_index].name})</h1>
                <div class="resources">
                    <p>Gem: ${ship.resources.gem}</p>
                    <p>Wood: ${ship.resources.wood}</p>
                    <p>Iron: ${ship.resources.iron}</p>
                    <p>Hide: ${ship.resources.hide}</p>
                    <p>Stone: ${ship.resources.stone}</p>
                    <p>Pineapple: ${ship.resources.pineapple}</p>
                    <p>Wool: ${ship.resources.wool}</p>
                    <p>Wheat: ${ship.resources.wheat}</p>
                    <p>Gold: ${ship.resources.gold}</p>
                </div>
            </div>
        `;
        document.getElementById('closeBtn')!.addEventListener('click', () => {
            Stats.stats.innerHTML = '';
        })
    }

    static showHarborStats(harbor: Harbor | null) {
        if (!harbor) {
            Stats.stats.innerHTML = '';
            Stats.currentShown = null;
            return;
        }
        Stats.currentShown = harbor;
        Stats.stats.innerHTML = `
            <div class="stats">
                <button class="close" id="closeBtn">
                    <span class="material-icons-round">close</span>
                </button>
                <h1>Harbor</h1>
                <div class="resources">
                    ${Stats.showHarborResource('gem')}
                    ${Stats.showHarborResource('wood')}
                    ${Stats.showHarborResource('iron')}
                    ${Stats.showHarborResource('hide')}
                    ${Stats.showHarborResource('stone')}
                    ${Stats.showHarborResource('pineapple')}
                    ${Stats.showHarborResource('wool')}
                    ${Stats.showHarborResource('wheat')}
                    ${Stats.showHarborResource('gold')}
                </div>
                <h2>Ships here</h2>
                <div class="ships">
                    ${Stats.shipsOnTile(harbor)}
                </div>
            </div>
        `;

        document.querySelectorAll('.ship').forEach((ship) => {
            ship.querySelector('button')!.addEventListener('click', (e) => {
                const shipIndex = parseInt((e.currentTarget as HTMLElement).id.replace('ship', ''));
                Stats.ShowShipStats(Playback.turn.ships[shipIndex]);
            })
        })

        document.getElementById('closeBtn')!.addEventListener('click', () => {
            Stats.stats.innerHTML = '';
        })
    }

    static showHarborResource(name: string) {
        const harbor = Stats.currentShown as Harbor;
        const displayName = name[0].toUpperCase() + name.slice(1);
        //@ts-ignore
        const production = harbor.production[name];
        //@ts-ignore
        const storage = harbor.storage[name];
        return `
            <p>${displayName}: <strong>${storage}</strong> <span style="color:${production >= 0 ? 'greenyellow' : 'tomato'}">${production}</span></p>
        `;
    }

    static shipsOnTile(tile: Harbor | Base) {
        const ships = Object.values(Playback.turn.ships).filter((ship) => {
            return ship.x === tile.x && ship.y === tile.y;
        });

        if (ships.length === 0) {
            return '<p style="color:tomato;">No ships here</p>';
        }

        return ships.map((ship) => {
            return `
                <div class="ship">
                    <p>Ship ${ship.index} (${Playback.turn.players[ship.player_index].name})</p>
                    <button id="ship${ship.index}">View</button>
                </div>
            `;
        }).join('');
    }

    Update(turn: Turn) {
        if (Stats.currentShown?.type === 'Ship') {
            Stats.ShowShipStats(turn.ships[Stats.currentShown.index]);
        }
        else if (Stats.currentShown?.type === 'Harbor') {
            const harbor = turn.harbors.find(h => h.x === Stats.currentShown!.x && h.y === Stats.currentShown!.y);
            Stats.showHarborStats(harbor!);
        }
    }

    static showBaseStats(base: Base | null) {
        if (!base) {
            Stats.stats.innerHTML = '';
            Stats.currentShown = null;
            return;
        }
        Stats.currentShown = base;
        const player = Playback.turn.players[base.player];
        Stats.stats.innerHTML = `
            <div class="stats">
                <button class="close" id="closeBtn">
                    <span class="material-icons-round">close</span>
                </button>
                <h1>${player.name}'s base</h1>
                <h2>Ships here</h2>
                <div class="ships">
                    ${Stats.shipsOnTile(base)}
                </div>
            </div>
        `;

        document.querySelectorAll('.ship').forEach((ship) => {
            ship.querySelector('button')!.addEventListener('click', (e) => {
                const shipIndex = parseInt((e.currentTarget as HTMLElement).id.replace('ship', ''));
                Stats.ShowShipStats(Playback.turn.ships[shipIndex]);
            })
        })

        document.getElementById('closeBtn')!.addEventListener('click', () => {
            Stats.stats.innerHTML = '';
        })
    }

}