## Get Account

Now you will interact with the node's API

### Instructions

For this exercise you will have to create a node script that :

- Loads an ethereum javascript library. I recommend `ethers.js`, `Web3.js` is a common alternative.
- create a function `getAccount()` that connects to a local node (`http://localhost:8545`) and returns the address of the first account available.
- export the function

```js
function getAccount() {
  //...
}
module.exports = getAccount
```

### Notions

- [hardhat](https://hardhat.org)
- [ganache](https://www.trufflesuite.com/ganache)
- [ethers](https://docs.ethers.io/)
- [web3](https://web3js.readthedocs.io/)
