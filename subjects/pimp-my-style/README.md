## Pimp my style

### Instructions

Check out that button on the HTML page:

```html
<button class="button">pimp my style</div>
```

For now, it's only a lonely, basic and sad element ; let's pimp it up!

On each click on the page, a function `pimp` is triggered.
Write the body of that function so that the button's class is altered:

- Add in order the next class of the `styles` array provided in the data file below
- When the end of the array is reached, remove backwards each class
- Toggle the class 'unpimp' when removing classes

```
Example for a `styles` array with only 3 classes:

Page load --> <button class="button"></div>

...adding
Click 1 --> <button class="button one"></div>
Click 2 --> <button class="button one two"></div>

...toggling `unpimp`
Click 3 --> <button class="button one two three unpimp"></div>

...and removing backwards
Click 4 --> <button class="button one two unpimp"></div>
Click 5 --> <button class="button one unpimp"></div>
Click 6 --> <button class="button"></div>
```

### Notions

- [`classList`](https://developer.mozilla.org/en-US/docs/Web/API/Element/classList): `add()`, `remove()`, `toggle()`

### Provided files

- Check the HTML file [index.html](/public/subjects/pimp-my-style/index.html), which includes:

  - the JS script running some code, and which will also allow to run yours
  - some CSS pre-styled classes: feel free to use those as they are, or modify them

- Import `styles` from the data file [data.js](/public/subjects/pimp-my-style/data.js)

### Expected result

You can see an example of the expected result [here](https://youtu.be/VIRf3TBDTN4)
