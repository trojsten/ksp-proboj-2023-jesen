import Konva from 'konva';
import { Resources, Ship, ShipType } from './observer';
import Playback from './playback';
import Stats from './stats';


export default class ShipClass {
    ship: Konva.Group;
    mouseOver = false;
    constructor(private data: Ship, shipLayer: Konva.Layer, tileSize: number, type: ShipType) {
        const image = new Image();
        image.src = `./Ships/${type}.png`;
        const konvaImage = new Konva.Group({
            y: data.y * tileSize + tileSize / 2,
            x: data.x * tileSize + tileSize / 2,
            width: tileSize,
            height: tileSize,
            rotation: 0,
            opacity: 0,
            listening: false,
        });

        konvaImage.on('mouseenter', () => {
            this.mouseOver = true;
            konvaImage.find('#circle')[0].opacity(1);
        });
        konvaImage.on('mouseleave', () => {
            this.mouseOver = false;
            konvaImage.find('#circle')[0].opacity(0.2);
        });

        konvaImage.add(new Konva.Circle({
            radius: 0.8 * tileSize / 2,
            fill: 'black',
            opacity: 0.2,
            id: 'circle',
        }), new Konva.Image({
            image: image,
            x: -tileSize / 2,
            y: -tileSize / 2,
            width: tileSize,
            height: tileSize,
        }))

        konvaImage.on('click', () => {
            Stats.ShowShipStats(this.data);
        })
        new Konva.Animation(() => {
            if ((Stats.currentShown as Ship)?.index == this.data.index) {
                konvaImage.find('#circle')[0].opacity(1);
            }
            else if ((Stats.currentShown == null || (Stats.currentShown as Ship).index != this.data.index) && !this.mouseOver) {
                konvaImage.find('#circle')[0].opacity(0.2);
            }
        }).start();
        shipLayer.add(konvaImage);
        this.ship = konvaImage;
    }

    move(newX: number, newY: number) {
        if (this.data.is_wreck) return;
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
            rotation: Math.atan2(delta.y, delta.x) * 180 / Math.PI,
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
        this.ship.opacity(0.2);
    }

    setHealth(health: number) {
        this.data.health = health;
    }

    remove() {
        this.ship.remove();
    }

    setResources(resources: Resources) {
        this.data.resources = resources;
    }

    deselect() {    
        this.ship.find('#circle')[0].opacity(0.2);
    }
}