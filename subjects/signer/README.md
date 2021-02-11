## Signer

### Instructions

- Create a function `init()` which generates a cryptographic key pair on any elliptic curve and returns the public key in any format.
- Create a function `signer(message)` which takes as arguments a message and returns the signature of the message (using the `sha256` algorithm).
- Create a function `verifier(message, pubKey, signature)` which takes as arguments a message, a public key (which will be the return of the function `init()`), a signature (which will be the return of the function `signer(message)`) and returns a boolean if the signature is valid.

### Example

```js
const message = 'This is a message to sign'
const pubKey = init()
const signature = signer(message)
console.log(verifier(message, pubKey, signature))
// expected result :
// true
```

### Relevance

Elliptic curve cryptography is used in most blockchain projects to sign transactions.

### Notions

- [Module crypto, sign](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_class_sign)
- [ECDSA](https://cryptobook.nakov.com/digital-signatures/ecdsa-sign-verify-messages)
