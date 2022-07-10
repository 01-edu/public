## First words

### Resources

We provide you with some content to get started smoothly, check it out!

- Video [DOM JS - createElement & append](https://www.youtube.com/watch?v=J-A_pqTqGBU&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=13)
- Video [Set an element's className](https://www.youtube.com/watch?v=h3b7H1ZKvFE&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=16)
- Video [DOM JS - Set an element's properties](https://www.youtube.com/watch?v=4O6zSVR0ufw&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=14)
- Video [DOM JS - Add an event listener to an element](https://www.youtube.com/watch?v=ydRv338Fl8Y)
- [Memo DOM JS](https://github.com/nan-academy/js-training/blob/gh-pages/examples/dom.js)

### Instructions

Now that you know how to make your creation move, what about making it communicate its first words to the world?

Let's put a second button in the top right corner of the page, that will add some text when clicked.
Add it in the HTML structure:

```
<button id="speak-button">Click me to speak</button>
```

Add the button style in the CSS file:

```
button#speak-button {
  top: 100px;
}
```

Also add this class to style the text we will add:

```
.words {
  text-align: center;
  font-family: sans-serif;
}
```

In the JS file, like in the previous exercise, get the HTML button element with `id` `speak-button` and add an event listener on `click` event, triggering a function that will:

- [create a new HTML element](https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement) of type `div`
- set its [text content](https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent) to "Hello there!"
- set its [`className`](https://developer.mozilla.org/en-US/docs/Web/API/Element/className) to `words`, that we just added earlier in the CSS
- use the [`append`](https://developer.mozilla.org/en-US/docs/Web/API/ParentNode/append) method to add it inside the `torso` element

### Code examples

Create a new element and add it inside the body:

```js
// create a new `div` element
const div = document.createElement('div') // the argument passed (string) is the html tag

// select the `body` and add the new `div` inside it
const body = document.querySelector('body')
body.append(div)
```

### Expected output

[This](https://youtu.be/Eq9liRCc-zA) is what you should see in the browser.

### Notions

- [`getElementById`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById)
- [`addEventListener`](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener) / [`click` event](https://developer.mozilla.org/en-US/docs/Web/API/Element/click_event#javascript)
- [`createElement`](https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement)
- [Text content of a HTML element](https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent)
- [`className`](https://developer.mozilla.org/en-US/docs/Web/API/Element/className)
- [`append`](https://developer.mozilla.org/en-US/docs/Web/API/ParentNode/append)
