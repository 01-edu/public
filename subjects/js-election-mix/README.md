## Election Mix

### Instructions

Write a function `createCurriedFilterAndMap(criteria, mapper)` that:

- Accepts a `criteria` function and a `mapper` function.
  - `criteria(key, value)` returns `true` if the key-value pair should be included, `false` otherwise.
  - `mapper(value)` transforms the value of each included key.
- Returns a **curried function** that:
  - Takes an object as input.
  - Returns a new object containing only the key-value pairs that passed the criteria, with their values transformed by the mapper.
- Alongside the filtered and transformed object, return:
  - `keysKept`: The number of keys that passed the criteria.
  - `keysFilteredOut`: The number of keys that were filtered out.

### Expected Function

```js
function createCurriedFilterAndMap(criteria, mapper) {
  // Your implementation here
}
```

### Usage

```js
const isNumberValue = (key, value) => typeof value === "number";
const increment = (val) => (typeof val === "number" ? val + 1 : val);

const filterAndIncrementNumbers = createCurriedFilterAndMap(
  isNumberValue,
  increment,
);

const inputObject = {
  name: "Alice",
  age: 25,
  score: 99,
  active: true,
};

const result = filterAndIncrementNumbers(inputObject);

console.log(result);
```

### Example Output

```sh
$ node main.js
{
  filteredObject: { age: 26, score: 100 },
  keysKept: 2,
  keysFilteredOut: 2
}
```

### Notes

- Ensure your function returns a partially applied (curried) function. That means `createCurriedFilterAndMap(criteria, mapper)` returns a function that takes the object separately.
- `criteria` receives `(key, value)` and returns a boolean.
- `mapper` receives `(value)` and returns a transformed value.
- Count `keysKept` and `keysFilteredOut` based on the object provided to the curried function.
