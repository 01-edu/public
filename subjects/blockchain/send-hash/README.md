## Send Hash

In this exercise, you will register the hash of a document on the blockchain.

### Instructions

Write a node script that provides a function `sendHash()` that:

- creates a hash `sha256` of a string provided as parameter
- connects to a local node (`http://localhost:8545`)
- sends a transaction to an address with the hash of the document in the `data` field

```js
function sendHash(text) {
  //...
}
module.exports = sendHash
```

### Notions

- [hardhat](https://hardhat.org)
- [ganache](https://www.trufflesuite.com/ganache)
- [ethers](https://docs.ethers.io/)
- [web3](https://web3js.readthedocs.io/)
