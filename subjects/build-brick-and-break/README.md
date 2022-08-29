## Build brick and break

### Instructions

Today, your mission is to build a 3-column brick tower, maintain it and finally break it.

- Create a function `build` which will create and display the amount of bricks passed as argument:

  - each brick has to be created as a `div` and added to the page at a regular interval of 100ms.
  - each brick will receive a unique `id` property, like the following:

  ```html
  <div id="brick-1"></div>
  ```

  - each brick in the middle column has to be set with the custom data attribute `foundation`, receiving the value `true`.

- Each one of the two emojis in the top-right corner fires a function on click:
  - 🔨: triggers the function `repair`. Write the body of that function. It receives any number of `ids`. For each `id`, it retrieves the HTML element, and sets the `repaired` custom attribute to `in progress` if it is a brick situated in the middle column, and `true` if not.
  - 🧨: triggers the `destroy` function. Write the body of that function. It removes the current last brick in the tower.

### Files

You only need to create & submit the JS file `build-brick-and-break.js`, We're providing you the following file to download and test locally:

- the HTML file [build-brick-and-break.html](./build-brick-and-break.html) can be opened in the browser, which includes:
  - the JS script running some code, and which will enable you to run your code.
  - some CSS pre-styled classes: feel free to use those as they are, or modify them.

### Expected result

You can see an example of the expected result [here](https://youtu.be/OjSP_7u9CZ4)

### Notions

- [createElement](https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement)
- [append](https://developer.mozilla.org/en-US/docs/Web/API/ParentNode/append)
- [Element](https://developer.mozilla.org/en-US/docs/Web/API/Element)
- [setInterval](https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/setInterval) / [clearInterval](https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/clearInterval)
- [hasAttribute](https://developer.mozilla.org/en-US/docs/Web/API/Element/hasAttribute)
- [dataset](https://developer.mozilla.org/en-US/docs/Web/API/HTMLOrForeignElement/dataset) / [data-*](https://developer.mozilla.org/en-US/docs/Web/HTML/Global_attributes/data-*)
- [remove](https://developer.mozilla.org/en-US/docs/Web/API/ChildNode/remove)
