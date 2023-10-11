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
        });

        konvaImage.add(new Konva.Image({
            image: image,
            x: -tileSize / 2,
            y: -tileSize / 2,
            width: tileSize,
            height: tileSize,
        }))
        konvaImage.cache();
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
        new Konva.Tween({
            node: this.ship,
            duration: 0.1,
            x: newX * 20 + 10,
            y: newY * 20 + 10,
            rotation: Math.atan2(delta.y, delta.x) * 180 / Math.PI + 180,
        }).play();
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
                <button class="close">X</button>
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
        stats.getElementsByClassName('close')[0].addEventListener('click', () => {
            stats.innerHTML = '';
        })
    }
}