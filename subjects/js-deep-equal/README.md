## Deep Equal

### Instructions

Write a function `deepEqual(obj1, obj2)` that checks whether two objects are **deeply equal**. Two objects are considered deeply equal if:

- They have the same structure and content.
- Their nested objects and arrays are also equal.
- Primitive values (numbers, strings, booleans) are strictly equal.
- Order matters for arrays but not for object properties.

The function should return `true` if `obj1` and `obj2` are deeply equal, and `false` otherwise.

### Expected Function

```js
function deepEqual(obj1, obj2) {
  // Your implementation here
}
```

### Usage

Given the following objects:

```js
const objA = {
  name: "Alice",
  age: 25,
  hobbies: ["reading", "gaming"],
  address: {
    city: "Wonderland",
    zip: 12345,
  },
};

const objB = {
  name: "Alice",
  age: 25,
  hobbies: ["reading", "gaming"],
  address: {
    city: "Wonderland",
    zip: 12345,
  },
};

const objC = {
  name: "Alice",
  age: 25,
  hobbies: ["reading", "gaming", "chess"],
  address: {
    city: "Wonderland",
    zip: 12345,
  },
};

console.log(deepEqual(objA, objB)); // true
console.log(deepEqual(objA, objC)); // false
```

### Example Output

```sh
$ node main.js
true
false
```
