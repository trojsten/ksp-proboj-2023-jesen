import Konva from 'konva';
import ShipImage from './assets/Ship.png';

import { Ship } from './observer';
import Playback from './playback';
import Stats from './stats';


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
            opacity: 0,
            listening: false,
        });

        konvaImage.add(new Konva.Image({
            image: image,
            x: -tileSize / 2,
            y: -tileSize / 2,
            width: tileSize,
            height: tileSize,
        }))

        konvaImage.on('click', () => {
            Stats.ShowShipStats(this.data);
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

        const newTile = Playback.map.tiles[newY][newX];
        if (newTile.type == 2 || newTile.type == 3) {
            tween._addAttr('opacity', 0);
            this.ship.listening(false);
        } else {
            tween._addAttr('opacity', 1);
            this.ship.listening(true);
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


}