## Cut Corners

### Instructions

Create some functions which behave like JavaScript's `Math` rounding functions:

- `round`: which behaves similar to `Math.round()`.
- `ceil`: which behaves similar to `Math.ceil()`.
- `floor`: which behaves similar to `Math.floor()`.
- `trunc`: which behaves similar to `Math.trunc()`.

> Some restrictions apply:
>
> - You may not use strings conversion to do it
> - No bitwise operator
> - No `%` operator

### Usage

```js
const nums = [3.7, -3.7, 3.1, -3.1]
console.log(nums.map(round))
console.log(nums.map(floor))
console.log(nums.map(trunc))
console.log(nums.map(ceil))
```

Output:

```console
[ 4, -4, 3, -3 ]
[ 3, -4, 3, -4 ]
[ 3, -3, 3, -3 ]
[ 4, -3, 4, -3 ]
```

### Notions

- [Math](https://devdocs.io/javascript/global_objects/math)

### Code provided

> The provided code will be added to your solution, and does not need to be submitted.

```js
Math.round = Math.ceil = Math.floor = Math.trunc = undefined
```
