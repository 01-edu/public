## send-transaction

The purpose of this exercise is to create a simple Bitcoin transaction locally using the command line. As we use new tools this exercise is guided.

### Instructions

- Install a Bitcoin node. There are two solutions :

  - Use a preconfigured image for a virtual machine such as [cryptotux](https://cryptotux.org/)

  - Install manually by downloading and installing the Bitcoin Core software from [GitHub](https://github.com/bitcoin/bitcoin/releases). You can copy and paste the following text in your shell and execute.

```bash
export bitcoinCoreVersion="24.0.1"
wget -q "https://bitcoincore.org/bin/bitcoin-core-$bitcoinCoreVersion/bitcoin-$bitcoinCoreVersion-x86_64-linux-gnu.tar.gz"
tar xzf "bitcoin-$bitcoinCoreVersion-x86_64-linux-gnu.tar.gz"
sudo mv bitcoin-$bitcoinCoreVersion/bin/* /usr/local/bin
rm -rf bitcoin-$bitcoinCoreVersion/
rm "bitcoin-$bitcoinCoreVersion-x86_64-linux-gnu.tar.gz"
```

- Launch Bitcoin core daemon on _regtest_ network. You can do so with the following command assuming the executable is in your shell's path.

```console
$ bitcoind -chain=regtest -daemon -fallbackfee=0.00000003
$
```

- Create manually a wallet and two addresses on _Bitcoin regtest_. You may do this by running the following commands

```console
$ bitcoin-cli -regtest createwallet "testwallet"
...
$ bitcoin-cli -regtest getnewaddress
...
$ bitcoin-cli -regtest getnewaddress
...
$
```

- Generate 101 blocks to get fresh bitcoins. You need 101 blocks as the Bitcoin you receive from mining are locked for 100 blocks. You may use the following command:

```console
$ bitcoin-cli -regtest generatetoaddress 101 <your address>
...
$
```

- Send a transaction of 2 bitcoins to the second address.

```console
$ bitcoin-cli -regtest sendtoaddress <the second address> 2
...
$
```

- List the last transactions

```console
$ bitcoin-cli -regtest listtransactions
...
$
```

- Create a JS file `send-transaction.js` with a variable `txid` that will store the hash of your transaction. The variable `txid` needs to be importable from an external file! Below an example on how you should export your variable.

```js
exports.txid = ...
```

Congrats for your first Bitcoin transaction!

### Notions

- Bitcoin core node [bitcoin.org](https://bitcoin.org/)
- [Bitcoin for developers introduction](https://wbnns.github.io/bitcoin-dev-docs/examples/intro.html)
- A Linux image with developer tools [cryptotux.org](https://cryptotux.org/)
