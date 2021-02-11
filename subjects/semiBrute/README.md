## semiBrute

### Instructions

Create a function `semiBrute` which takes one argument, a target, and returns string which hash `sha256` starts with the target. The target is an hexadecimal number with two characters.

### Example

```js
solution = semiBrute("e2")
console.log(solution)
// One valid result : 'abcdefghijklmnopqrs'
// Depending on your algorithm you might find other valid solutions
```

### Notions

- [Module crypto: hash](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_class_hash)
