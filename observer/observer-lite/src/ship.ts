import Konva from 'konva';
import ShipImage from './assets/Ship.png';
import { Ship } from './observer';
import Playback from './playback';


export default class ShipClass {
    ship: Konva.Group;
    constructor(private data: Ship, shipLayer: Konva.Layer, tileSize: number) {
        const image = new Image();
        image.src = ShipImage;
        const konvaImage = new Konva.Group({
            y: data.y * tileSize + tileSize / 2,
            x: data.x * tileSize + tileSize / 2,
            width: tileSize,
            height: tileSize,
            rotation: 0,
            opacity: 0
        });

        konvaImage.add(new Konva.Image({
            image: image,
            x: -tileSize / 2,
            y: -tileSize / 2,
            width: tileSize,
            height: tileSize,
        }))
        
        konvaImage.on('click', () => {
            ShipClass.ShowShipStats(this.data);
        })
        shipLayer.add(konvaImage);
        this.ship = konvaImage;
    }

    move(newX: number, newY: number) {
        const delta = {
            x: newX - this.data.x,
            y: newY - this.data.y,
        }
        this.data.x = newX;
        this.data.y = newY;
        const tween = new Konva.Tween({
            node: this.ship,
            duration: 0.1,
            x: newX * 20 + 10,
            y: newY * 20 + 10,
        });

        new Konva.Tween({
            node: this.ship,
            rotation: Math.atan2(delta.y, delta.x) * 180 / Math.PI + 180,
            duration: 0.2,
        }).play();

        const newTile = Playback.turn.map.tiles[newY][newX];
        if(newTile.type == 2 || newTile.type == 3) {
            tween._addAttr('opacity', 0);
        } else {
            tween._addAttr('opacity', 1);
        }
        
        tween.play();
    }

    setWreck() {
        this.data.is_wreck = true;
        this.ship.opacity(0.5);
    }

    remove() {
        this.ship.remove();
    }

    static ShowShipStats(ship: Ship) {
        const stats = document.getElementById('shipStats')!;
        stats.innerHTML += `
            <div class="stats">
                <button class="close" id="closeBtn">X</button>
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
            stats.innerHTML = '';
        })
    }
}