## Pick & click

### Instructions

Today, you're gonna create your own color picker.

Write the function `pick` which creates a `hsl` color picker varying the `hue` and `luminosity` according to the position of the mouse, which:

- displays the `hue` value in text
- displays the `luminosity` value in text
- displays the full `hsl` value in text
- copies that value in the clipboard on click
- displays two lines, for X and Y axis, following the cursor

### Notions

- [Copy event](https://developer.mozilla.org/en-US/docs/Web/API/Element/copy_event)
- [Mouse move event](https://developer.mozilla.org/en-US/docs/Web/API/Element/mousemove_event)
- [SVG](https://developer.mozilla.org/en-US/docs/Web/SVG/Element/svg): [`createElementNS`](https://developer.mozilla.org/en-US/docs/Web/API/Document/createElementNS), [`setAttribute`](https://developer.mozilla.org/en-US/docs/Web/API/Element/setAttribute)
- Take a look at the [HSL section](https://developer.mozilla.org/en-US/docs/Web/HTML/Applying_color)

### Provided files

- Use the HTML file [index.html](/public/subjects/pick-and-click/index.html), which includes:

  - the JS script which will allow to run your code
  - some CSS pre-styled classes: feel free to use those as they are, or modify them

### Expected result

You can see an example of the expected result:
[![video](https://img.youtube.com/vi/eE4eE9_eKZI/0.jpg)](https://www.youtube.com/watch?v=eE4eE9_eKZI)
