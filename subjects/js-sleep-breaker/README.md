## Sleep Breaker

### Instructions

Write a function `sleepBreaker(delay, breaker)` that:

- Takes two parameters:
  - `delay`: A positive integer representing the delay time in milliseconds.
  - `breaker`: An asynchronous function that will determine if the sleep should be interrupted.
- Returns a `Promise` that:
  - Resolves after the given `delay` in milliseconds.
  - If the `breaker` function resolves before the delay is completed, the `Promise` should resolve immediately, interrupting the sleep.

### Expected Function

```js
function sleepBreaker(delay, breaker) {
  // Your implementation here
}
```

### Usage

```js
const breaker = async () => {
  await new Promise((res) => setTimeout(res, 100));
  return true;
};

sleepBreaker(1000, breaker).then(() => {
  console.log("Wake up!");
});
```

### Example Output

```sh
$ node main.js
Wake up! // After ~100ms because the breaker resolved
```
