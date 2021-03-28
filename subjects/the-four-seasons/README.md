## The Four Seasons

Sometimes we don't need a key, we just want a list of things, JS has a special
type for that

### Arrays

Let's see an example of an `Array`:

```js
let numberArray = [
  10, // <- no keys !
  20,
  30,
  40,
]

// or for brievety, we often write them on a single line like so:
let numberArray = [10, 20, 30, 40]
```

`Arrays` are like objects but use `[]` square bracket delimiters and only
specify values.

The `key` of an element of an array is its position, starting from `0`. We call
that its `index`

So our `numberArray` is **roughly equivalent** to writing this object:

```js
let numberObject = {
  0: 10,
  1: 20,
  2: 30,
  3: 40,
}
```

### Instructions

You must declare a variable `seasons` that contains 4 strings, one for each
seasons.

Starting with Spring, like the work of the Maestro Vivaldi.

> BGM:
> [Antonio Vivaldi - Le Quattro Stagioni](https://www.youtube.com/watch?v=b4YNYf39mcg)
