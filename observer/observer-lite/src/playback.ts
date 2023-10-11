import { shipLayer } from "./canvas";
import { Turn } from "./observer";
import ShipClass from "./ship";


export default class Playback {
    currentTurn: number = 0;
    turnTime = 0.2;
    playing = false;
    playInterval: number | null = null;
    ships: Record<number, ShipClass> = {};
    playButton: HTMLButtonElement;
    static turn: Turn;
    constructor(private data: Turn[], private slider: HTMLInputElement) {
        document.getElementById('forward')!.addEventListener('click', () => {
            this.next();
        });
        document.getElementById('reverse')!.addEventListener('click', () => {
            this.previous();
        });
        this.playButton = document.getElementById('play')! as HTMLButtonElement;
        this.playButton.addEventListener('click', () => {
            if(this.playing) {
                this.stop();
            }
            else {
                this.play();
            }
        });
        slider.onchange = () => {
            this.seek(parseInt(slider.value));
        }
    }

    seek(time: number) {
        this.currentTurn = time;
        this.renderTurn();
    }

    next() {
        if(this.currentTurn == this.data.length - 1) return;
        this.currentTurn++;
        this.renderTurn();
    }

    previous() {
        if(this.currentTurn == 0) return;
        this.currentTurn--;
        this.renderTurn();
    }

    play() {
        (this.playButton.childNodes[0] as HTMLSpanElement).innerHTML = 'pause';
        this.playing = true;
        this.playInterval = setInterval(() => {
            if(this.currentTurn == this.data.length - 1) {
                this.stop();
            }
            this.next();
        }, this.turnTime * 1000);
    }

    stop() {
        (this.playButton.childNodes[0] as HTMLSpanElement).innerHTML = 'play_arrow';
        this.playing = false;
        if(this.playInterval){
            clearInterval(this.playInterval);
        }
    }

    renderTurn() {
        const turn = this.data[this.currentTurn];
        Playback.turn = turn;
        this.slider.value = this.currentTurn.toString();
        const updated = new Set<number>();
        for (const ids of Object.keys(turn.ships)) {
            const id = parseInt(ids);
            const ship = turn.ships[id];
            console.log(ship);
            updated.add(id);
            if(!this.ships[id]) {
                this.ships[id] = new ShipClass(ship, shipLayer, 20);
            }
            else {
                this.ships[id].move(ship.x, ship.y);
            }
        }       

        for (const id of Object.keys(this.ships)) {
            const ship = parseInt(id);
            if(!updated.has(ship)) {
                this.ships[ship].remove();
                delete this.ships[ship];
            }
        }
    }
}