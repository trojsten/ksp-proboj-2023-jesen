import Konva from "konva";
import harbor from "./assets/Harbor.png";
import water from "./assets/Water.png";
import base from "./assets/Base.png";
import ground from "./assets/Ground.png";
import { GameMap } from "./observer";

export function createMap(mapLayer: Konva.Layer, map: GameMap) {
    const tileSize = 20;
    const images = [water, ground, harbor, base].map(createImage);
    
    const grp = new Konva.Group({
    });

    const konvaImages = images.map((image) => {
        const img = new Konva.Image({
            image: image,
            width: tileSize,
            height: tileSize
        });
        img.cache();
        return img;
    });

    for (let i = 0; i < map.width; i++) {
        for (let j = 0; j < map.height; j++) {
            const tile = map.tiles[j][i];
            const x = i * tileSize;
            const y = j * tileSize;
            const konvaImage = konvaImages[tile.type].clone({
                x: x,
                y: y,
            });
            grp.add(konvaImage);
        }
    }

    mapLayer.add(grp);

    grp.toDataURL({
        x: 0,
        y: 0,
        width: map.width * tileSize,
        height: map.height * tileSize,
        callback: (dataUrl) => {
            const image = new Image();
            image.src = dataUrl;
            
            mapLayer.add(new Konva.Image({
                image: image,
                width: map.width * tileSize,
                height: map.height * tileSize,
            }))
            mapLayer.draw();
        }
    })
    
}

function createImage(src: string) {
    const image = new Image();
    image.src = src;
    return image;
}