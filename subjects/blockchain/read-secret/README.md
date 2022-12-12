# Read a Secret (Optional)

### Instructions

Create a web page, `readSecret.html` that loads an ethereum library, connects to ethereum testnet `kovan` and displays a secret contained in a smart contract.

The smart contract is available at the address `0x27c8dda37a22a29cb78320bf5e1c81ca087b2f8e` on Kovan testnet. It contains one function `getSecret()` that returns a string.

You might use the following interface :

```js
const abi = ["function getSecret() view returns (string)"];
```

### Notions

- [ethers contract](https://docs.ethers.io/v5/api/contract/contract/)
- [Infura](https://infura.io/)
- [Alchemy](https://www.alchemy.com/supernode)
- [web3](https://web3js.readthedocs.io/en/v1.3.4/web3-eth.html)

### Relevance

Reading your first smart contract.