## Increment

Cryptographic algorithms use a `binary` representation of variables internally (`Buffer` in nodejs). `Hexadecimal` representation is used to facilitate human reading. To get more familiar with the hexadecimal form, we will do a simple operation.

### Instructions

Create a function `increment` that takes as argument a number written in hexadecimal form and returns the same number incremented by one.

### Usage

```js
increment('03') // expected : <Buffer 04>
increment('a0') // expected : <Buffer a1>
increment('ff') // expected : <Buffer 01 00>
increment('d537') // expected : <Buffer d5 38>
```

### Notions

- [Node.js buffer](https://nodejs.org/api/buffer.html)
- [Wikipedia: Hexadecimal number](https://en.wikipedia.org/wiki/Hexadecimal)
