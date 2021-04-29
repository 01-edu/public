## Two can play that game

### More Arguments

We have seen how to add one argument to the a function, We are now going to see
how to add two (or more).

All we need to do to add a second argument `arg2` is to add a comma `,` after
`arg1` and then add `arg2`.

```js
let myFirstFunction = (arg1, arg2) => {
  //<-arg1 and arg2 are inputed in between the parens
  console.log(arg1, arg2)
  //            ↖   arg1 and arg2 are "transfered" to be the args of console.log()
}
// Now we call the function
myFirstFunction('first arg', 'second arg')
// "first arg"
// "second arg"
```

For more args, you will need to simply repeat the same process. Example with
`arg3` and `arg4`:

```js
let myFirstFunction = (arg1, arg2, arg3, arg4) => {
  //<as many arguments as needed
  console.log(arg1, arg2, arg3, arg4)
} //<-end of the scope of the function
// Now we call the function
myFirstFunction('first arg', 'second arg', 3, 4)
// "first arg"
// "second arg"
// 3
// 4
```

Finally, please note that you can name your arguments however you please. Just
make sure that you reuse the proper name inside the scope of your function.
Example:

```js
let myFirstFunction = (continent, country, city, temperature) => {
  //<as many arguments as needed
  console.log(continent, country, city, temperature)
} //<-end of the scope of the function
// Now we call the function
myFirstFunction('Europe', 'France', 'Paris', '30°C')
// "Europe"
// "France"
// "Paris"
// "30°C"
```

### Instructions

You want to conquer the world! _(of entertainment)_.

In this exercise you will have to declare the function `duos`.

This function will take two arguments and will log them together with an `and`
and an `!`

> Example: calling: `duos('Batman', 'Robin')` should log : `Batman and Robin !`

You will then declare the function `duosWork`. This function will take three
arguments and will log them togethers as in the example below.

> Example: calling: `duosWork('Batman', 'Robin', 'protect Gotham')` should log :
> `Batman and Robin protect Gotham !`
