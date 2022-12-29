## level up

You know now, in your function declarations, how to use arguments and returns.

Let's try to now use the `if` statements in a new function called `happy`. Here
is an example:

```js
let happy = (mood) => {
  if (mood === 'happy') {
    return true
  }
  return false
}

let result1 = happy('happy')
let result2 = happy('sad')

console.log(result1) // true
console.log(result2) // false
```

Here we used the `if` statement, and two `return` keywords to alternate between
the result `true` or `false` depending whether the argument `mood` is equal to
the string `happy` or not. The possibilities are becoming limitless...

### Instructions

As Rick's robot, you are continuing your training to add yourself new ... skills
(I could have said funtions). You want now to become a robot barista.

Define the function `shaker` which will take as arguments:

- `quantity`, which will be variable of type `Number`
- `fruit`, which will be a `String`
- `diet`, which will be a `Boolean`

`shaker` must return a `String`. Look at the examples below to understand how
`shaker` must mix its ingredients:

```js
console.log(shaker(1, 'strawberry', true))
//'1 skinny strawberry milkshake'
console.log(shaker(2, 'chocolate', false))
//'2 chocolate milkshakes'
console.log(shaker(2, 'strawberry', true))
//'2 skinny strawberry milkshakes'
console.log(shaker(1, 'chocolate', false))
//'1 chocolate milkshake'
```

Ps: watch out for your plurals!
