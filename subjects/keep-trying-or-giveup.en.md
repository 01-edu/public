## Keep Trying Or Giveup

### Instructions

Create a `retry` function, that takes 2 arguments
- a `count`, that tells how many retries must be done
- an async `callback`, that will be call every try

and it return a new function, passing arguments given to the
callback on every tries.

> for count of 3, the function will be called at most 4 times:
> the initial call + 3 retries.


Create a `timeout` function, that takes 2 arguments
- a `delay`, that tells how long to wait
- an async `callback`, that will be call

and it return a new function, passing arguments given to the callback
and either the async callback resolve before the delay is reached,
in that case we return the value from the callback,
or reject an error using the message `"timeout"`


### Notions

- https://nan-academy.github.io/js-training/examples/promise.js
- https://devdocs.io/dom/windoworworkerglobalscope/settimeout
- https://devdocs.io/javascript/global_objects/promise/race