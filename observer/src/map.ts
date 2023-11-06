import Konva from "konva";
import harbor from "./assets/Harbor.png";
import base from "./assets/Base.png";
import { GameMap } from "./observer";
import Stats from "./stats";
import Playback from "./playback";

export function createMap(mapLayer: Konva.Layer, shipLayer: Konva.Layer, map: GameMap) {
    const tileSize = 20;

    const grp = new Konva.Group({
        listening: false
    });

    for (let i = 0; i < map.width; i++) {
        for (let j = 0; j < map.height; j++) {
            const tile = map.tiles[j][i];
            const x = i * tileSize;
            const y = j * tileSize;
            const img = new Image();

            if (tile.type === 1 || tile.type === 3) {
                img.src = "./mapImages/ground" + getRandomInt(1, 4) + ".png";
            } else {
                const dirX = [1, 1, 0, -1, -1, -1, 0, 1];
                const dirY = [0, 1, 1, 1, 0, -1, -1, -1];
                const dirName = ["R", "BR", "B", "BL", "L", "TL", "T", "TR"];

                const dirs = [];
                for (let k = 0; k < dirName.length; k++) {
                    const tile = tryGetTile(i + dirX[k], j + dirY[k], map)
                    if (tile === 1 || tile === 3) {
                        dirs.push(dirName[k]);
                    }
                }

                if (dirs.includes("R")) {
                    dirs.remove("BR");
                    dirs.remove("TR");
                }
                if (dirs.includes("L")) {
                    dirs.remove("BL");
                    dirs.remove("TL");
                }
                if (dirs.includes("T")) {
                    dirs.remove("TR");
                    dirs.remove("TL");
                }
                if (dirs.includes("B")) {
                    dirs.remove("BR");
                    dirs.remove("BL");
                }

                dirs.sort();
                if (dirs.length === 0) {
                    img.src = "./mapImages/water.png";
                } else {

                    img.src = "./mapImages/water-" + dirs.join(',') + ".png";
                }
            }
            const konvaImage = new Konva.Image({
                x: x,
                y: y,
                image: img,
                width: tileSize,
                height: tileSize,
                listening: false
            });
            grp.add(konvaImage);

            if (tile.type > 1) {
                const img = new Image();
                if (tile.type === 2) {
                    img.src = harbor;
                } else if (tile.type === 3) {
                    img.src = base;
                }

                const konvaImage = new Konva.Image({
                    x: x,
                    y: y,
                    image: img,
                    width: tileSize,
                    height: tileSize
                });

                if (tile.type === 2) {
                    konvaImage.on('click', () => {
                        const harborStats = Playback.turn.harbors.find((harbor) => {
                            return harbor.x === i && harbor.y === j;
                        });
                        Stats.showHarborStats(harborStats || null);
                    });
                }
                if (tile.type === 3) {
                    konvaImage.on('click', () => {
                        const baseStats = Playback.turn.bases.find((base) => {
                            return base.x === i && base.y === j;
                        });
                        Stats.showBaseStats(baseStats || null);
                    });
                }
                shipLayer.add(konvaImage);
            }
        }
    }

    mapLayer.add(grp);

    // grp.toDataURL({
    //     x: 0,
    //     y: 0,
    //     width: map.width * tileSize,
    //     height: map.height * tileSize,
    //     callback: (dataUrl) => {
    //         const image = new Image();
    //         image.src = dataUrl;

    //         mapLayer.add(new Konva.Image({
    //             image: image,
    //             width: map.width * tileSize,
    //             height: map.height * tileSize,
    //         }))
    //         mapLayer.draw();
    //     }
    // })

}

function tryGetTile(i: number, j: number, map: GameMap) {
    if (i < 0 || i >= map.width || j < 0 || j >= map.height) {
        return 0;
    }
    return map.tiles[j][i].type;
}

function getRandomInt(min: number, max: number) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

Object.defineProperty(Array.prototype, "remove", {
    value: function <T>(this: T[], item: T) {
        var index = this.indexOf(item);
        if (index !== -1) {
            this.splice(index, 1);
        }
    }
});

declare global {
    interface Array<T> {
        remove(item: T): void;
    }
}
