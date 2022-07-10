## Keycodes symphony

### Instructions

Like an inspired Beethoven who's going to write his Moonlight Sonata, you're about to compose a colourful symphony of letters with your keyboard.

Write the function `compose`:

- Make it fire every time a key is pressed
- Create a new `div` with the class `note` when a letter of the lowercase alphabet is pressed, which has a unique background color generated using the `key` of the `event`, and displays the corresponding letter pressed
- If the pressed key is the `Backspace` one, delete the last note
- If the pressed key is the `Escape` one, clear all the notes

### Notions

- [Keyboard event](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent): [`keydown`](https://developer.mozilla.org/en-US/docs/Web/API/Document/keydown_event), [`key`](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key)

### Files

You only need to create & submit the JS file `keycodes-symphony.js` ; we're providing you the following file to download (click right and save link) & test locally:

- the HTML file [keycodes-symphony.html](./keycodes-symphony.html) to open in the browser, which includes:

  - the JS script which will allow to run your code
  - some CSS pre-styled classes: feel free to use those as they are, or modify them

### Expected result

You can see an example of the expected result [here](https://youtu.be/5DdijwBnpAk)
