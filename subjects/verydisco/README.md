## verydisco

### Instructions

Create a `verydisco.mjs` script that :

- takes the 1st argument from the command line (after the program name)
- makes it very disco:
  - cut each word from this argument in 2 (rounded up),
  - re-compose a word by concatenating the chunks in the other order
- display the result in console

If the argument passed is a sentence, each word of the sentence must be "very disco".

For example:

- `discovery` would print `verydisco` (🕺🏼) in console.
- `Node is awesome` would print `deNo si omeawes` in the console.

### Notions

- [Node process: `argv`](https://nodejs.org/docs/latest/api/process.html#process_process_argv)
- [nan-academy.github.io/js-training/examples/methods.js](https://github.com/nan-academy/js-training/blob/master/examples/methods.js)
- [`slice()` method](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/slice)
- [`includes()` method](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/includes)
- [`split()` method](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/split)
- [`join()` method](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/join)
- [`Math.ceil()` function](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/ceil)
