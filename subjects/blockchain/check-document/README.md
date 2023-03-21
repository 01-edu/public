## Check Document

### Instructions

Write a node script that exports the function `checkDocument()` that :

- takes as parameter a transaction id and a string
- connects to a local node (`http://localhost:8545`)
- returns the date, as a unix timestamp, that the document was added to the blockchain if the transaction id and the string match and returns 0 otherwise

```js
function checkDocument(text, txID) {
  //...
}
module.exports = checkDocument
```

### Usage

```js
checkDocument(
  'Hello world!',
  '0x49c8803ea126179502d59707dbcd9e9de15f5d441920936e9ec6fd78dd6468d8',
)
// Expected :
//1611104541
```

### Notions

- [ethers provider](https://docs.ethers.io/v5/api/providers/provider/#Provider--transaction-methods)
