## Pipeline

### Instructions

Write a function `pipeline(initialValue, pipeline)` that:

- Takes an initial value `initialValue` and an array of functions `functions`.
- Each function in `functions` takes the current value and returns a new value.
- After executing all pipeline functions in sequence, return an object with:
  - `finalValue`: The final transformed value.
  - `steps`: An array of objects, each describing a pipeline step:
    - `index`: The step index.
    - `input`: The input value to that step.
    - `output`: The output value from that step.
- The pipeline should be executed in a purely functional manner (no external side effects needed).

### Expected Function

```js
function pipeline(initialValue, functions) {
  // Your implementation here
}
```

### Usage

```js
const functions = [
  (val) => val + 1,
  (val) => val * 2,
  (val) => `Result: ${val}`,
];

const result = pipeline(0, functions);

console.log(result);
```

### Example Output

```sh
$ node main.js
{
  finalValue: 'Result: 2',
  steps: [
    { index: 0, input: 0, output: 1 },
    { index: 1, input: 1, output: 2 },
    { index: 2, input: 2, output: 'Result: 2' }
  ]
}
```

### Notes

- You can use `Array.prototype.reduce` or a similar approach to run through the pipeline.
- Ensure that each step’s input and output is recorded accurately.
- Keep the design purely functional as much as possible.
