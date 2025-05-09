## Bubble Sort Analyzer

### Instructions

Write a function `bubbleSortAnalyzer(arr, comparator)` that:

- Takes an array `arr` and a comparator function `comparator`.
- Sorts the array using the **bubble sort** algorithm and returns an object with the following properties:
  - `sortedArray`: The sorted array.
  - `iterations`: The total number of iterations (steps) performed during the sort.
  - `swaps`: The total number of swaps performed during the sort.

### Expected function

```js
function bubbleSortAnalyzer(arr, comparator) {}
```

### Usage

Here is a possible program to test your function:

```js
const comparator = (a, b) => a - b;

const result = bubbleSortAnalyzer([4, 2, 7, 1, 3], comparator);

console.log(result);
```

And its output:

```console
$ node main.js
{
  sortedArray: [1, 2, 3, 4, 7],
  iterations: 10,
  swaps: 7
}
$
```

### Notes

- The `comparator` function determines the sorting order:
  - If `comparator(a, b)` returns a value < 0, `a` is considered smaller than `b`.
  - If it returns 0, `a` and `b` are considered equal.
  - If it returns a value > 0, `a` is considered greater than `b`.
- You should implement a **step-by-step bubble sort**, counting both the number of iterations through the array and the number of swaps performed.
