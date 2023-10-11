import './style.css'
import { createCanvas } from './canvas'
import { loadGame } from './observer'
import Playback from './playback';

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <div>
    <div id="shipStats"></div>
    <div id="canvas"></div>
    <div class="bottomPanel">
      <input type="file" id="fileInput" />
      <input type="range" id="turnSlider" class="slider" />
      <div class="controls">
        <button id="reverse"><span class="material-symbols-outlined">fast_rewind</span></button>
        <button id="play"><span class="material-symbols-outlined">play_arrow</span></button>
        <button id="forward"><span class="material-symbols-outlined">fast_forward</span></button>
      </div>
    </div>
  </div>
`


const fileInput = document.querySelector<HTMLInputElement>('#fileInput')!;
const turnSlider = document.querySelector<HTMLInputElement>('#turnSlider')!;
fileInput.addEventListener('change', async () => {
  const game = await loadGame(fileInput.files![0])
  createCanvas('canvas', game);
  turnSlider.max = game.length.toString();
  turnSlider.value = '0';
  turnSlider.style.display = 'block';
  new Playback(game, turnSlider);
  fileInput.remove();
})

const urlParams = new URLSearchParams(window.location.search);
if(urlParams.get('file')) {
  fetch(urlParams.get('file')!).then(async (response) => {
    const game = await loadGame(await response.blob());
    createCanvas('canvas', game);
    turnSlider.max = game.length.toString();
    turnSlider.value = '0';
    turnSlider.style.display = 'block';
    new Playback(game, turnSlider);
    fileInput.remove();
  })
}
