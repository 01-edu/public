## Local Node

An Ethereum node provides different functionalities :

1. To be connected to other nodes and exchange blocks and transactions
2. To verify each transaction by running the Ethereum Virtual Machine
3. To provide an API to interact with the node
4. To provide a basic wallet to sign transactions and collect fees when mining

For now you will simply learn to launch a local node.

### Instructions

Launch a local test node. You can use any client (geth, openEthereum...), but I recommend javascript test nodes `hardhat` and `ganache-cli`.

**hardhat node**

```shell
npm i hardhat
npx hardhat node
```

**ganache**

```shell
npm i ganache-cli
npx ganache-cli
```

_You can also use ganache ("non cli") that provides a nice graphical interface. For compatibility, change the listening port to 8545 in the settings_

Noticeably, local javascript nodes are not connected to any network and provide already 10 account populated with test ether.

> There is no code to submit for this exercise. Submit an empty file.

### Notions

- [hardhat](https://hardhat.org)
- [ganache](https://www.trufflesuite.com/ganache)
