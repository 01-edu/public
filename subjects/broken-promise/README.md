## kept-promise

### Instructions

Sometimes, even with the best of intentions, some promises are not kept.

Imagine what should happen if a request is made to some API, but the server is offline.

A `Promise` has two possible outcomes. A `Promise` will eventually "resolve" to some value, or "reject" with some error information.

For your Javascript code to properly function, you'll need to write your code to handle both `Promise` eventualities (resolve and reject).

Create an asynchronous function `doubleIt` in a file `broken-promise.js`. The function `doubleIt` will take an asynchronous function as an input. The asynchronous function passed as an argument will return a `Promise` which will resolve to a number.

If the `Promise` resolves, your `doubleIt` function should double the value before returning it.

If the `Promise` rejects, your `doubleIt` function should prepend "`Error: `" to the error message before returning it.

> Assume that your function will always get a valid input function

### Example

The following `main.js` can be used to test your function:

```js
import doubleIt from './broken-promise.js'

const asyncFive = async () => new Promise((resolve) => resolve(5))

const asyncSeven = async () => new Promise((resolve) => resolve(7))

const asyncReject = async () => new Promise((_, reject) => reject("There are no numbers available"))

console.log(await doubleIt(asyncFive))
console.log(await doubleIt(asyncSeven))
console.log(await doubleIt(asyncReject))
```

The output should be the following:

```console
$ node main.js
10
14
Error: There are no numbers available
$
```

### Hints

You may be wondering why it is important to handle resolve and reject cases?

Imagine if you expect a `Promise` to resolve to a `number`. If the `Promise` rejects, and you try to interpret the value as a `number`, this could cause your program to crash.

In true Javascript fashion, "there are many ways to skin a cat". It is possible to deal with resolved and rejected promises in more than one way.

There is the `async`/`await` way, which makes use of `try`/`catch`. This functionality is only available inside `async` functions. This method allows you to write "synchronous" looking code, which is blocking while the `Promise` is waiting to resolve.

There is the `then`/`catch` way, which is non-blocking. The rest of the code executes while the `Promise` is waiting to resolve.

Say we have some function, which waits a while before returning a value.
```js
const someFunc = async (delay) => new Promise((resolve, reject) => {
  if (delay > 0) {
    setTimeout(() => {
      resolve(delay)
    }, delay * 1000)
  } else {
    reject("Delay must be positive")
  }
})
```

If we invoke `someFunc` in the `async`/`await` way:
```js
const asyncAwait = async () => {
  try {
    const value = await someFunc(2)
    // Do something with the value once the Promise resolves
    console.log(value)
  } catch (error) {
    // Do something with the error if the Promise is rejected
    console.log(error)
  }
  console.log("End asyncAwait")
}
```

It will output like this, because the last `console.log` is blocked until the `Promise` is resolved. This is because the `await` keyword is blocking:

```console
2
End asyncAwait
```

If we invoke `someFunc` in the `then`/`catch` way:
```js
const thenCatch = () => {
  someFunc(4)
    .then((value) => {
      // Do something with the value once the Promise resolves
      console.log(value)
    })
    .catch((error) => {
      // Do something with the error if the Promise is rejected
      console.log(error)
    })
  console.log("End thenCatch")
}
```

It will output like this, because the last `console.log` is not blocked. The code continues to execute until the `Promise` is resolved:
```console
End thenCatch
4
```

### Notions

- [Introducing asynchronous JavaScript](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Asynchronous/Introducing)
- [await](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/await)
- [Using promises](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Using_promises)
- [async function](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/async_function)
- [then](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/then)
