## Transfers History

### Instructions

- Create a JavaScript file `transfersHistory.js`that exports the function `transfersHistory(address, address)` that takes as parameters the address of a `EventfulToken` Smart Contract instance and the address of an user.

The function will connect to a local node and retrieve the events. It returns an array with all the transfers from that account, starting from the most recent, positive if it is toward the user, and negative if it is outside.

#### Structure

```js
const { ethers } = require("ethers");

async function transfersHistory(contractAddress, userAddress) {...}
exports.transfersHistory = transfersHistory;
```

#### Usage

```js
const { transfersHistory } = require('./transfersHistory.js')

transfersHistory(
  '0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc',
  '0x70997970c51812dc3a010c7d01b50e0d17dc79c8',
)
/* expected:
[50, -25, 200, 30, -230]
*/
```

### Resources

- [solidity docs](https://docs.soliditylang.org/)
- [ethers](https://learnxinyminutes.com/docs/solidity/)
- [web3 : contracts events](https://web3js.readthedocs.io/en/v1.2.11/web3-eth-contract.html#contract-events)
