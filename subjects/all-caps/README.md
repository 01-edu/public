## ALL CAPS

### Methods

Methods are a special kind of functions, they are functions called **from**
another value.

So that's what the `.` in `console.log` is for, we call the function `log`
**from** the `console`, so it will run its actions using the console.

Now every value types have methods in JS, for example, numbers have a special
`toFixed` method that allows you to specify how many decimals you want when you
convert them to string:

```js
let thirdOf10 = 10 / 3
console.log(thirdOf10) // -> 3.3333333333333335 that's a lot of precision...
console.log(thirdOf10.toFixed(3)) // -> '3.333' that's quite enough !
console.log(thirdOf10.toFixed()) // -> with no argument, we get just '3'
```

So here in that second to last line we call `toFixed` from the value of the
variable `thirdOf10`, with the argument `3`, saying that we want 3 decimal
numbers in our string conversion.

Since they are functions, they can also return values.

### Instructions

One of the value with the most methods are strings, you can do a lot of things
with them.

For this exercise you will have to use the methods `toUpperCase` and
`toLowerCase` from the provided variable `message`.

- Create a `noCaps` variable of the value of `message` but in lower case.
- Create an `allCaps` variable of the value of `message` but in upper case.

> Just remember ALL CAPS when you spell the man name \
> ― MF DOOM
