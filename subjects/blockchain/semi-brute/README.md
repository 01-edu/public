## Semi Brute

Hash functions are used to secure information. A piece of data, for instance a password, is hashed and only its hash is stored. While there is no way to compute on the original information, one can try every possible value. This is called a brute force attack.

Proof of Work algorithms work in a similar manner. Miners hash a block and modify it continuously to obtain a certain value.

We will create an example of this with a function that finds a string which hash starts with a given value.

### Instructions

Create a function `semiBrute()` that takes as argument a target, which is a two characters hexadecimal number, and returns a string which hash `sha256` starts with the target.

### Usage

```js
solution = semiBrute('e2')
console.log(solution)
// One valid result : 'abcdefghijklmnopqrs'
// You might find other valid solutions as we only check the first two characters
```

### Notions

- [Module crypto: hash](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_class_hash)
- [Password cracking (video)](https://www.youtube.com/watch?v=7U-RbOKanYs)
