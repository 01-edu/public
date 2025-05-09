## Grid Words Finder 2

### Instructions

Write a function `gridWordFinder2(grid, word)` that:

- Takes a 2D grid of characters and a target word as input.
- Searches for all occurrences of the word in the grid.
- For each occurrence, return an object containing:
  - x and y: The starting coordinates of the word.
  - direction: Either "horizontal" or "vertical" depending on the word's orientation.

> ⚠️ Note: You are not allowed to use any kind of loops

### Expected function

```js
function gridWordFinder2(grid, word) {}
```

### Usage

Here is a possible program to test your function:

```js
const grid = [
  ["c", "a", "t"],
  ["d", "o", "g"],
  ["r", "a", "t"],
];

console.log(gridWordFinder2(grid, "cat"));
```

And its output:

```console
$ node main.js
[{x:0, y:0, direction:"horizontal"}]
$
```
