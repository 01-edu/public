## Currify

### Instructions

Create the function `currify` that will curry any functions put as argument.

example:

```js
const mult2 = (el1, el2) => el1 * el2
console.log(mult2(2, 2)) // result expected 4

const mult2Curried = currify(mult2)

console.log(mult2Curried(2)(2)) // result expected 4
// (same result, with a function that has technically only one argument)
```

### Notions

- [stackoverflow.com/questions/36314/what-is-currying](https://stackoverflow.com/questions/36314/what-is-currying)
