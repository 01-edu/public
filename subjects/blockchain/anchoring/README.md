## Anchoring

A simple use of blockhains is to anchor documents. We will hash documents and store them in a blockchain.

For this, we will use a local test node and interact with it using JavaScript

## Content

### Mandatory

1. `Local Node`
2. `Get Account`
3. `Send Ether`
4. `Send Hash`
5. `Check Document`
6. `Register`

### Optional

8. `Register With Events`

## Integration

In this Quest, tests can be run with

```sh
npm i
npx hardhat node&
npx mocha sendHash.test.js
```

---

---

## Exercise 1: Local Node

An Ethereum node provides different functionalities :

1. To be connected to other nodes and exchange blocks and transactions
2. To verify each transaction by running the Ethereum Virtual Machine
3. To provide an API to interact with the node
4. To provide a basic wallet to sign transactions and collect fees when mining

For now you will simply learn to launch a local node.

### Instructions

Launch a local test node. You can use any client (geth, openEthereum...), but I recommend javascript test nodes `hardhat` and `ganache-cli`.

**hardhat node**

```sh
npm i hardhat
npx hardhat node
```

**ganache**

```
npm i ganache-cli
npx ganache-cli
```

_You can also use ganache ("non cli") that provides a nice graphical interface. For compatibility, change the listening port to 8545 in the settings_

Noticeably, local javascript nodes are not connected to any network and provide already 10 account populated with test ether.

### Notions

- [hardhat](https://hardhat.org)
- [ganache](https://www.trufflesuite.com/ganache)

---

---

## Exercise 2: Get Account

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
module.exports = getAccount;
```

### Notions

- [hardhat](https://hardhat.org)
- [ganache](https://www.trufflesuite.com/ganache)
- [ethers](https://docs.ethers.io/)
- [web3](https://web3js.readthedocs.io/)

---

---

## Exercise 3: Send Ether

In this exercise, you will send ETH using a script.

### Instructions

Write a node script that provides a function `sendEther(amount, address)` which connects to a local node (`http://localhost:8545`) and sends an amount in Ether to the specified address.

```js
function sendEther(amount, address) {
  //...
}
module.exports = sendEther;
```

### Notions

- [hardhat](https://hardhat.org)
- [ganache](https://www.trufflesuite.com/ganache)
- [ethers](https://docs.ethers.io/)
- [web3](https://web3js.readthedocs.io/)

---

---

## Exercise 4: Send Hash

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
module.exports = sendHash;
```

### Notions

- [hardhat](https://hardhat.org)
- [ganache](https://www.trufflesuite.com/ganache)
- [ethers](https://docs.ethers.io/)
- [web3](https://web3js.readthedocs.io/)

---

---

## Exercise 5: Check Document

### Instructions

Write a node script that exports the function `checkDocument()` that :

- takes as parameter a transaction id and a string
- connects to a local node (`http://localhost:8545`)
- returns the date, as a unix timestamp, that the document was added to the blockchain if the transaction id and the string match and returns 0 otherwise

```js
function checkDocument(text, txID) {
  //...
}
module.exports = checkDocument;
```

### Usage

```js
checkDocument(
  "Hello world!",
  "0x49c8803ea126179502d59707dbcd9e9de15f5d441920936e9ec6fd78dd6468d8"
);
// Expected :
//1611104541
```

### Notions

- [ethers provider](https://docs.ethers.io/v5/api/providers/provider/#Provider--transaction-methods)

---

---

## Exercise 6: Register

To conveniently check for a given string or document, you will create a register smart contract.

### Instructions

Write a solidity smart contract `Register` that provides the following :

- a function `addDocument()` that takes as parameter a hash (bytes32 in solidity)
- a function `getDate()` that takes as parameter a hash and returns the corresponding unix timestamp of the document

```js
contract Register {
  function getDate(bytes32 documentHash) public view returns (uint) {
  }
  funct Hint

You may use block.timestamp to retrieve the current block date and a mapping (see Noion addDocument(bytes32 documentHash) public {
  }
}
```

###tions).

### Notions

- [Block information](https://docs.soliditylang.org/en/v0.4.21/units-and-global-variables.html#block-and-transaction-properties)
- [mapping](https://docs.soliditylang.org/en/v0.8.4/types.html#mapping-types)

---

---

## Exercise 7: Register With Events (Optional)

In order to facilitate the use of your Smart Contract, you will add events.

Events in solidity are used to facilitate the retrieval of information.

### Instructions

Write solidity smart contract `RegisterWithEvents` that provides:

- The functions `addDocument()` and `getDate()` as specified in the prior exercise.
- An event `event DocumentAdded(bytes32, uint)` that triggers when a new document is added to the register

### Notions

- [Events](https://docs.soliditylang.org/en/v0.8.4/contracts.html#events)
