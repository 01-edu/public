## Exam Grader

### Instructions

Write a function `examGrader(timeout, exercises)` that:

- Takes two parameters:
  - `timeout`: A positive integer representing the maximum allowed time for the entire exam in milliseconds.
  - `exercises`: An array of Promises, where each Promise represents an exercise in the exam. Each Promise resolves with an object containing:
    - `note`: A number from `0` to `20` representing the grade for that exercise.
- Executes the exercises in the same order they are provided.
- Calculates the **final grade** by:
  - Summing the grades (`note`) of all exercises that completed execution within the `timeout`.
  - If the total time of the exercises exceeds the `timeout`, the function should immediately stop processing further exercises.
- Returns the **final grade** as a number.

### Expected Function

```js
function examGrader(timeout, exercises) {
  // Your implementation here
}
```

### Usage

```js
const exercises = [
  () =>
    new Promise((res) => setTimeout(() => res({ time: 1000, note: 15 }), 1000)),
  () =>
    new Promise((res) => setTimeout(() => res({ time: 2000, note: 18 }), 2000)),
  () =>
    new Promise((res) => setTimeout(() => res({ time: 500, note: 12 }), 500)),
];

examGrader(3000, exercises).then((finalGrade) => {
  console.log(`Final Grade: ${finalGrade}`);
});

examGrader(4000, exercises).then((finalGrade) => {
  console.log(`Final Grade: ${finalGrade}`);
});
```

### Example Output

```sh
$ node main.js
Final Grade: 15
Final Grade: 45
```
