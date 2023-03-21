## Retrieve transaction value

### Instructions

Targeting a local bitcoin node, create a function `retrieveTransactionValue()` that takes as input a hash of a simple transaction, and returns the total value transferred in this transaction in bitcoins. If several outputs are sent, the value should be summed up.

The Bitcoin node RPC interface is expected to be running with the following parameters:

```bash
rpcallowip=127.0.0.1
rpcport=18443
rpcuser=leeloo
rpcpassword=multipass
```

### Usage

```js
const { retrieveTxValue } = require('./retrieveTransactionValue.js')
retrieveTxValue(
  'd030023d96b9170af9ec2fe5d9b62a5eacbcbf144c68f3f45d68bca72d1d3649',
) // Expected : 50
```

### Hint

Internally, Bitcoin uses satoshis, 1 satoshi = 10^-8 bitcoin

### Notions

- [Node.js https module](https://nodejs.org/api/https.html)
- [Async function in Javascript](https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Statements/async_function)
- [satoshi](<https://en.bitcoin.it/wiki/Satoshi_(unit)>)
