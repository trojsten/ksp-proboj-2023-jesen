import Konva from 'konva';
import ShipImage from './assets/Ship.png';
import { Ship } from './observer';


export default class ShipClass {
    image: Konva.Image;
    constructor(private data: Ship, shipLayer: Konva.Layer, tileSize: number) {
        const image = new Image();
        image.src = ShipImage;
        const konvaImage = new Konva.Image({
            image: image,
            x: data.x * tileSize,
            y: data.y * tileSize,
            width: tileSize,
            height: tileSize
        });
        shipLayer.add(konvaImage);
        this.image = konvaImage;
    }

    move(newX: number, newY: number) {
        this.data.x = newX;
        this.data.y = newY;
        new Konva.Tween({
            node: this.image,
            duration: 0.1,
            x: newX * 20,
            y: newY * 20
        }).play();
    }

    setWreck() {
        this.data.is_wreck = true;
        this.image.opacity(0.5);
    }

    remove() {
        this.image.remove();
    }
}