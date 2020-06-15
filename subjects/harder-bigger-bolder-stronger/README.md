## Harder, bigger, bolder, stronger

### Instructions

Being stuck at home, bored, desperate and coming up with a lot of weird ideas, a friend asks you to develop a tool to measure his ocular skills: one of those [Monoyer charts](https://en.wikipedia.org/wiki/Monoyer_chart) that ophthalmologists use.

Generate a board where each new letter is harder, bigger, bolder and stronger!

Create the function `generateLetters` which creates 100 `div`, each containing a letter randomly picked through the alphabet, and whose style properties have to be increased:

- `font-size` has to grow from `20` to at least `100` pixels
- `font-weigth` has to be `300` for the first third of the letters, `400` for the second third, and `600` for the last third

### Notions

- [`style`](https://developer.mozilla.org/en-US/docs/Web/API/ElementCSSInlineStyle/style)
- [`textContent`](https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent)
