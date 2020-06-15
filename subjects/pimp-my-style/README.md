## Pimp my style

### Instructions

Check out that button on the HTML page:

```html
<div class="button">pimp my style</div>
```

For now, it's only a lonely, basic and sad element ; let's pimp it up!

On each click on the page, a function `pimp` is triggered.
Write the body of that function so that the button's class is altered:

- Add in order the next class of the `styles` array provided in the data file below
- When the end of the array is reached, remove backwards each class
- Toggle the class 'unpimp' when removing classes

```
Example for a `styles` array with only 3 classes:

Page load --> <div class="button"></div>

...adding
Click 1 --> <div class="button one"></div>
Click 2 --> <div class="button one two"></div>

...toggling `unpimp`
Click 3 --> <div class="button one two three unpimp"></div>

...and removing backwards
Click 4 --> <div class="button one two unpimp"></div>
Click 5 --> <div class="button one unpimp"></div>
Click 6 --> <div class="button"></div>
```

### Notions

- [`classList`](https://developer.mozilla.org/en-US/docs/Web/API/Element/classList): `add()`, `remove()`, `toggle()`

### Provided files

- Use this CSS file: [https://mariemalarme.github.io/dom-js/assets/style/pimp-my-style.css](https://mariemalarme.github.io/dom-js/assets/style/pimp-my-style.css)
- Import the `styles` from the data file: [https://mariemalarme.github.io/dom-js/assets/data/pimp-my-style.js](https://mariemalarme.github.io/dom-js/assets/data/pimp-my-style.js)

### Expected result

You can see an example of the expected result [here](https://youtu.be/VIRf3TBDTN4)
