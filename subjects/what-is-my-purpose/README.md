## What is my purpose ?

![robot](https://cdn.discordapp.com/attachments/489466992286498816/828181029772197888/butter-purpose.png)

You have seen how to call functions that were stored in object properties.

Remember this example of function call ?

```js
//       ↙ identifier, like variables
console.log('Hello There !') // <- function call happening here
//          ↖ open paren + argument + close paren
```

or these ones:

```js
let roundedValue = Math.round(1.8) // another function call happening here
console.log(roundedValue) // and another one here
```

There, we saw how to call and use _"built-in"_ functions.

Here, now, we are going to learn how to declare our owns. This will gives us
even more freedom to build our own logic.

### Declaring a function

We mentioned before that a function had to be either assigned to a variable or
an object property.

We will now declare one in a variable. Let's call it: `myFirstFunction`.

Once a variable is declared; Remember how an array can be recognized with this
syntax : `[]` ? or an object with this one : `{}` ?

Well, we will use another syntax for declaring our function. We will do so using
the `ArrowFunctionExpression` syntax : `() => {}`

- So we first put parens `()`, These, are the containers of the `arguments` that
  go in the function. For now we will leave them empty with no arguments. (More
  on those later on)
- We then add the arrow `=>` which is the distinguishing feature of the
  `ArrowFunctionExpression` syntax.
- Finally, we add the curly brackets `{}` to delimite the scope of our newly
  created function. Note: They are not always necessary, you will probably find
  examples of this function syntax without the `{}`. However, for now because
  you are learning. We will put them most of the time.

```js
//    ↙ normal variable     ↙ beginning of the scope of the function
let myFirstFunction = () => {
  //                    ↖ parens () for arguments and the arrow => for syntax
} // <-end of the scope of the function
```

It's now possible to call this function using the `()`, like any pre-declared
function:

```js
myFirstFunction() // nothing happens
```

This function if called, does not do anything, since it doesn't contain any
code.

### The scope of a function `{}`

Very much like an `if` statement a function has a scope. The scope in between
the curly braces`{}` is where the action happens. Let's add something to the
scope of our function.

```js
let myFirstFunction = () => {
  console.log('Calling my first function')
  // ↖ some instructions to do when the function is called !
}
```

Now the function if called, display the output the `console.log()`.

```js
myFirstFunction() // "Calling my first function"
```

We actually used a function and gave it this single instruction:

- call another function `console.log()`.

### Instructions

You are a robot made by a scientist called Rick and you want to know your
purpose.

- Declare a function named `ask` that log `'What is my purpose ?'` in the
  console
- Declare a function named `reply` that log `'You pass butter.'` in the console

Then first call the `ask` then the `reply` once, in that order.
