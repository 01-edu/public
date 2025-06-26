## Deep Find

### Instructions

Write a function `deepFind(obj, path)` that retrieves a deeply nested value from an object. The `path` is provided as a string using **dot notation**.

The function should:

- Traverse the object following the provided path.
- Return `undefined` if any part of the path is missing.
- Support accessing nested objects and arrays.

### Expected Function

```js
function deepFind(obj, path) {
  // Your implementation here
}
```

### Usage

```js
const obj = {
  user: {
    profile: {
      name: "Alice",
      details: {
        age: 25,
        hobbies: ["reading", "gaming"],
      },
    },
  },
};

console.log(deepFind(obj, "user.profile.name")); // "Alice"
console.log(deepFind(obj, "user.profile.details.age")); // 25
console.log(deepFind(obj, "user.profile.details.hobbies.1")); // "gaming"
console.log(deepFind(obj, "user.profile.address.city")); // undefined
```

### Example Output

```sh
$ node main.js
Alice
25
gaming
undefined
```
