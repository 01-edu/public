## Register

To conveniently check for a given string or document, you will create a register smart contract.

### Instructions

Write a solidity smart contract `Register` that provides the following :

- a function `addDocument()` that takes as parameter a hash (bytes32 in solidity)
- a function `getDate()` that takes as parameter a hash and returns the corresponding unix timestamp of the document

```js
contract Register {
  function getDate(bytes32 documentHash) public view returns (uint) {
  }
  function addDocument(bytes32 documentHash) public {
  }
}
```

### Hint

You may use block.timestamp to retrieve the current block date and a mapping (see Notions).

### Notions

- [Block information](https://docs.soliditylang.org/en/v0.4.21/units-and-global-variables.html#block-and-transaction-properties)
- [mapping](https://docs.soliditylang.org/en/v0.8.4/types.html#mapping-types)
