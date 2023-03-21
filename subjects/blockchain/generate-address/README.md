## Generate Address

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
