## Is

### Instructions

Add new function properties to the `is` object to check value types. Each function should take one argument, and return a `boolean`.

- `is.num`: value is a `number`.
- `is.nan`: value is `NaN`.
- `is.str`: value is a `string`.
- `is.bool`: value is a `boolean`.
- `is.undef`: value is `undefined`.
- `is.def`: value is defined.
- `is.arr`: value is an `array`.
- `is.obj`: value is a simple object or `null` objects.
- `is.fun`: value is a function.
- `is.truthy`: value is truthy.
- `is.falsy`: value is falsy.

### Examples

```js
console.log(is.num(5))
// output: true

console.log(is.num('ciao'))
// output: false
```

### Notions

- [Primitives and operators](https://nan-academy.github.io/js-training/examples/primitive-and-operators.js)
- [typeof](https://devdocs.io/javascript/operators/typeof)
- [Truthy](https://developer.mozilla.org/en-US/docs/Glossary/Truthy)
- [Falsy](https://developer.mozilla.org/en-US/docs/Glossary/Falsy)

### Code provided

> The provided code will be added to your solution, and does not need to be submitted.

```js
const is = {}
```
