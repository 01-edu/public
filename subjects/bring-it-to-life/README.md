## Bring it to life

### Resources

We provide you with some content to get started smoothly, check it out!

- Video [Link a JS script to your HTML file](https://www.youtube.com/watch?v=jMvsQm-p1gM&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=7)
- Video [DOM JS - getElementById](https://www.youtube.com/watch?v=34kAR8yBtDM&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=8)
- Video [DOM JS - Set an element's inline style](https://www.youtube.com/watch?v=pxlYKvju1z8&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=15)
- Video [DOM JS - classList: add & remove](https://www.youtube.com/watch?v=uQEM-3_4vPA&list=PLHyAJ_GrRtf979iZZ1N3qYMfsPj9PCCrF&index=17)
- [Memo DOM JS](https://github.com/nan-academy/js-training/blob/gh-pages/examples/dom.js)

### Instructions

Still there? Well done! But hold on, here begins the serious part... In order to control your creation, you're going to plug its brain: JavaScript.

[`Link a JS script`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script) to your HTML file.
In this script, you're going to close the left eye of your entity.
To do so, you have to target the `eye-left` HTML element by its `id` thanks to the [`getElementById`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById) method. Then, [set the `style`](https://developer.mozilla.org/en-US/docs/Web/API/ElementCSSInlineStyle/style#setting_styles) of your `eye-left` to set its background color to "black". We also need to modify its shape ; for that we are going to add a new class to it.
First, define this new class in your CSS file:

```
.eye-closed {
  height: 4px;
  padding: 0 5px;
  border-radius: 10px;
}
```

In the JS file, add the freshly-created class `eye-closed` to the [`classList`](https://css-tricks.com/snippets/javascript/the-classlist-api/) of your `eye-left` element.

Reload the page - it's alive! Your JS brain has control and orders your HTML/CSS body to close one eye.

### Code examples

Get a HTML element by its `id` & set its inline style:

```js
const myDiv = document.getElementById('my-div')
myDiv.style.color = 'green'
```

### Expected output

This is what you should see in the browser:
![](https://github.com/01-edu/public/raw/master/subjects/bring-it-to-life/bring-it-to-life.png)

### Notions

- [Link a JS script](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script)
- [`getElementById`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById)
- [`classList` / `add()`](https://css-tricks.com/snippets/javascript/the-classlist-api/)
- [Setting style with JS](https://developer.mozilla.org/en-US/docs/Web/API/ElementCSSInlineStyle/style#setting_styles)
