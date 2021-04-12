## Griswold the Blacksmith ⚒️

Methods and built-in functions are good, but most of the time we have to write
our own logic and the first block for that are **conditions**.

### The `if` keyword

The `if` keyword allows you to run lines of code _only if_ the condition is
right, example:

```js
if (age > 48) {
  console.log('You are over 48 years old')
}
```

### Condition `()`

following our `if` is a `condition` delimited by `()` parentheses,

### Comparaison operators `>`

Inside our condition is a comparaison (using the _greater than_ operator `>`).\
This code `if (age > 48)` reads:

> If age is greater than 48, then do the following code

There are 6 different comparaison opperators:

- `>` _greater than_
- `<` _lesser than_
- `<=` _lesser than or equal_
- `>=` _greater than or equal_
- `===` _equal to_
- `!==` _not equal to_

Every comparaison operator are _like_ functions, they take 2 arguments, one on
the right and one on the left, and returns a boolean value. Either `true`, if
the condition is met, or `false` if it's not.

Since they return value, you can assign them to variables, just like functions
return values:

```js
let age = 5
let ageEqual5 = age === 5 // age is equal to 5
let ageNotEqual5 = age !== 5 // age is not equal to 5
console.log(ageEqual5) // true
console.log(ageNotEqual5) // false
```

But they are commonly used directly inside an `if` condition.

### Scope `{}`

After the condition, a curly brace `{` signal the begining of a scope. The scope
ends at the enclosing `}` a few lines after.

Scopes are a way to group lines of code, this allow us to do multiple lines of
code if a condition is true.

```js
if (age > 48) {
  // <-beginning of the scope of the if condition
  console.log('You are over 48 years old')
} // <- end of the scope of the if condition
```

### Indentation `..` _(2 spaces)_

Upon writing code inside a scope, it's an important convention to **indent** it.

Indenting is when spaces are added at the beginning of the line, here is an
example of bad code:

<!-- prettier-ignore-start -->
```js
if (age > 48) {
// <- without indentation ! bad bad ! unreadable !!
console.log('You are over 48 years old')
}
```
<!-- prettier-ignore-end -->

good code:

```js
if (age > 48) {
  // <- with indentation, omg so clean, amazing !
  console.log('You are over 48 years old')
}
```

Indenting add a visual indication that the code is inside a scope, while it's
not strictly necessary for code to work, it will become very important to keep
the code clear.

### Instructions

You are Griswold the Blacksmith, and you must give the list of items the player
can buy with the money he has, here is what you are selling:

- arrows for 3 coins
- boots for 44 coins
- sword for 299.99 coins _(limited offer)_

Declare a `purchasableGoods` array and _conditionally_ push to it all the goods
that the player can buy.

Use `if` conditions and compare the cost of the goods with the provided variable
`playerCoins` that contains the number of coins available to the player.

> You must order elements by price
