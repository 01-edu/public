## Action - reaction!

### Resources

We provide you with some content to get started smoothly, check it out!

- Video [querySelector](https://www.youtube.com/watch?v=m34qd7aGMBo&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=12)
- Video [DOM JS - Add an event listener to an element](https://www.youtube.com/watch?v=ydRv338Fl8Y)
- Video [DOM JS - Set an element's properties](https://www.youtube.com/watch?v=4O6zSVR0ufw&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=14)
- Video [DOM JS - classList: toggle, replace & contains](https://www.youtube.com/watch?v=amEBcoTYw0s&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=20)
- Video [DOM JS - Set an element's inline style](https://www.youtube.com/watch?v=pxlYKvju1z8&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=15)
- [Memo DOM JS](https://github.com/nan-academy/js-training/blob/gh-pages/examples/dom.js)

### Instructions

OK, you have now connected HTML, CSS and JS altogether ; big day! Excited? Exhausted?
Well so far, you've only scratched the surface... Let's go deeper into the power of JS! You're going to add some interaction ; the webpage will react when a user action will happen, called an [event](https://developer.mozilla.org/en-US/docs/Web/Events) (a click, a key pressed, a mouse move, etc.).

Let's put a button on the top right corner of the page, that will toggle (close or open) the left eye when clicked.
Add it in the HTML structure:

```html
<button>Click to close the left eye</button>
```

Add this CSS style:

```css
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
- [toggle](https://css-tricks.com/snippets/javascript/the-classlist-api/) the class `eye-closed` in the `classList` of the `eye-left` HTML element
- change the background color of the `eye-left`: if the eye is open, to "red", if the eye is closed, to "black"

### Code examples

Add an event listener on click on a button that triggers a function:

```js
// events allow you to react to user inputs
// (any action with the mouse, keyboard, etc.)
// it's the foundation of the interactivity of your website
// each event is linked to an element or the window

// for this example we will attach a click event to a button
// first we select the button HTML element
const button = document.querySelector('button')

// we need to create a function
// that will be executed when the event is triggered
// let's call it `handleClick`
const handleClick = (event) => {
  // do semething when the button has been clicked
}

// register the event:
button.addEventListener('click', handleClick)
// here we ask the button to call our `handleClick` function
// on the 'click' event, so every time it's clicked
```

### Expected output

[This](https://youtu.be/Wkar5SmswDo) is what you should see in the browser.

### Notions

- [`querySelector`](https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector)
- [Text content of a HTML element](https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent)
- [`addEventListener`](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener) / [`click` event](https://developer.mozilla.org/en-US/docs/Web/API/Element/click_event#javascript)
- [`classList` / `toggle`](https://css-tricks.com/snippets/javascript/the-classlist-api/)
- [Setting style with JS](https://developer.mozilla.org/en-US/docs/Web/API/ElementCSSInlineStyle/style#setting_styles)
