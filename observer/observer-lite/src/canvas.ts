import Konva from "konva";
import { createMap } from "./map";
import { Turn } from "./observer";

export let shipLayer: Konva.Layer;

export function createCanvas(id: string, turns: Turn[]): Konva.Stage {
  const stage = new Konva.Stage({
    container: id,
    width: window.innerWidth,
    height: window.innerHeight,
    draggable: true,
    offsetX: -window.innerWidth / 2,
  })
  const map = new Konva.Layer({
    imageSmoothingEnabled: false,
  })
  createMap(map, turns[0].map)
  const minScale = Math.min(window.innerHeight / (20*turns[0].map.height), window.innerWidth / (20*turns[0].map.width));
  console.log(minScale);
  
  stage.scale({x: minScale, y: minScale})
  stage.add(map)
  
  shipLayer = new Konva.Layer({
    imageSmoothingEnabled: false,
  })
  stage.add(shipLayer)
  const scaleBy = 1.1;
  stage.on('wheel', (e) => {
    // stop default scrolling
    e.evt.preventDefault();

    var oldScale = stage.scaleX();
    var pointer = stage.getPointerPosition() || {x: 0, y: 0};

    var mousePointTo = {
      x: (pointer.x - stage.x()) / oldScale,
      y: (pointer.y - stage.y()) / oldScale,
    };

    // how to scale? Zoom in? Or zoom out?
    let direction = e.evt.deltaY > 0 ? -1 : 1;

    // when we zoom on trackpad, e.evt.ctrlKey is true
    // in that case lets revert direction
    if (e.evt.ctrlKey) {
      direction = -direction;
    }

    var newScale = clamp(direction > 0 ? oldScale * scaleBy : oldScale / scaleBy, minScale, 15);

    stage.scale({ x: newScale, y: newScale });

    var newPos = {
      x: pointer.x - mousePointTo.x * newScale,
      y: pointer.y - mousePointTo.y * newScale,
    };
    stage.position(newPos);
  });

  return stage;
}

function clamp(val: number, min: number, max: number) {
  return Math.min(Math.max(val, min), max);
}