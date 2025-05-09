## Flatten Object

### Instructions

Write a function `flattenObject(obj)` that flattens a nested object into a single-level object where keys represent the nested path using dot notation (`.`).

The function should:

- Convert all nested properties into a single-level object.
- Use dot notation to represent the path of nested keys.
- Maintain array indices in the key path when flattening arrays.

### Expected Function

```js
function flattenObject(obj) {
  // Your implementation here
}
```

### Usage

Given the following object:

```js
const obj = {
  a: 1,
  b: { c: 2, d: { e: 3 } },
  f: [4, 5, { g: 6 }],
};

const flattened = flattenObject(obj);
console.log(flattened);
```

### Example Output

```sh
{
  "a": 1,
  "b.c": 2,
  "b.d.e": 3,
  "f.0": 4,
  "f.1": 5,
  "f.2.g": 6
}
```
