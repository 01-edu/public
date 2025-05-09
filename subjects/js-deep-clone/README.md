## Deep Clone

### Instructions

Write a function `deepClone(obj)` that takes an object and returns a **deep clone** of it. The function should recursively copy all nested objects and arrays, ensuring that modifying the original object does not affect the cloned object.

The cloning process should:

- Copy **all primitive values** (e.g., numbers, strings, booleans).
- Recursively copy **nested objects**.
- Recursively copy **arrays** while maintaining element order.
- Ensure that objects with circular references are handled safely.

### Expected Function

```js
function deepClone(obj) {
  // Your implementation here
}
```

### Usage

Given the following object:

```js
const original = {
  name: "Alice",
  age: 25,
  hobbies: ["reading", "gaming"],
  address: {
    city: "Wonderland",
    zip: 12345,
  },
};

const cloned = deepClone(original);

cloned.name = "Bob";
cloned.hobbies.push("chess");
cloned.address.city = "Nowhere";

console.log(original);
console.log(cloned);
```

### Example Output

```sh
Original:
{
  name: "Alice",
  age: 25,
  hobbies: ["reading", "gaming"],
  address: { city: "Wonderland", zip: 12345 }
}

Cloned:
{
  name: "Bob",
  age: 25,
  hobbies: ["reading", "gaming", "chess"],
  address: { city: "Nowhere", zip: 12345 }
}
```
