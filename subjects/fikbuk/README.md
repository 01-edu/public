## FikBuk

### Instructions

Create a `fikBuk` function that takes a number `n` and will start counting from
`1` to `n`.

You will return an array from those values

for each numbers in between `1` and `n` you must do the following:

- if the number is a multiple of `3` add `Fik` to the array
- if the number is a multiple of `5` add `Buk` to the array
- if the number is a multiple of both `3` and `5` add `FikBuk` to the array
- otherwise add the number to the array

**Example:**

```js
fikBuk(3) // -> [1, 2, 'Fik']
fikBuk(5) // -> [1, 2, 'Fik', 4, 'Buk']
fikBuk(16) // -> [...7, 8, 'Fik', 'Buk', 11, 'Fik', 13, 14, 'FikBuk', 16]
```
