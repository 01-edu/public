## Pick & click

### Instructions

Today, you're going to create your own color picker.

Write the function `pick` which turns the screen into a `hsl` color picker. It will vary the `hue` and `luminosity` according to the position of the mouse.

The `background` color of the `body` will change based on the position of the mouse on the screen.
  - The X axis will vary the hue value between 0 and 360.
  - The Y axis will vary the luminosity value between 0 and 100.

You'll need to display these three values:
- The full `hsl` value in a `div`, which has the class `hsl` in the middle of the screen.
- The `hue` value in a `div` with the class `hue` and `text` in the top right corner of the screen.
- The `luminosity` value will be displayed in the bottom left corner of the screen, in a `div` with the class `luminosity` and `text`.

When the mouse is clicked, the value of the `hsl` will need to be copied to the clipboard.

Two SVG lines with ids `axisX` and `axisY` will need to follow the cursor, like really long cross hairs.
  - the `axisX` attributes `x1` and `x2` need to be set to the X position of the cursor.
  - the `axisY` attributes `y1` and `y2` need to be set to the Y position of the cursor.

> The formatting of a `hsl` value: `hsl(45, 50%, 35%)`.

> Use `Math.round()` to round the values.

### Files

You only need to create & submit the JS file `pick-and-click.js`; we're providing you the following file to download and test locally:

- the HTML file [pick-and-click.html](./pick-and-click.html) to open in the browser, which includes:

  - the JS script which will enable you to run your code.
  - some CSS pre-styled classes: feel free to use those as they are, or modify them.

### Expected result

You can see an example of the expected result [here](https://www.youtube.com/watch?v=eE4eE9_eKZI)

### Notions

- [Copy event](https://developer.mozilla.org/en-US/docs/Web/API/Element/copy_event)
- [Mouse move event](https://developer.mozilla.org/en-US/docs/Web/API/Element/mousemove_event)
- [SVG](https://developer.mozilla.org/en-US/docs/Web/SVG/Element/svg)
- [createElementNS](https://developer.mozilla.org/en-US/docs/Web/API/Document/createElementNS)
- [setAttribute](https://developer.mozilla.org/en-US/docs/Web/API/Element/setAttribute)
- Take a look at the [HSL section](https://developer.mozilla.org/en-US/docs/Web/HTML/Applying_color)
