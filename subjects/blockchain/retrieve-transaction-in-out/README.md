## Retrieve Transaction In and Out

A Bitcoin transaction consists of one or several inputs and one or several outputs. The sum of the inputs is slightly superior to the sum of the output to take into account the fee. One transaction per block, called "coinbase", pays miners and does not have a valid input. For some transactions, outputs might be null too.

This model is referred to as `UTXO` for "Unspent Transaction Outputs", as there is a ongoing list of unspent outputs.

#### A simple transaction

| -> In | Out -> |
| ----- | ------ |
| 0.50  | 0.30   |
|       | 0.19   |

_+ 0.01 fee_

### Instructions

Using a local Bitcoin node RPC API, create a function `retrieveTransactionInOut()` that takes as input a hash of a simple transaction, and returns an object with an array of inputs values and an array outputs values in this transaction.

The Bitcoin node RPC interface is expected to be running with the following parameters:

```bash
rpcallowip=127.0.0.1
rpcport=18443
rpcuser=leeloo
rpcpassword=multipass
```

### Usage

```js
txHash = 'd030023d96b9170af9ec2fe5d9b62a5eacbcbf144c68f3f45d68bca72d1d3649'
retrieveTxData(txHash)
/* Expected : 
    { 
      in: [ 0.18075094 ], 
      out: [ 0.001, 0.1797493 ] 
    }
*/
```

### Notions

- [blockcypher API](https://www.blockcypher.com/dev/bitcoin/#blockchain-api)
- [Node.js https module](https://nodejs.org/api/https.html)
- [Node.js axios module](https://github.com/axios/axios)
- [Async function in Javascript](https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Statements/async_function)
