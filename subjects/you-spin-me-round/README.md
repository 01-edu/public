## You Spin Me Round

### Functions

Functions in JS allows you to describe multiple `instructions`, in other words,
it's a way to execute code inside your code !

That seems pretty complicated but it is the building block of your programs.

You already have been using a `function` as `console.log` is one ! You can easly
spot them because we stick parens after their identifiers `()`.

For example, here's a **function call**:

```js
//       ↙ identifier, like variables
console.log('Hello There !')
//          ↖ open paren + argument + close paren
```

#### Function name (`identifier` or `property.key`)

The first things that appears in a function is the `identifier`, in fact, in JS,
functions are values of type `function`. This means that in order to be used a
function **must** be either assigned to:

- a variable
- or an object property

#### Function calling

Functions **do** something, and they can **return a result**, a value.

`console.log` does not return a value, but it will do something: make its
arguments appear in the console.

To `call` the function, in other word run it, we need too add `()`.

```js
console.log // function is not used, nothing happen
console.log() // function was called ! an empty line appear in the console
console.log(4) // function was called with 4 and it appears in the console
```

#### Function `arguments`

So in that last example, number `4` was the argument of the `console.log`
function.

A function will execute the same code on different arguments, making them
flexible.

> Sometimes, `arguments` are named `parameters`. We just like to use 10
> different names for everything to sound _"smart"_.

JS gives you plenty of readymade functions, for now we are going to focus on
`Math` functions.

#### function `return values`

All functions from the JS `Math` object do nothing other than compute a new
value from its argument.

For example, the well named `Math.round` function will take a number as argument
and returns the rounded value of this number.

To use return values, assign them to `variables`:

```js
let roundedValue = Math.round(1.8) // Here we assign the result of the function call
console.log(roundedValue) // 2
```

The variable `roundedValue` value is the number `2`, the result of the function
call.

#### Nested function calls

It is also possible to use the return value of a function directly without using
an intermediary variable.\
For example we could have written:

```js
console.log(Math.round(3.2)) // double functions call !!!! woaaaa
```

Here we first call `Math.round(3.2)` which returns the number `3` and that will
be passed to `console.log` that will procede to display it.

### Instructions

We have prepared a variable `num`. Just as a warm up, use this variable `num` as
`argument` of some `Math` functions.

- Declare a `rounded` variable of `num` rounded value.
- Declare a `truncated` variable of `num` truncated value.

One of the necessary Math function is already used in the lessons examples
before the instructions. Explore the link below to see which others functions
the Math object contains in order to find what you need to complete this
exercise.

### Notions

- [Math](https://devdocs.io/javascript/global_objects/math)
