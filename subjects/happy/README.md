## 🌟 Happy ?

### Notions

- [Includes](https://devdocs.io/javascript/global_objects/string/includes)
- [Conditional (ternary) Operator](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Conditional_Operator)

### Instructions

With all that you know now you are not a truly happy, efficient and independent
robot.

You have decided that if someone ask you the question `Are you happy ?` or
questions that includes the word `happy` you will always says that it is `true`.
To do so you will declare the function `happy` which takes only one argument and
returns a boolean.

Examples:

```js
console.log(happy('Are you happy ?')) // true
console.log(happy('Happy ?')) // true
console.log(happy('Are you happy!')) // false  -> this is not a question
console.log(happy('Are you sad ?')) // false  -> wrong question
```

Any other requests will be denied with a `false`. (if it is not a question of if
it does not have the word `happy`).

And because you aim for efficiency you will contruct your fonction without the
curlys brackets `{}` and without the keyword `return` or `if`.
