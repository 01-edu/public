## Deep Freeze

### Instructions

Write a function `deepFreeze(obj)` that makes an object **completely immutable**, preventing any modifications at all levels.

The function should:

- Use `Object.freeze()` to prevent modifications to the object itself.
- Recursively apply `Object.freeze()` to all nested objects and arrays.
- Ensure that properties cannot be **added**, **deleted**, or **modified** at any depth.

Once frozen, any attempt to modify the object should **fail silently in non-strict mode** and **throw an error in strict mode**.

### Expected Function

```js
function deepFreeze(obj) {
  // Your implementation here
}
```

### Usage

Given the following object:

```js
const data = {
  user: "Alice",
  stats: {
    score: 100,
    achievements: ["gold", "silver"],
  },
};

deepFreeze(data);

data.user = "Bob"; // This should fail
data.stats.score = 200; // This should fail
data.stats.achievements.push("bronze"); // This should fail

console.log(data);
```

### Example Output

```sh
$ node main.js
{
  user: "Alice",
  stats: {
    score: 100,
    achievements: ["gold", "silver"]
  }
}
```
