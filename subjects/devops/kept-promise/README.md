## kept-promise

### Instructions

Thing in your code does not always happen as fast as you would like.

Sometimes it is important to wait for a function to finish before keep going on with your life.

Create an asynchronous function `processInfo` in the file `kept-promise.js`. The function `processInfo` will take an asynchronous function as an input, and it will print the message `Ok!` if the return of the input function is zero or a multiple of two, `Error!` otherwise.

The following function can be used as an example of input function for `processInfo`:

```js
const getImportantInfo = _ =>
  new Promise(resolve => {
    setTimeout(() => resolve(Math.round(Math.random() * 10)), 1000)
  })
```

> Assume that your function will always get a valid input function 

### Example

The following script `main.js` can be used to test your function:

```js
import processInfo from './kept-promise.js'

const getImportantInfo = _ =>
  new Promise(resolve => {
    setTimeout(() => resolve(Math.round(Math.random() * 10)), 1000)
  })

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

### Notions

- [Using promises](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Using_promises)
- [async function](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/async_function)
