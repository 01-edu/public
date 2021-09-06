## Pick & click

### Instructions

Today, you're gonna create your own color picker.

Write the function `pick` which turns the screen into a `hsl` color picker, varying the `hue` and `luminosity` of according to the position of the mouse, which:

- changes the `background` color of the `body`, so the `hsl` value is different on each mouse position on the screen:
  - on the axis X, the hue value has to vary between 0 and 360
  - on the axis Y, the luminosity value has to vary between 0 and 100
- displays those 3 values using the `text` class:
  - the full `hsl` value in a `div` with the class `hsl` in the middle of the screen
  - the `hue` value in a `div` with the class `hue` in the top right corner of the screen
  - the `luminosity` value in a `div` with the class `luminosity` in the bottom left corner of the screen
- copies that value in the clipboard on click
- displays two SVG lines, with respective ids `axisX` and `axisY`, following the cursor like so:
  - the axisX has to set the attributes `x1` and `x2` to the mouse X position
  - the axisY has to set the attributes `y1` and `y2` to the mouse Y position

> Here is how a hsl value is formatted: `hsl(45, 50%, 35%)`

> Use `Math.round()` to round the values

### Notions

- [Copy event](https://developer.mozilla.org/en-US/docs/Web/API/Element/copy_event)
- [Mouse move event](https://developer.mozilla.org/en-US/docs/Web/API/Element/mousemove_event)
- [SVG](https://developer.mozilla.org/en-US/docs/Web/SVG/Element/svg): [`createElementNS`](https://developer.mozilla.org/en-US/docs/Web/API/Document/createElementNS), [`setAttribute`](https://developer.mozilla.org/en-US/docs/Web/API/Element/setAttribute)
- Take a look at the [HSL section](https://developer.mozilla.org/en-US/docs/Web/HTML/Applying_color)

### Files

You only need to create & submit the JS file `pick-and-click.js` ; we're providing you the following file to download (click right and save link) & test locally:

- the HTML file [pick-and-click.html](./pick-and-click.html) to open in the browser, which includes:

  - the JS script which will allow to run your code
  - some CSS pre-styled classes: feel free to use those as they are, or modify them

### Expected result

You can see an example of the expected result:
[![video](https://img.youtube.com/vi/eE4eE9_eKZI/0.jpg)](https://www.youtube.com/watch?v=eE4eE9_eKZI)
