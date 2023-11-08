import { shipLayer } from "./canvas";
import { leaderboard } from "./leaderboard";
import { GameMap, Turn } from "./observer";
import ShipClass from "./ship";
import Stats from "./stats";


export default class Playback {
    currentTurn: number = 0;
    turnTime = 0.2;
    playing = false;
    playInterval: number | null = null;
    ships: Record<number, ShipClass> = {};
    playButton: HTMLButtonElement;
    stats: Stats;
    static turn: Turn;
    static map: GameMap;
    constructor(private data: Turn[], private setSlider: (value: number) => void, slider: HTMLInputElement) {
        document.getElementById('forward')!.addEventListener('click', () => {
            this.next();
        });
        document.getElementById('reverse')!.addEventListener('click', () => {
            this.previous();
        });
        this.playButton = document.getElementById('play')! as HTMLButtonElement;
        this.playButton.addEventListener('click', this.togglePlay.bind(this));
        console.log(data[0]);
        window.addEventListener('keydown', (e) => {
            if (e.key == ' ') {
                this.togglePlay();
            }
            else if (e.key == 'ArrowRight') {
                this.next();
            }
            else if (e.key == 'ArrowLeft') {
                this.previous();
            }
            else if (e.key == 'Home') {
                this.seek(0);
            }
            else if (e.key == 'End') {
                this.seek(this.data.length - 1);
            }
        });
        setSlider(0);
        slider.onchange = () => {
            const v = parseInt(slider.value) / parseInt(slider.max) * 100;
            console.log(v);
            this.seek(parseInt(slider.value));
        }
        Playback.map = data[0].map;
        this.stats = new Stats();
    }

    seek(time: number) {
        this.currentTurn = time;
        this.renderTurn();
    }


    togglePlay() {
        if (this.playing) {
            this.stop();
        }
        else {
            this.play();
        }
    }

    next() {
        if (this.currentTurn == this.data.length - 1) return;
        this.currentTurn++;
        this.renderTurn();
    }

    previous() {
        if (this.currentTurn == 0) return;
        this.currentTurn--;
        this.renderTurn();
    }

    play() {
        (this.playButton.childNodes[0] as HTMLSpanElement).innerHTML = 'pause';
        this.playing = true;
        this.playInterval = setInterval(() => {
            if (this.currentTurn == this.data.length - 1) {
                this.stop();
            }
            this.next();
        }, this.turnTime * 1000);
    }

    stop() {
        (this.playButton.childNodes[0] as HTMLSpanElement).innerHTML = 'play_arrow';
        this.playing = false;
        if (this.playInterval) {
            clearInterval(this.playInterval);
        }
    }

    renderTurn() {
        const turn = this.data[this.currentTurn];
        Playback.turn = turn;
        this.stats.Update(turn);
        this.setSlider(this.currentTurn);
        leaderboard(turn.players);
        const updated = new Set<number>();
        for (const ids of Object.keys(turn.ships)) {
            const id = parseInt(ids);
            const ship = turn.ships[id];
            console.log(ship);
            updated.add(id);
            if (!this.ships[id]) {
                this.ships[id] = new ShipClass(ship, shipLayer, 20, turn.ship_types[ship.index]);
            }
            else {
                this.ships[id].move(ship.x, ship.y);
                this.ships[id].setHealth(ship.health);
                if (ship.is_wreck) {
                    this.ships[id].setWreck();
                }
            }
        }

        for (const id of Object.keys(this.ships)) {
            const ship = parseInt(id);
            if (!updated.has(ship)) {
                this.ships[ship].remove();
                delete this.ships[ship];
            }
        }
    }
}