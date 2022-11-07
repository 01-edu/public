## Reduce

### Instructions

Create four functions:

- `fold` that receives an array, a function and an accumulator, in this order,
  and applies the function in the elements of the array starting on the left.

- `foldRight` that receives an array, a function and an accumulator, in this order,
  and applies the function in the elements of the array starting on the right.

- `reduce` that works just like the method `[].reduce` when you don't
  specify an accumulator.
  The arguments should be an array and a function.
  The starting value of your accumulator must be the first value of the array.
  If your array is less than 1 argument you should throw an error.

- `reduceRight` like reduce, from the last value to the first

Example:

```js
const adder = (a, b) => a + b
fold([1, 2, 3], adder, 2) // returns 8 (2 + 1 + 2 + 3)
foldRight([1, 2, 3], adder, 2) // returns 8 (2 + 3 + 2 + 1)
reduce([1, 2, 3], adder) // returns 6 (1 + 2 + 3)
reduceRight([1, 2, 3], adder) // returns 6 (3 + 2 + 1)
```

The use of `[].reduce` and `[].reduceRight` is forbidden for this exercise.

### Notions

- [devdocs.io/javascript/global_objects/array/reduce](https://devdocs.io/javascript/global_objects/array/reduce)

### Code provided

> all code provided will be added to your solution and doesn't need to be submited.

```js
Array.prototype.reduce = undefined
Array.prototype.reduceRight = undefined
```
