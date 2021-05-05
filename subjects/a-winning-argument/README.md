## A winning argument

### Arguments

We mentioned it before with methods, functions can take arguments. They are
always in between parens `()`.

Let's use the same examples that we used for function calls:

Remember this example of function call ?

```js
//       ↙ method
console.log('Hello There !') //<-
//                 ↖ The String 'Hello There!' is
//                   the argument of console.log()
```

or these ones ?

```js
let roundedValue = Math.round(1.8) // The number 1.8 is the arg
console.log(roundedValue) // the variable roundedValue is the arg
```

We are now going to adapt `myFirstFuntion` so that it takes one argument :
`arg1`.

```js
let myFirstFunction = (arg1) => {
  //<-arg1 is inputed in between the parens
  console.log(arg1) // arg1 can be use inside the scope of the function
  //            ↖   arg1 is "transfered" to be the arg of console.log()
} //<-end of the scope of the function
```

Now the function if called, display the output the `console.log(arg1)`.

```js
myFirstFunction('using my first arg') // "using my first arg"
```

But let's say we want to change what the function logs. Now, instead of
modifying `myFirstFunction` we just need to modify the argument in the function
call.

```js
myFirstFunction('another arg') // "another arg"
myFirstFunction('and another one') // "and another one"
myFirstFunction('and one more') // "and one more"
```

> Waste no more time arguing about what a good person should be. Be one.
>
> - Marcus Aurelius

### Instructions

You are a general's aide who has to transmit the communications to the other
soldiers.

In order to do so you will create the function `battleCry`. This function will
take one argument and will display it in the console.

The battlefield is big so make sure that the argument is uppercased before
displaying it.

Now, sometimes the communications will have to given quietly.

For this you will create the function `secretOrders` which does the same as
`battleCry` except that it lowercases the argument before sending it.
