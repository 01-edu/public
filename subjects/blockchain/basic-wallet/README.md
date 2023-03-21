## Basic Wallet

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
let address = generateAddress()
console.log(address)
// 01ae9596973b45d10e6a65c7fdcc593f63283bd405539c77e83938750d29729fac
let transaction = createTransaction(
  200,
  '01ffa17f4e0bbb2f049f5d1f9d4ab9fad967e6b0ed9a8fc94563d0b4e47b62e169',
)
console.log(transaction)
// 01ae9596973b45d10e6a65c7fdcc593f63283bd405539c77e83938750d29729fac01ffa17f4e0bbb2f049f5d1f9d4ab9fad967e6b0ed9a8fc94563d0b4e47b62e1692003045022017cdeb0e982e0705044caeb73db8d86992282f466c8b124cff0bfef64ae8be72022100ec4c6025d83a726bb483ebd05e53da37de28eaeb5db85068e355974c3caffaa4
```

### Notions

- [Node.js _crypto_ sign](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_class_sign)
- [Node.js _fs_ writeFileSync](https://nodejs.org/docs/latest-v14.x/api/fs.html#fs_fs_writefilesync_file_data_options)
- [Node.js _fs_ readFileSync](https://nodejs.org/docs/latest-v14.x/api/fs.html#fs_fs_readfilesync_path_options)
