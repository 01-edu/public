## Send Transaction to Peer

You will send your first peer-to-peer transaction between two students.

### Instructions

- Create manually an address on **Bitcoin testnet**. You can use bitcoin-cli configured on testnet or any other wallet.

- Get testnet Bitcoin from a faucet (see below).

- Send a transaction of 0.00001337 bitcoins to another student.

- Retrieve the main information from the transaction from a public API or a node (using for instance `bitcoin-cli gettransaction`). You might need to convert some fields depending on the API used.

- Store the following values in a variable `tx` in your solution `send-transaction-to-peer.js`:
  - Transaction hash `txid`.
  - Transaction fee `fee` (in Bitcoins, negative integer).
  - Transaction amount `amount` (in Bitcoins, negative integer).
  - Transaction date `time.

> The submitted transaction information needs to be related to a recent (no more than 48 h) transaction!

### Expected file

sendTransactionToPeer.js

```js
exports.tx = {
  amount: -0.00001337,
  fee: -0.00000003,
  txid: '95952d9bf7542dfa0c98486495f1ae432a8738bbd7da051915d0aca1bec1f9',
  time: 1789670282,
}
```

Congrats for your first real peer to peer Bitcoin transaction!

### Utilities

- A fully open source Bitcoin wallet software [electrum](https://electrum.org/)
  In order to use it on the testnet you will need to launch it from the command line with the `--testnet` option
- A full list of Bitcoin compatible wallets [bitcoin wallets](https://bitcoin.org/en/choose-your-wallet)
- A Bitcoin explorer [bitcoin testnet explorer](https://blockstream.info/testnet/)
- A Bitcoin public API documentation [esplora](https://github.com/Blockstream/esplora/blob/master/API.md)
- [faucet 1](https://kuttler.eu/en/bitcoin/btc/faucet/)
- [faucet 2](https://bitcoinfaucet.uo1.net/)
- [faucet 3](https://testnet-faucet.com/btc-testnet/)
