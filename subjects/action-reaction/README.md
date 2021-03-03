### Action - reaction!

OK, you have now connected HTML, CSS and JS altogether ; big day! Excited? Exhausted?
Well so far, you've only scratched the surface... Let's go deeper into the power of JS! You're going to add some interaction ; the webpage will react when a user action will happen, called an [event](https://developer.mozilla.org/en-US/docs/Web/Events) (a click, a key pressed, a mouse move, etc.).

Let's put a button on the top right corner of the page, that will toggle (close or open) the left eye when clicked.
Add it in the HTML structure:

```
<button>Click to close the left eye</button>
```

Add the style in the CSS file:

```
button {
  z-index: 1;
  position: fixed;
  top: 30px;
  right: 30px;
  padding: 20px;
}
```

In the JS file, get the HTML button element with [`querySelector`](https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector), and [add an event listener](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener) on [`click` event](https://developer.mozilla.org/en-US/docs/Web/API/Element/click_event#javascript), triggering a function that will:

- change the [text content](https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent) of the button: if the eye is open, write "Click to close the left eye", if the eye is closed, write "Click to open the left eye"
- [toggle](https://css-tricks.com/snippets/javascript/the-classlist-api/) the class `eye-open` in the `classList` of the `eye-left` HTML element
- change the background color of the `eye-left`: if the eye is open, to "orange", if the eye is closed, to "black"

### Expected output

[This](https://youtu.be/Wkar5SmswDo) is what you should see in the browser.

### Notions

- [`querySelector`](https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector)
- [Text content of a HTML element](https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent)
- [`addEventListener`](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener) / [`click` event](https://developer.mozilla.org/en-US/docs/Web/API/Element/click_event#javascript)
- [`classList` / `toggle`](https://css-tricks.com/snippets/javascript/the-classlist-api/)
- [Setting style with JS](https://developer.mozilla.org/en-US/docs/Web/API/ElementCSSInlineStyle/style#setting_styles)
