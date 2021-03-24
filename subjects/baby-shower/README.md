## Baby shower

### Instructions

It's time to celebrate the arrival of your creation in this world ; let's organize a baby shower!

But you already worked a lot to give birth, so now you need some help from others to set the party.
You're going to use a JS library - meaning some code made by other people, made available for anyone to use. Very cool & helpful, isn't it?
There is a huge amount of packages of any type - really any - and it's usually quite simple to install and use.
The one we're going to use today is [canvas confetti][0].

From your script, [`import`][1] the package as indicated in the documentation, and launch the function to pop confetti when the page loads.
Then, [add an event listener][2] on [`keydown` event][3] that triggers the same confetti popping function any time a key of the keyboard is pressed.

### Code examples

Import code from libraries & use it:

```js
// import and use the `hello` function from a library
import { hello } from 'hello-library-url'

hello()
```

### Expected output

This is what you should see in the browser:

![](https://raw.githubusercontent.com/01-edu/public/master/subjects/baby-shower/baby-shower.gif)

### Notions

- [`import`][1]
- [`addEventListener`][2] / [`keydown` event][3]

[0]: https://www.skypack.dev/view/canvas-confetti
[1]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import
[2]: https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
[3]: https://developer.mozilla.org/en-US/docs/Web/API/Element/keydown_event#javascript
