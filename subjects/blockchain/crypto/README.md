## Crypto

_When in doubt, use brute force. - Ken Thompson_

Today , we will learn the fundamentals of cryptography that underlies all blockchain projects. We will practice binary variables, hash functions and digital signatures. By the end of the quest, you should be able to create a basic wallet to generate keys, store them and sign transaction.

Buffers are a builtin type in nodejs used to represent binary objects. You create a buffer from any object by using the function `Buffer.from()`. As usual in JavaScript there are a lot of implicit conversions. Be aware for instance that a string representing a number is different from the actual number.

For cryptography, you may have heard how it is used the encryption and decryption of messages. While this is often true, in the blockchain and cryptoassets industries this is not how we will primarily use it. The two families of algorithms we will see are:

**Cryptographic hash** functions are algorithms that take as input any data and produce an unique fingerprint of this data. Those functions are meant to be fast, one way, deterministic.

**Digital signatures** allow to identify the author of a message. It relies on asymmetric cryptography. You first generate randomly a key pair with a public key and a private key. The public key is shared publicly. The private key will allow you to sign a message. More precisely the hash of this message. The public key will allow anyone to verify that this message was properly signed.

From a practical point of view, for today, you only need the builtin library `crypto`.

# Content

## Mandatory

1. `Increment` _Binary variables in nodejs_
2. `Hash File` _Hash functions, file read in nodejs_
3. `Hash 160` _Hash functions, double hash, sha256, Ripemd160_
4. `Semi Brute` _Hash bruteforce_
5. `Signer` _ECDSA_

## Optional

6. `Generate Address` _Key generation, Crypto address_
7. `Basic Wallet` _Key storage, transaction signature_

## Integration:

Launch tests with

```sh
node test.mjs <exercices>
```

---

---

## Exercise 1: Increment

Cryptographic algorithms use a `binary` representation of variables internally (`Buffer` in nodejs). `Hexadecimal` representation is used to facilitate human reading. To get more familiar with the hexadecimal form, we will do a simple operation.

### Instructions

Create a function `increment` that takes as argument a number written in hexadecimal form and returns the same number incremented by one.

### Usage

```js
increment("03"); // expected : <Buffer 04>
increment("a0"); // expected : <Buffer a1>
increment("ff"); // expected : <Buffer 01 00>
increment("d537"); // expected : <Buffer d5 38>
```

### Notions

- [Node.js buffer](https://nodejs.org/api/buffer.html)
- [Wikipedia: Hexadecimal number](https://en.wikipedia.org/wiki/Hexadecimal)

---

---

## Exercise 2: Hash File

### Instructions

Create a function `hashFile(file)` that given the name of a file in the current folder, expected to be a string, returns the hash `sha256` of the content of this file as a string.

### Usage

_textfile.txt_

```
Sometimes science is more art than science.
```

```js
let hash = hashFile("textfile.txt");
console.log(hash);
// expected result : 575e6ccc6aa3882344d97a8ebae4b7fc4852f9149ad14007d11164450f751eb1
```

### Notions

- [Node.js _crypto_ hash](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_class_hash)
- [Node.js _fs_ readFileSync](https://nodejs.org/docs/latest-v14.x/api/fs.html#fs_fs_readfilesync_path_options)
- [Wikipedia: Hash_function](https://en.wikipedia.org/wiki/Hash_function)

---

---

## Exercise 3: Hash 160

### Instructions

Create a function `hash160` that takes one argument, expected to be a string, and returns the hash `sha256` of this argument hashed again with the `ripemd160` algorithm. The return value must be in a binary form.

Formally, it could be described as:

```
hash160 = sha256(ripemd160(input))
```

### Usage

```js
let hash = hash160("Ducks");
console.log(hash); // expected result : <Buffer de 5b 73 aa 85 84 02 d8 8c 36 d4 ff 85 29 65 d3 76 ac 6d 19>
```

### Notions

- [Node.js _crypto_](ttps://nodejs.org/docs/latest-v14.x/api/crypto.html)
- [Wikipedia: Hash_function](https://en.wikipedia.org/wiki/Hash_function)

---

---

## Exercise 4: Semi Brute

Hash functions are used to secure information. A piece of data, for instance a password, is hashed and only its hash is stored. While there is no way to compute on the original information, one can try every possible value. This is called a brute force attack.

Proof of Work algorithms work in a similar manner. Miners hash a block and modify it continuously to obtain a certain value.

We will create an example of this with a function that finds a string which hash starts with a given value.

### Instructions

Create a function `semiBrute()` that takes as argument a target, which is a two characters hexadecimal number, and returns a string which hash `sha256` starts with the target.

### Usage

```js
solution = semiBrute("e2");
console.log(solution);
// One valid result : 'abcdefghijklmnopqrs'
// You might find other valid solutions as we only check the first two characters
```

### Notions

- [Module crypto: hash](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_class_hash)
- [Password cracking (video)](https://www.youtube.com/watch?v=7U-RbOKanYs)

---

---

## Exercise 5: Signer

Elliptic curve cryptography is used in most blockchain projects to sign transactions. Using Node.js base library we will practise simple signatures.

### Instructions

- Create a function `init()` that generates a cryptographic key pair on any elliptic curve and returns the public key in any format.

- Create a function `signer(message)` that takes as arguments a message and returns the signature of the message (using the `sha256` algorithm and the keys generated with init).

- create a function `verifier(message, pubKey, signature)` that takes as arguments a message, a public key and a signature and returns a boolean if the signature is valid.

### Usage

```js
const message = "This is a message to sign";
const pubKey = init();
const signature = signer(message);
console.log(verifier(message, pubKey, signature));
// expected result :
// true
```

### Notions

- [Module crypto, sign](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_class_sign)
- [ECDSA](https://cryptobook.nakov.com/digital-signatures/ecdsa-sign-verify-messages)

---

---

## Exercise 6: Generate Address (Optional)

### Instructions

Create a function `generateAddress()` that generates a pair of cryptographic keys using the elliptic curve `secp256k1`. The return value is an object with members:

- `privateKey`, a private key as a string (`pkcs8/pem` format)
- `publicKey`, the corresponding public key as a string (`spki/pem`)
- `address`, an address to identify this account. The address is the hash `sha256` of the publicKey (in format `spki/der`) prepended with '01'

### Usage

```js
console.log(generateAddress())

// expected result :
{
  privateKey: '-----BEGIN PRIVATE KEY-----\n' +
    'MIGEAgEAMBAGByqGSM49AgEGBSuBBAAKBG0wawIBAQQg6wEoiAc7tEeOfh1HPVMw\n' +
    'z/eNG/D30JziVWjq8bW35MChRANCAAQP0WsGc3VmBsXO6Iz+LKxPPWNRb2DqUJKY\n' +
    '8Qh7qVvB4rExjUR0DhgUJrTC9tCPy3aiYGbXYLzuXTuOBcPbUiJq\n' +
    '-----END PRIVATE KEY-----\n',
  publicKey: '-----BEGIN PUBLIC KEY-----\n' +
    'MFYwEAYHKoZIzj0CAQYFK4EEAAoDQgAED9FrBnN1ZgbFzuiM/iysTz1jUW9g6lCS\n' +
    'mPEIe6lbweKxMY1EdA4YFCa0wvbQj8t2omBm12C87l07jgXD21Iiag==\n' +
    '-----END PUBLIC KEY-----\n',
  address: '01598fbb0e68eaa369d361f7b59157d80a91b2b78de74bc2ff71f6a44b272af368'
}
```

### Relevance

- The `secp256k1` curve is used several blockchains to generate keys.
- An user account is identified using an `address`. The address is computed from the public key using hash functions, splits and concatenations.
  - In Bitcoin `hash160` is used in the process and the address starts with '1', '3' or 'bc1'.
  - In Ethereum `keccak-256` is used and addresses are appended with 'Ox' that denotes an hexadecimal number.

### Notions

- [Node.js _crypto_ generateKeyPairSync](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_crypto_generatekeypairsync_type_options)
- [Bitcoin invoice address](https://en.bitcoin.it/wiki/Invoice_address)

---

---

## Exercise 7: Basic Wallet (Optional)

_- Do you always look at it encoded ?_
_- Well you have to, the image translators work for the construct program, but there is way too much information to decode the Matrix_

### Instructions

- Create a function `generateAddress()` that generates a cryptographic key pair on any elliptic curve and stores locally the private key in a `wallet.pem` file and returns the hash of the public key (`spki/der`) prepended with '01'

- Create a function `createTransaction(amount, recipient)` that takes as arguments an integer and a string. It returns a string with the following format:

`sender + recipient + hexAmount + signature`

- `sender` is the hash of the public key (`spki/der`) of the current user prepended with '01'
- `recipient` is the recipient taken as parameter
- `amount` is the amount taken as parameter in hexadecimal form
- `signature` is a digital signature of the first three fields, using the private key stored in `wallet.pem`

### Usage

```js
let address = generateAddress();
console.log(address);
// 01ae9596973b45d10e6a65c7fdcc593f63283bd405539c77e83938750d29729fac
let transaction = createTransaction(
  200,
  "01ffa17f4e0bbb2f049f5d1f9d4ab9fad967e6b0ed9a8fc94563d0b4e47b62e169"
);
console.log(transaction);
// 01ae9596973b45d10e6a65c7fdcc593f63283bd405539c77e83938750d29729fac01ffa17f4e0bbb2f049f5d1f9d4ab9fad967e6b0ed9a8fc94563d0b4e47b62e1692003045022017cdeb0e982e0705044caeb73db8d86992282f466c8b124cff0bfef64ae8be72022100ec4c6025d83a726bb483ebd05e53da37de28eaeb5db85068e355974c3caffaa4
```

### Notions

- [Node.js _crypto_ sign](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_class_sign)
- [Node.js _fs_ writeFileSync](https://nodejs.org/docs/latest-v14.x/api/fs.html#fs_fs_writefilesync_file_data_options)
- [Node.js _fs_ readFileSync](https://nodejs.org/docs/latest-v14.x/api/fs.html#fs_fs_readfilesync_path_options)
