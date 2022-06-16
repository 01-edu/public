## Gossip grid

### Instructions

Create the function `grid` which displays all the `gossips`, provided in the data file below, as cards on a grid (in the same order).

They will each be represented as a `div` with the `gossip` class.

The first `gossip` card must be a `form`. It will need a `textarea`, and a submit button with the text `"Share gossip!"`. It will add new gossip to the list.

Create 3 `type="range"` inputs with the class `range`, all wrapped in a `div` with the class `ranges`.
- `id="width"`: that controls the width of cards from 200 to 800 pixels.
- `id="fontSize"`: that controls the font size from 20 to 40 pixels.
- `id="background"`: that control the background lightness from 20% to 75%.

> Use `hsl` for colors

### Files

You only need to create & submit the JS file `gossip-grid.js`; we're providing you the following files to download and test locally:

- the HTML file [gossip-grid.html](./gossip-grid.html) to open in the browser, which includes:

  - the JS script which will enable you to run your code.
  - some CSS pre-styled classes: feel free to use those as they are, or modify them.

- the data file [gossip-grid.data.js](./gossip-grid.data.js) from which you can import `gossips`.

### Expected result

You can see an example of the expected result [here](https://youtu.be/nbR2eHBqTxU)

### Notions

- [form](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Form)
- [input](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Input)
- [text](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/text)
- [range](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/range)
