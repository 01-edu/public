## Pimp my style

### Instructions

Check out that button on the HTML page:

```html
<button class="button">pimp my style</div>
```

For now, it's only a lonely, basic and sad element. let's pimp it up.

On each click on the page, a function `pimp` is triggered.
Write the body of that function so that the button's class is altered:

- From the data file provided, add each of the `styles` array elements as classes, in order.
- When the end of the array is reached, remove the classes in a LIFO fashion.
- While removing classes, toggle the `unpimp` class on. And toggle it off again while adding classes.


Example for a `styles` array with only 3 classes:
```js
["one", "two", "three"]
```

```
Page load --> <button class="button"></button>

...adding
Click 1 --> <button class="button one"></button>
Click 2 --> <button class="button one two"></button>

...toggling `unpimp`
Click 3 --> <button class="button one two three unpimp"></button>

...and removing backwards
Click 4 --> <button class="button one two unpimp"></button>
Click 5 --> <button class="button one unpimp"></button>
Click 6 --> <button class="button"></button>
```

### Files

You only need to create & submit the JS file `pimp-my-style.js`. We're providing you the following files to download and test locally.:

- the HTML file [pimp-my-style.html](./pimp-my-style.html) to open in the browser, which includes:

  - the JS script running some code, and which will enable you to run yours.
- the data file [pimp-my-style.data.js](./pimp-my-style.data.js) from which you can import `styles`.

- feel free to use the CSS file [pimp-my-style.data.css](./pimp-my-style.data.css) as it is or you can also modify it.

### Expected result

You can see an example of the expected result [here](https://youtu.be/VIRf3TBDTN4)

### Notions

- [classList](https://developer.mozilla.org/en-US/docs/Web/API/Element/classList): `add()`, `remove()`, `toggle()`
