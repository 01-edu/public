## Gossip grid

### Instructions

Good information is the pillar of society, that's why you've decided to dedicate your time to reveal the powerful truth to the world and deliver essential and strong news: you're launching a gossip grid.

Create the function `grid` which displays all the `gossips`, provided in the data file below, as cards on a grid (in the same order).
They must be `div` with the `gossip` class.

The first `gossip` card must be a `form` with a `textarea` and a submit button with the text `Share gossip!` that allows to add a new gossip to the list.

Create 3 `type="range"` inputs with the class `range`, all wrapped in a `div` with the class `ranges`:

- one with `id="width"` that control the width of cards _(from 200 to 800 pixels)_
- one with `id="fontSize"` that control the font size _(from 20 to 40 pixels)_
- one with `id="background"` that control the background lightness _(from 20% to 75%)_

> _tips:_ use `hsl` for colors

### Notions

- [`<form>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Form)
- [`<input>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Input): [`text`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/text), [`range`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/range)

### Files

You only need to create & submit the JS file `gossip-grid.js` ; we're providing you the following files to download (click right and save link) & test locally:

- the HTML file [gossip-grid.html](./gossip-grid.html) to open in the browser, which includes:

  - the JS script which will allow to run your code
  - some CSS pre-styled classes: feel free to use those as they are, or modify them

- the data file [gossip-grid.data.js](./gossip-grid.data.js) from which you can import `gossips`

### Expected result

You can see an example of the expected result [here](https://youtu.be/nbR2eHBqTxU)
