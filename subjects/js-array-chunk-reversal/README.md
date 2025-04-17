## Array Reverse Chunks

### Instructions

Write a function that takes an array `arr` and an integer `chunkSize` (where chunkSize > 0). The function should reverse the elements within each chunk of size `chunkSize` in the array and return the modified array. If the last chunk is smaller than `chunkSize`, reverse all remaining elements.

### Expected function

```js
function reverseChunks(arr, chunkSize) {}
```

### Usage

Here is a possible program to test your function:

```js
console.log(reverseChunks([1, 2, 3, 4, 5, 6, 7], 3));
console.log(reverseChunks([10, 20, 30, 40, 50], 2));
console.log(reverseChunks([1, 2, 3, 4, 5], 4));
```

And its output:

```console
$ node main.js
[3, 2, 1, 6, 5, 4, 7]
[20, 10, 40, 30, 50]
[4, 3, 2, 1, 5]
$
```
