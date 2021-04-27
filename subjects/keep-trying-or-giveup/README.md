## Keep Trying Or Giveup

### Instructions

Create a `retry` function, that takes 2 arguments

- a `count` indicates maximum amount of retries
- an async `callback`, that will be called on every try

`retry` returns a function that calls and returns value from `callback`
function passing its arguments and catches errors. If error is caught it
should return the `callback` function with catch. If number of errors
exceeds `count` then throw an `Error`.

> for count of 3, the function will be called at most 4 times:
> the initial call + 3 retries.

Create a `timeout` function, that takes 2 arguments

- a `delay` indicates maximum wait time
- an async `callback`, that will be called

`timeout` returns a function either that calls and returns value from `callback`
function passing its arguments or returns `Error('timeout')` if `callback` didn't
resolve before `delay` time has reached.

### Notions

- [nan-academy.github.io/js-training/examples/promise.js](https://nan-academy.github.io/js-training/examples/promises.js)
- [devdocs.io/dom/windoworworkerglobalscope/settimeout](https://devdocs.io/dom/windoworworkerglobalscope/settimeout)
- [devdocs.io/javascript/global_objects/promise/race](https://devdocs.io/javascript/global_objects/promise/race)
