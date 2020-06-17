## Build brick and break

### Instructions

Today, your mission is to build a 3-column brick tower, maintain it and finally break it!

Create a function `build` which will create and display the given amount of bricks passed as argument:

- each brick has to be created and added to the page at a regular interval of time (at least 50ms),
- each brick will receive a unique `id` property, like following:

```html
<div id="brick-1"></div>
```

- each brick in the middle column has to be set with the custom attribute `foundation` receiving the value `true`

Each one of the two emojis in the top-right corner fires a function on click:

- 🔨 triggers the function `repair`
- 🧨 triggers the function `destroy`

Write the body of the `repair` function, which receives any number of `ids`, and for each `id`, retrieves the HTML element and set a custom attribute `repaired` set to `in progress` if it is a brick situated in the middle column, and `true` if not.

Write the body of the `destroy` function, which removes the current last brick in the tower.

### Notions

- [`createElement()`](https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement)
- [`append()`](https://developer.mozilla.org/fr/docs/Web/API/ParentNode/append)
- [Element](https://developer.mozilla.org/en-US/docs/Web/API/Element)
- [`setInterval()`](https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/setInterval) / [`clearInterval()`](https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/clearInterval)
- [`hasAttribute()`](https://developer.mozilla.org/en-US/docs/Web/API/Element/hasAttribute)
- [dataset](https://developer.mozilla.org/en-US/docs/Web/API/HTMLOrForeignElement/dataset)
- [`remove()`](https://developer.mozilla.org/en-US/docs/Web/API/ChildNode/remove)

### Provided files

- Use this CSS file: [style.css](./style.css)

### Expected result

You can see an example of the expected result [here](https://youtu.be/OjSP_7u9CZ4)
