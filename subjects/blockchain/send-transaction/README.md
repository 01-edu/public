# Send Transaction

The purpose of this exercise is to create a simple Bitcoin transaction locally using the command line. As we use new tools this exercise is guided.

### Instructions

- Install a Bitcoin node. There are two solutions :
  - Use a preconfigured image for a virtual machine such as [cryptotux](https://cryptotux.org/)
  - Install manually by downloading and installing the Bitcoin Core software from [github](https://github.com/bitcoin/bitcoin/releases).
- Launch Bitcoin core daemon on _regtest_ network. You can do so with `bitcoind -regtest -fallbackfee=0.00000003` assuming the executable is in your shell's path.
- Create manually a wallet and two addresses on _Bitcoin regtest_. You may do this by running the following commands
  ```bash
  bitcoin-cli createwallet "testwallet"
  bitcoin-cli getnewaddress
  bitcoin-cli getnewaddress
  ```
- Generate 101 blocks to get fresh bitcoins. You need 101 blocks as the Bitcoin you receive from mining are locked for 100 blocks.
  You may use the following command `bitcoin-cli -regtest generatetoaddress 101 <your address>`

- Send a transaction of 2 bitcoins to the second address.

- List the last transactions `bitcoin-cli listtransactions`
- Create a js file and exports the hash of your transaction in a variable `txid`.

### Usage

```js
exports.txid =
  "be3d0d245e7dce50964ac9157aee7e18a3828e11d89f72ee0bc3f76b526e5bb";
```

Congrats for your first Bitcoin transaction!

### Notions

- A linux image with developer tools [cryptotux.org](https://cryptotux.org/)
- Bitcoin core node [bitcoin.org](https://bitcoin.org/)
- Send to address [reference](https://wbnns.github.io/bitcoin-dev-docs/reference/rpc/sendtoaddress.html)
