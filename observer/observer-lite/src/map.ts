import Konva from "konva";
import harbor from "./assets/Harbor.png";
import water from "./assets/Water.png";
import base from "./assets/Base.png";
import ground from "./assets/Ground.png";
import { GameMap } from "./observer";

export function createMap(mapLayer: Konva.Layer, map: GameMap) {
    const tileSize = 20;
    const images = [water, ground, harbor, base].map(createImage);
    console.log(map);
    
    for (let i = 0; i < map.width; i++) {
        for (let j = 0; j < map.height; j++) {
            const tile = map.tiles[i][j];
            const image = images[tile.type];
            const x = i * tileSize;
            const y = j * tileSize;
            const konvaImage = new Konva.Image({
                image: image,
                x: x,
                y: y,
                width: tileSize,
                height: tileSize
            });
            // konvaImage.cache();
            mapLayer.add(konvaImage);
        }
    }    
}

function createImage(src: string) {
    const image = new Image();
    image.src = src;
    return image;
}