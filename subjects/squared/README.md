## Squared

`.filter` is not the only useful array method that do loops for you.

### `map`

The `.map` method is another very powerful tool once mastered, let's see it in
action:

```js
const time10 = [1, 2, 3, 4, 5].map((num) => {
  return `#${num}`
})

console.log(time10) // [`#1`, `#2`, `#3`, `#4`, `#5`]
```

Map takes a function and apply it to each elements of the array.

Note that map will never change the number of element of the array

For example if your function return nothing:

```js
const nothingX3 = [1, 2, 3].map((num) => {
  // Not doing anything today...
})

console.log(nothingX3) // [undefined, undefined, undefined]
```

We still get an array of 3 elements, but they are `undefined`.

You should use map everytime you want to repeat the same action for all
elements.

### Instructions

Declare a function `toSquares` that takes an array of numbers and return an
array of those squared numbers

**Example:**

```js
const result = toSquares([1, 2, 3, 4])

console.log(result) // [1, 4, 9, 16]
```
