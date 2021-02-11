# Introduction

Today , we will learn the fundamentals of cryptography that underlies all blockchain projects. We will practise binary variables, hash functions and digital signatures. By the end of the quest, you should be able to create a basic wallet to generate keys, store them and sign transaction.

Buffers are a builtin type in nodejs used to represent binary objects. You create a buffer from any object by using the function `Buffer.from()`. As usual in JavaScript there are a lot of implicit conversions. Be careful that a string representing a number is not the same as the actual number.

For cryptography, you may have heard it is used the encryption and decryption of messages. While this is often true, in the blockchain and cryptoassets industry this is not how we will primarily use it. The two families of algorithms we will see are :

**Cryptographic hash** functions are algorithms that take as input any data and produce an unique fingerprint of this data. Those functions are meant to be fast, one way, deterministic.

**Digital signatures** allow to identify the author of a message. It relies on asymmetric cryptography. You first generate randomly a key pair with a public key and a private key.The public key is shared publicly. The private key will allow you to sign a message. More precisely the hash of this message. The public key will allow anyone to verify that this message was properly signed.

From a practical point of view, for today, you only need the builtin library `crypto`.

# Content

## Mandatory

1. [Increment](increment.md) _Binary variables in nodejs_
2. [hashFile](hashFile.md) _Hash functions, file read in nodejs_
3. [hash160](hash160.md) _Hash functions, double hash, sha256, Ripemd160_
4. [semiBrute](semiBrute.md) _Hash bruteforce_
5. [signer](signer.md) _ECDSA_

## Optional

6. [generateAddress](generateAddress.md) _Key generation, Crypto address_
7. [basicWallet](basicWallet.md) _Key storage, transaction signature_
