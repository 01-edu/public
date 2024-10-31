## Mouse trap

### Instructions

Develop a trap to capture the elements when the mouse is getting too close to the center of the page.

Create the following functions:

- `createCircle`: make it fire on every click on the page, and create a `div` at the position of the mouse on the screen, setting its `background` to `white` and its class to `circle`.

- `moveCircle`: make it fire when the mouse moves, and get the last circle created and makes it move along with the mouse.

- `setBox`: which creates a box with the class `box` in the center of the page. When a circle is entirely inside that box, it has to be purple (use the CSS global variable `var(--purple)` as its `background`). Once a circle enters the box, it is trapped inside and cannot escape.

> Hint: Be careful, a circle cannot overlap the box which has walls of `1px`. It has to be trapped **strictly** inside.

### Provided files

### Files

You only need to create & submit the JS file `mouse-trap.js`, we're providing you the following files to download and test locally:

- the HTML file [mouse-trap.html](./mouse-trap.html) to open in the browser, which includes:

  - the JS script which will enable you to run your code.

- feel free to use the CSS file [mouse-trap.data.css](./mouse-trap.data.css) as it is or you can also modify it.

### Expected result

You can see an example of the expected result [here](https://youtu.be/qF843P-V2Yw)

### Notions

- [addEventListener](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener)
- [removeEventListener](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener)
- [Mouse event](https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/MouseEvent)
- [click](https://developer.mozilla.org/en-US/docs/Web/API/Element/click_event)
- [mousemove](https://developer.mozilla.org/en-US/docs/Web/API/Element/mousemove_event)
- [clientX](https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/clientX)
- [clientY](https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent/clientY)
- [getBoundingClientRect](https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect)
