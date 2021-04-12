## Seeker of Truth

### Truthy and Falsy

In JS, all values are either **truthy** or **falsy**, this means if used in a
condition, truthy values will validate the condition while falsy values would
not.

Here is the list of falsy values:

- `undefined` and `null`
- the numbers `0` and `NaN`
- the empty string `''`
- and the boolean `false` of course

All other values are truthy, note that empty arrays and empty objects are
truthy, but the empty string is not.

```js
if ('') {
  console.log('Since empty string are falsy, this will never log')
}

if ('hello') {
  console.log('this will always log as the string is not empty')
}
```

### Operator `!`

The `!` (NOT operator) can be used to convert a truthy value to `false` or a
falsy value to `true`.

example:

```js
let money = 0
let noMoney = !money

console.log(noMoney) // true
```

In this case, 0 is falsy, so the `!` return the value `true`

### Instructions

Seek the truth and claim your verdict !

- Log `'The truth was spoken.'` if the value of the provided variable `truth` is
  truthy
- Log `'Lies !!!!'` if the value of the provided variable `truth` is falsy
