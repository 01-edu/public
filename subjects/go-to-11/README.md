## Go to 11

### Looping with the `while` keyword

Loops are one of the last missing basic blocks of programing left to learn,
let's see it in action:

```js
while (count < 30) {
  // Code to be repeated here
}
```

As you can see, it is exactly like an `if`, we have a scope `{}` and a
`(condition)` just a new keyword.

Unlike the `if` however, the code inside the scope will be executed as long as
the condition is `true`

It is easy to get stuck in an never ending loop as if your condition is never
`false` you will always be stuck in a loop

As such it is common to have some kind of iteration inside our looping code:

```js
while (count < 30) {
  count = count + 1
}
```

Now everytime our code is executed, the value of `count` will be `incremented`
by `1`, until it reach `30` and stop.

> If your code is stuck in an infinite loop, you can do a Ctrl+C shortcut in the
> terminal to interrupt the code

### Instructions

Write a function `goTo11` that takes a `start` number argument and call
`console.log` on each numbers from the `start` (included) up to and including
`11`.

If the `start` is higher than `11` your function must do nothing.

[These go to Eleven... — "This is Spinal Tap" (1984)](https://www.youtube.com/watch?v=hW008FcKr3Q)
