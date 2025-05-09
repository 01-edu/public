## Final Attempt

### Instructions

Write a function `FinalAttempt(callback, count)` that:

- Takes two parameters:
  - `callback`: An asynchronous function to be invoked.
  - `count`: A positive integer representing the maximum number of retries.
- Returns a function that:
  - Accepts any number of arguments. These arguments are passed to the `callback` when invoked.
  - Attempts to call the `callback` with the provided arguments. If the `callback` throws an error or rejects, retry the call until the `count` limit is reached.
  - If all attempts fail, it should return `"Final Attempt Fail"`.
  - If the `callback` succeeds within the retry limit, return the result of the `callback`.

### Expected Function

```js
function FinalAttempt(callback, count) {
  // Your implementation here
}
```

### Usage

```js
const unreliableAsyncFunction = async (param) => {
  if (Math.random() > 0.7) {
    return `Success with ${param}`;
  }
  throw new Error("Failure");
};

const finalAttempt = FinalAttempt(unreliableAsyncFunction, 3);

finalAttempt("test")
  .then((result) => console.log(result))
  .catch((err) => console.error(err));
```

### Example Output

```sh
$ node main.js
Success with test

$ node main.js
Final Attempt Fail
```
