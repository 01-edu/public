## Fifty shades of cold

### Instructions

You've been asked to freshen a webpage atmosphere by displaying shades of cold colors.

Check the `colors` array provided in the data file below.

- Write the `generateClasses` function which creates a `<style>` tag in the `<head>` tag and generates, for each color of `colors`, a class setting the `background` attribute and taking the color as value, like following:

```css
.blue {
  background: blue;
}
```

- Write the `generateColdShades` function which creates a `<div>` for each color of the `colors` array whose name contains `aqua`, `blue`, `turquoise`, `green`, `cyan`, `navy` or `purple`.\
  Each `<div>` must have the corresponding generated class and display the name of the color, like following:

```html
<div class="blue">blue</div>
```

- The function `choseShade` is triggered when clicking on a `div`.\
  Write the body of this function, which receives the shade of the clicked element as argument, and replaces all the other elements class by the chosen shade.

### Notions

- [`head`](https://developer.mozilla.org/en-US/docs/Web/API/Document/head) / [style tag](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style)
- [`className`](https://developer.mozilla.org/en-US/docs/Web/API/Element/className)
- [`classList`](https://developer.mozilla.org/en-US/docs/Web/API/Element/classList): `contains()`, `replace()`

### Files

You only need to create & submit the JS file `fifty-shades-of-cold.js` ; we're providing you the following files to download (click right and save link) & test locally:

- the HTML file [fifty-shades-of-cold.html](./fifty-shades-of-cold.html) to open in the browser, which includes:

  - the JS script running some code, and which will also allow to run yours
  - some CSS pre-styled classes: feel free to use those as they are, or modify them

- the data file [fifty-shades-of-cold.data.js](./fifty-shades-of-cold.data.js) from which you can import `colors`

### Expected result

You can see an example of the expected result [here](https://youtu.be/a-3JDEvW-Qg)
