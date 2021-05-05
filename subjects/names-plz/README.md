## Names plz

### Chaining methods

One common pattern in javascript is to chain methods to achieve a specific goal:

```js
let title = 'Title:   THE MATRIX  '
  .slice(6) //          '   THE MATRIX  '
  .trim() //            'THE MATRIX'
  .toLowerCase() //     'the matrix'
  .replace(' ', '-') // 'the-matrix'

// Final result:
console.log(title) // 'the-matrix'
```

This is an example of a chaining 4 different methods to transform the original
string `'Title: THE MATRIX '` into `'the-matrix'`.

### Instructions

It's now your turn to chain methods!

Declare a `namePlz` function that takes an array of objects with a `.name`
string property and does:

- `.map` form objects to get only the property `name`
- `.join` each names with a `,` and a space

**Example:**

```js
const result = namePlz([
  { name: 'Wasan' },
  { name: 'Alanoud' },
  { name: 'Salama' },
  { name: 'Yousuf' },
  { name: 'Maitha' },
])

console.log(result) // 'Wasan, Alanoud, Salama, Yousuf, Maitha'
```
