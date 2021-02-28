### First words

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

### Expected output

[This](https://youtu.be/Eq9liRCc-zA) is what you should see in the browser.

### Notions

- [`getElementById`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById)
- [`addEventListener`](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener) / [`click` event](https://developer.mozilla.org/en-US/docs/Web/API/Element/click_event#javascript)
- [`createElement`](https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement)
- [Text content of a HTML element](https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent)
- [`className`](https://developer.mozilla.org/en-US/docs/Web/API/Element/className)
- [`append`](https://developer.mozilla.org/en-US/docs/Web/API/ParentNode/append)
