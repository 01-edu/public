## Transform Keys

### Instructions

Write a function `transformKeys(obj, transformFn)` that recursively applies a transformation function to all keys of an object.

The function should:

- Apply `transformFn` to every key in the object.
- Recursively apply the transformation to all nested objects.
- Preserve all values without modification.

### Expected Function

```js
function transformKeys(obj, transformFn) {
  // Your implementation here
}
```

### Usage

Given the following object:

```js
const obj = {
  FirstName: "Alice",
  LastName: "Doe",
  Address: {
    StreetName: "Main St",
    ZipCode: 12345,
  },
};

const toSnakeCase = (key) =>
  key
    .replace(/[A-Z]/g, (letter) => `_${letter.toLowerCase()}`)
    .replace(/^_/, "");
const transformed = transformKeys(obj, toSnakeCase);
console.log(transformed);
```

### Example Output

```sh
{
  "first_name": "Alice",
  "last_name": "Doe",
  "address": {
    "street_name": "Main St",
    "zip_code": 12345
  }
}
```
