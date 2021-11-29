## You pass butter

### Return values

We are now going to see how to declare a function that returns a value.

Let's say we declare the variable `ten` the following way.

```js
let ten = 5 + 5
console.log(ten) // 10
```

We could replace those `5` with a function that `returns` this value. Let's call
this function `returnsFive`. The only new concept is the `return` keyword. It
returns the specified value **and** stop the function execution.

```js
let returnsFive = () => {
  return 5
  //     ↖ the keyword `return`, returns the value right after it,
  //       in this case the number 5.
}
```

Now that the function is declared, we call it where we need it.

```js
let ten = returnsFive() + returnsFive()
console.log(ten) // 10
```

Now a question that you might ask yourself is: What if we had several `return`
keywords in the same function ? Well as mentioned before, the `return` also
stops the function execution. So only the first `return` would matter. In fact
that means that anything after the `return` would not be executed. Example:

```js
let returnsFive = () => {
  return 5 // ONLY this return is executed. Everything else is ignored.
  return 10 // not executed (useless)
  return 'I am useless' // not executed either
  console.log('I am also useless') // nor this one
}
let ten = returnsFive() + returnsFive()
console.log(ten) // 10
//exactly the same result as the previous example
```

As you may see, we get exactly the same result as the previous example.
`returnsFive` only returns `5`. :)

### Instructions

As Rick's robot, you now know your purpose. (Remember ? `'You pass butter.'`)

Define the function `passButter` that returns the string `'The butter'`.

![robot](https://media.discordapp.net/attachments/489466992286498816/828181031991377930/butter-disapointed.png?width=717&height=241)
