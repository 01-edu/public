## No Adults Wanted

### `filter`

In JS most of the time, a loop is not needed and specials array methods can help
doing the heavy work.

The first one we are going to see is the `filter` array method:

```js
let values = [1, 23, 3, 43, 78, 9, 23, 12]

let isOver20 = (value) => {
  if (value > 20) {
    return true
  }
}

let valuesUnder10 = values.filter(isOver20)

console.log(valuesUnder10) // [23, 43, 78, 23]
```

In this case, `values` is an array and the `.filter` method is used to filter
only those that are over `20`.

This is achieved by giving `filter` a special function that will return `true`
if we want to keep this value or not.

Filtering is the best way to remove elements from an array, when you only need
to do a selection but not modify them.

### Instructions

Create a function `childrenOnly` that takes an array of numbers and returns only
those that are lower than 18

**Example:**

```js
const result = childrenOnly([12, 11, 23, 44, 10])

console.log(result) // [12, 11, 10]
```
