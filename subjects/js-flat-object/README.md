## Flat Object

### Instructions

Write a function `flattenAndMap(obj, mapper)` that:

- Takes a deeply nested object `obj` where values might be objects or primitives.
- Flattens this object so that each key in the resulting object represents the path to the value in the original object, using dot notation for nesting (e.g. `user.address.city`).
- Applies a mapper function to each value (primitive) before placing it into the flattened result.
- Returns an object with:
  - `flattened`: The flattened object after mapping values.
  - `originalKeysCount`: The total number of keys (including nested) found in the original object.
  - `transformedKeysCount`: The total number of keys whose values were transformed by the mapper.

### Expected Function

```js
function flattenAndMap(obj, mapper) {
  // Your implementation here
}
```

### Usage

```js
const nestedObj = {
  user: {
    name: "Alice",
    address: {
      city: "Wonderland",
      zip: 12345,
    },
  },
  meta: {
    roles: ["admin", "editor"],
    active: true,
  },
};

const toUpperCaseStrings = (value) =>
  typeof value === "string" ? value.toUpperCase() : value;

const result = flattenAndMap(nestedObj, toUpperCaseStrings);

console.log(result);
```

### Example Output

```sh
$ node main.js
{
  flattened: {
    'user.name': 'ALICE',
    'user.address.city': 'WONDERLAND',
    'user.address.zip': 12345,
    'meta.roles': ['ADMIN', 'EDITOR'],
    'meta.active': true
  },
  originalKeysCount: 5,
  transformedKeysCount: 5
}
```

### Notes

- The mapper function should be applied to every value, regardless of type. If it's not a string or the mapper doesn't change it, that's fine.
- Arrays are considered values too, and should be mapped. If mapping arrays, apply the mapper to each element.
- Count all final flattened paths as keys in the resulting object.
- `originalKeysCount` is how many ultimate primitive or array-containing keys were found in the original nested structure.
- `transformedKeysCount` is how many values ended up being transformed by the mapper (even if that transformation doesn't change the value).
