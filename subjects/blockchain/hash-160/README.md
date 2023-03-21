## Hash160

### Instructions

Create a function `hash160` that takes one argument, expected to be a string, and returns the hash `sha256` of this argument hashed again with the `ripemd160` algorithm. The return value must be in a binary form.

Formally, it could be described as:

```
hash160 = sha256(ripemd160(input))
```

### Usage

```js
let hash = hash160('Ducks')
console.log(hash) // expected result : <Buffer de 5b 73 aa 85 84 02 d8 8c 36 d4 ff 85 29 65 d3 76 ac 6d 19>
```

### Notions

- [Node.js _crypto_](ttps://nodejs.org/docs/latest-v14.x/api/crypto.html)
- [Wikipedia: Hash_function](https://en.wikipedia.org/wiki/Hash_function)
