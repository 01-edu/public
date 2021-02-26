### Bring it to life

Still there? Well done! But hold on, here begins the serious part... In order to control your creation, you're going to plug its brain: JavaScript.

[`Link a JS script`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script) to your `bring-it-to-life.html` file.
In this script, you're going to close the left eye of your entity.
To do so, you have to target the `eye-left` HTML element by its `id` thanks to the [`getElementById`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById) method. Then, change the `style` of your `eye-left` to set its background color to black. We also need to modify its shape ; for that we are going to add a new class to it.
First, define it in your CSS file:

```
.eye-closed {
  height: 2px;
  padding: 0 5px;
  border-radius: 10px;
}
```

In the JS file, add the freshly-created class `eye-closed` to the [`classList`](https://css-tricks.com/snippets/javascript/the-classlist-api/) of your `eye-left` element.

Reload the page - it's alive! Your JS brain has control and orders your HTML/CSS body to close one eye.

### Expected output

This is what you should see in the browser:

<!-- screenshot -->

### Notions

- [Link a JS script](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script)
- [`getElementById`](https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById)
- [`classList` / `add()`](https://css-tricks.com/snippets/javascript/the-classlist-api/)
