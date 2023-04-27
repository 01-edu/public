## kept-promise

### Instructions

Things in your code do not always happen as fast as you would like.

Sometimes it is important to wait for a function to finish before keep going on. Other times you don't want your code to be blocking while waiting.

Create an asynchronous function `processInfo` in a file `kept-promise.js`. The function `processInfo` will take an asynchronous function as an input, and it will print the message `Ok!` if the return of the input function is zero or a multiple of two, `Error!` otherwise.

The following function can be used as an example of input function for `processInfo`:

```js
const getImportantInfo = async () =>
  new Promise(resolve => resolve(Math.round(Math.random() * 10)))
```

> Assume that your function will always get a valid input function

### Example

The following script `main.js` can be used to test your function:

```js
import processInfo from './kept-promise.js'

const getImportantInfo = () =>
  new Promise(resolve => resolve(Math.round(Math.random() * 10)))

console.log(await processInfo(getImportantInfo))
```

The output should be the following:

```console
$  node main.js
Ok!
$  node main.js
Error!
$  node main.js
Ok!
$
```

### Hints

- Asynchronous behavior, in the context of server-side JavaScript, refers to the ability of running code in a non-blocking way, meaning for instance that the server can handle multiple requests simultaneously without waiting for a long-running task to complete, avoiding blocking the server responsiveness. It is essential to achieve scalable and functional back-end Node applications.

- A `Promise` is a special JS object that represent the result (success or failure) of an asynchronous operation. This special object is usually used to "wrap" asynchronous operations.
  `Promise`s can have three different states: _pending_ - the asynchronous operation has not finished yet, _fulfilled_/_resolved_ - the asynchronous operation has finished successfully - or _rejected_ - the asynchronous operation has finished, but something went wrong. When defined from scratch, it is possible to define a `resolve` and `reject` callback function for a new `Promise` that will define the results of the success or failure of asynchronous operation happening inside the `Promise`.

- It is possible to wait for an asynchronous function with the keyword `await`. Alternatively, it can be clearer to use the method `.then`. Below an example of how to handle promises and, more generally, asynchronous operations.

```js
const promise1 = new Promise((resolve, reject) => {
  resolve('Success!')
})

console.log(await promise1)
// Expected output: "Success!"

promise1.then(value => {
  console.log(value)
  // Expected output: "Success!"
})
```

### Notions

- [Introducing asynchronous JavaScript](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Asynchronous/Introducing)
- [await](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/await)
- [Using promises](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Using_promises)
- [async function](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/async_function)
- [then](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/then)
