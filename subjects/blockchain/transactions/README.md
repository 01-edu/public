## Transactions

_Sometimes science is more art than science. - Rick Sanchez_

Welcome and congratulations for choosing to learn blockchains and cryptocurrencies.

Today we get a first experience in the most core element, transactions. We will use the actual Bitcoin software to create transactions and send them. We will learn to interact with command lines, APIs and wallets to move cryptocurrencies around.

Most open blockchains offer three categories of networks:

- The main network `mainnet` on which value is transacted
- One or several `testnet`, that work more or less similarly to the mainnet with no value
- Local networks, to test locally with a node. It is called `regtest` on Bitcoin.

_🚨 Caution 🚨_
You should not use any crypto with value for any exercise of the module. If you already own crypto, we recommend that you use a new separate wallet to avoid any loss of funds. We will never use the main network of any blockchain `mainnet`.

## Content

### Mandatory

1. `Send Transaction` _Send a Bitcoin transaction_
2. `Retrieve Block Date` _get a block date_
3. `Retrieve Transaction Value` _get the value of a transaction_
4. `Send Transaction To Peer` _send a bitcoin transaction to a peer
5. `Send Ethereum Transaction` _Send a transaction to an address on a testnet_

### Optional

6. `Retrieve Transaction In Out` _get inputs and outputs from a transaction_

---

---

## Exercise 1: Send Transaction

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

---

---

## Exercise 2: Retrieve Block Date

### Instructions

Using Node.js, create and exports a function `retrieveBlockDate()` that takes as input the height of a block and returns the date of this block. The date must be a `date` javascript object. The script must connect to a local node available on `localhost:18443`, with the user `leeloo` and the password `multipass`.

### Usage

```js
retrieveBlockDate(1881467); // Expected : 2020-11-05T20:18:48.000Z
```

### Notions

- [Node.js https module](https://nodejs.org/api/https.html)
- [Async function in Javascript](https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Statements/async_function)
- [Exports module](https://nodejs.org/api/modules.html#exports-shortcut)
- [bitcoin core rpc](https://bitcoin.org/en/developer-reference#remote-procedure-calls-rpcs)
- [bitcoin-core js](https://www.npmjs.com/package/bitcoin-core)

---

---

## Exercise 3: Retrieve Transaction Value

### Instructions

Targeting a local bitcoin node, create a function `retrieveTransactionValue()` that takes as input a hash of a simple transaction, and returns the total value transferred in this transaction in bitcoins.

The Bitcoin node RPC interface is expected to be running with the following parameters:

```bash
rpcallowip=127.0.0.1
rpcport=18443
rpcuser=leeloo
rpcpassword=multipass
```

### Usage

```js
const { retrieveTxValue } = require("./retrieveTransactionValue.js");
retrieveTxValue(
  "d030023d96b9170af9ec2fe5d9b62a5eacbcbf144c68f3f45d68bca72d1d3649"
); // Expected : 50
```

### Hint

Internally, Bitcoin uses satoshis, 1 satoshi = 10^-8 bitcoin

### Notions

- [Node.js https module](https://nodejs.org/api/https.html)
- [Async function in Javascript](https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Statements/async_function)
- [satoshi](<https://en.bitcoin.it/wiki/Satoshi_(unit)>)

---

---

## Exercise 4: Send Transaction To Peer

You will send your first peer-to-peer transaction between two students.

### Instructions

- Create manually an address on **Bitcoin testnet** using any wallet. You can use bitcoin-cli configured on testnet or any other wallet.

- Get testnet Bitcoin from a faucet (see below).

- Send a transaction of 0.00001337 bitcoins to another student.

- Retrieve the main information from the transaction from a public API or a node (using for instance `bitcoin-cli gettransaction`) - Transaction hash "txid - Transaction fee "fee" - Transaction amount "amount" - Transaction date "time"
- Store the

### Usage

```js
exports.tx = {
  amount: -0.00001337,
  fee: -0.00000003,
  txid: "95952d9bf7542dfa0c98486495f1ae432a8738bbd7da051915d0aca1bec1f9",
  time: 1789670282,
  timereceived: 1789670282,
  hex: "02000000000101ab6873a9b39bc5be93ca6f75794aa235000000000feffffff0245e7052a0100000016001fb9bb786ac90f008c513fb4c545f21d561fa00000000146beada555da374454e9460220fced3cbbbd7e8ba02473044022b4a4e68743a0a51edb346228a54c0b7b1c00000000",
};
```

Congrats for your first real peer to peer Bitcoin transaction!

### Utilities

- A cross platform multi-currency wallet [Bitpay wallet](https://bitpay.com/wallet/)
- A Bitcoin explorer [bitcoin testnet explorer](https://blockstream.info/testnet/)
- [faucet 1](https://kuttler.eu/en/bitcoin/btc/faucet/)
- [faucet 2](https://bitcoinfaucet.uo1.net/)
- [faucet 3](https://testnet-faucet.com/btc-testnet/)

---

---

## Exercise 5: Send Ethereum Transaction

You will create your first transaction on Ethereum using one of its testnets. We recommend Goerli

### Instructions

- Create manually a wallet on any **Ethereum testnet**

- Get testnet Ether on a faucet (You might have to try several)

- Send a transaction of 0.000000000000001337 Ethers to the address '0x7A7a4EdC679bC4E29F74E32E9eEDd256cd435FBb'

- Create a js file and store the hash of your transaction in a variable `txid`

### Usage

```js
exports.txid =
  "0xf02c4a1487aa2e45fc2c77cb5a28713a1474d86d5f4292b264875ccc5da82b67";
```

### Notions

- A cross platform multi-currency wallet [Bitpay wallet](https://bitpay.com/wallet/)
- A cross platform Ethereum only wallet [MetaMask](https://metamask.io/)
- Find a [goerli faucet](https://letmegooglethat.com/?q=goerli+faucet)

---

---

## Exercise 6: Retrieve Transaction In Out (Optional)

A Bitcoin transaction consists of one or several inputs and one or several outputs. The sum of the inputs is slightly superior to the sum of the output to take into account the fee. One transaction per block, called "coinbase", pays miners and does not have a valid input. For some transactions, outputs might be null too.

This model is referred to as `UTXO` for "Unspent Transaction Outputs", as there is a ongoing list of unspent outputs.

#### A simple transaction

| -> In | Out -> |
| ----- | ------ |
| 0.50  | 0.30   |
|       | 0.19   |

_+ 0.01 fee_

### Instructions

Using a local Bitcoin node RPC API, create a function `retrieveTxValue()` that takes as input a hash of a simple transaction, and returns an object with an array of inputs values and an array outputs values in this transaction.

The Bitcoin node RPC interface is expected to be running with the following parameters:
rpcallowip=127.0.0.1
rpcport=18443

### Usage

```js
txHash = "d030023d96b9170af9ec2fe5d9b62a5eacbcbf144c68f3f45d68bca72d1d3649";
retrieveTxData(txHash);
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
