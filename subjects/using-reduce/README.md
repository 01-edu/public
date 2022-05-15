## Using Reduce

### Instructions

Create the following functions:

> Your solutions **must** use `reduce`.

- `adder`: accepts an array of numbers, and returns the sum as a `number`.

- `sumOrMul`: accepts an array of numbers and adds or multiplies its elements depending on whether the element is odd or even. Even = multiply. Odd = add.

- `funcExec`: accepts an array of functions and executes them using `reduce`, returning the result.

> Each function may accept an optional argument, which should be the initial value for the function's execution.

#### Example:

```js
sumOrMul([1, 2, 3, 5, 8], 5) // (((((5 + 1) * 2) + 3) + 5) * 8) -> 160
```

### Notions

- [Array.prototype.reduce](https://devdocs.io/javascript/global_objects/array/reduce)
