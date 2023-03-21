## Read a Secret (Optional)

### Instructions

Create a web page, `read-secret.html` that loads an ethereum library, connects to ethereum testnet `Sepolia` and displays a secret contained in a smart contract.

The smart contract is available at the address `0x467782A5ab90af6baA6f8af0b4E69A7Ddb197fF0` on Sepolia testnet. It contains one function `getSecret()` that returns a string.

You might use the following interface :

```js
const abi = ['function getSecret() view returns (string)']
```

⚠️ As the test does not reach the internet:

- Store in a hardcoded manner the result in an element with `storedSecret` as id.
- Detect when internet is not available and skip the connection to the provider.

### Notions

- [ethers contract](https://docs.ethers.io/v5/api/contract/contract/)
- [Infura](https://infura.io/)
- [Alchemy](https://www.alchemy.com/supernode)
- [web3](https://web3js.readthedocs.io/en/v1.3.4/web3-eth.html)

### Relevance

Reading your first smart contract.
