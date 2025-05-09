## Trap Object

### Instructions

Write a function `trapObject(obj, fn)` that takes an object and a function. The returned object should behave normally but **run `fn` every time a property is accessed or modified**.

The function `fn` should receive the following arguments:

- `"get"` when a property is accessed, along with the key and value.
- `"set"` when a property is modified, along with the key, old value, and new value.

### Expected Function

```js
function trapObject(obj, fn) {
  // Your implementation here
}
```

### Usage

```js
const logger = (action, key, value, newValue) => {
  if (action === "get") {
    console.log(`Accessed ${key}: ${value}`);
  } else if (action === "set") {
    console.log(`Modified ${key}: ${value} → ${newValue}`);
  }
};

const obj = trapObject({ name: "Alice", age: 25 }, logger);

console.log(obj.name); // Logs: Accessed name: Alice
obj.age = 30; // Logs: Modified age: 25 → 30
console.log(obj.age); // Logs: Accessed age: 30
```

### Example Output

```sh
$ node main.js
Accessed name: Alice
Modified age: 25 → 30
Accessed age: 30
```
