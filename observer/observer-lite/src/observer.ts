import pako from 'pako';

export interface Turn {
    players: Player[];
    ships: Record<number, Ship>;
    bases: Base[];
    harbors: Harbor[];
    map: GameMap;
}

export interface Player {
    index: number;
    name: string;
    gold: number;
}

export interface Ship {
    index: number;
    player_index: number;
    x: number;
    y: number;
    resources: Resources;
    is_wreck: boolean;
}

export interface Base {
    player: number;
    x: number;
    y: number;
}

export interface Harbor {
    production: Resources;
    storage: Resources;
    x: number;
    y: number;
}

export interface GameMap {
    tiles: Maptile[][];
    width: number;
    height: number;
}

export interface Maptile {
    type: 0 | 1 | 2 | 3;
    index: number;
}

export interface Resources {
    wood: number;
    stone: number;
    iron: number;
    gem: number;
    wool: number;
    hide: number;
    wheat: number;
    pineapple: number;
    gold: number;
}

export async function loadGame(file: File | Blob): Promise<Turn[]> {
    const unzipped = pako.inflate(await file.arrayBuffer());
    const json = new TextDecoder().decode(unzipped);
    
    const turns = json.split('\n');
    turns.splice(turns.length - 1, 1);
    return turns.map(turn => JSON.parse(turn));
}