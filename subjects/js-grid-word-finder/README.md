## Grid Words Finder

### Instructions

Write a function gridWordsFinder(grid, words) that takes a 2D grid of characters and an array of words.
Finds and returns all words from the array that can be traced horizontally (left to right) or vertically (top to bottom) in the grid.

### Expected function

```js
function gridWordsFinder(grid, words) {}
```

### Usage

Here is a possible program to test your function:

```js
const grid = [
  ["c", "a", "t"],
  ["a", "a", "t"],
  ["r", "a", "t"],
  ["d", "o", "g"],
];

console.log(
  gridWordsFinder(grid, ["cat", "dog", "rat", "tar", "car", "rac", "g"]),
);
console.log(gridWordsFinder(grid, []));
```

And its output:

```console
$ node main.js
["cat", "dog", "rat", "car", "g"]
[]
$
```
