import './style.css'
import {createCanvas} from './canvas'
import {loadGame} from './observer'
import Playback from './playback';
import 'material-icons/iconfont/material-icons.css';

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <div>
    <div id="stats"></div>
    <div id="leaderboard"></div>
    <div id="mapCanvas" style="display:none;"></div>
    <div id="canvas"></div>
    <div class="bottomPanel">
      <input type="file" id="fileInput" />
      <div class="slider" id="turnSlider">
        <input type="range" />
        <p>10 / 100</p>
      </div>
      <div class="controls">
        <button id="reverse"><span class="material-icons-round">fast_rewind</span></button>
        <button id="play"><span class="material-icons-round">play_arrow</span></button>
        <button id="forward"><span class="material-icons-round">fast_forward</span></button>
      </div>
    </div>
  </div>
`


const fileInput = document.querySelector<HTMLInputElement>('#fileInput')!;
const turnSlider = document.querySelector<HTMLInputElement>('#turnSlider')!;
const slider = turnSlider.querySelector<HTMLInputElement>('input')!;
const sliderText = turnSlider.querySelector<HTMLParagraphElement>('p')!;

fileInput.addEventListener('change', async () => {
    const game = await loadGame(fileInput.files![0])
    createCanvas('canvas', game);
    turnSlider.style.display = 'flex';
    slider.max = game.length.toString();
    setSlider(0);
    new Playback(game, setSlider, slider);
    fileInput.remove();
})

function setSlider(value: number) {
    slider.value = (value + 1).toString();
    sliderText.innerHTML = `${value + 1} / ${slider.max}`;
    const v = (value + 1) / parseInt(slider.max) * 100;
    slider.style.background = 'linear-gradient(to right, cornflowerblue 0%, cornflowerblue ' + v + '%, #616161 ' + v + '%, #616161 100%)';
    console.log(slider.style.background);
}

const urlParams = new URLSearchParams(window.location.search);
if (urlParams.get('file')) {
    fetch(urlParams.get('file')!).then(async (response) => {
        turnSlider.style.display = 'flex';
        const game = await loadGame(await response.blob());
        createCanvas('canvas', game);
        slider.max = game.length.toString();
        setSlider(0);
        new Playback(game, setSlider, slider);
        fileInput.remove();
    })
}
