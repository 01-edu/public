## Signer

Elliptic curve cryptography is used in most blockchain projects to sign transactions. Using Node.js base library we will practise simple signatures.

### Instructions

- Create a function `init()` that generates a cryptographic key pair on any elliptic curve and returns the public key in any format.

- Create a function `signer(message)` that takes as arguments a message and returns the signature of the message (using the `sha256` algorithm and the keys generated with init).

- create a function `verifier(message, pubKey, signature)` that takes as arguments a message, a public key and a signature and returns a boolean if the signature is valid.

### Usage

```js
const message = 'This is a message to sign'
const pubKey = init()
const signature = signer(message)
console.log(verifier(message, pubKey, signature))
// expected result :
// true
```

### Notions

- [Module crypto, sign](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_class_sign)
- [ECDSA](https://cryptobook.nakov.com/digital-signatures/ecdsa-sign-verify-messages)
