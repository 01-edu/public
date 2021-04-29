## Total

JS also provide another type of loop than the `while`

### for..of

`for..of` loops are used to iterate _over_ a value, for example you can iterate
over an array:

```js
let heroes = ['Superman', 'Batman', 'Onepuchman', 'Spiderman', 'Ultraman']

// ↙ the for of  ↙  keywords
for (let heroe of heroes) {
  // <- the value we iterate over
  //  ↖ variable declaration, will be each elements from the value
  // in this case, heroes, an array
  console.log(heroe)
}
```

From the value we have used so far, only `string` and `array` are iterable.

### Instructions

Declare the function `total` that compute the sum of all the values from the
given array

If the array is empty, `total` should return 0

**Example:**

```js
total([1, 1, 3]) // -> 5
total([6, 2, -3]) // -> 5
total([1, 2, 2]) // -> 5
total([0, -2]) // -> -2
```
