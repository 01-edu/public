## Nested Array Reverser

### Instructions

Write a function `nestedArrayReverser(words)` that takes a 2D array of `words`, reverses the order of the inner arrays, reverses the order of words within each inner array, and returns a single string with all the words joined by a space.

### Expected function

```javascript
nestedArrayReverser(words);
```

### Usage

Here is a possible program to test your function:

```javascript
console.log(
  nestedArrayReverser([
    ["hello", "world"],
    ["this", "is"],
    ["a", "test"],
  ]),
  console.log(nestedArrayReverser([])),
);
```

And its output:

```console
$ node main.js
test a is this world hello
$
$
```
