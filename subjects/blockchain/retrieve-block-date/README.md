## Retrieve Block Date

### Instructions

Using Node.js, create and exports a function `retrieveBlockDate()` that takes as input the height of a block and returns the date of this block. The script must connect to a local node available on `localhost:18443`, with the user `leeloo` and the password
`multipass`.

### Usage

```js
retrieveBlockDate(1881467) // Expected : 1604607528

```

### Notions

- [Node.js https module](https://nodejs.org/api/https.html)
- [Async function in Javascript](https://developer.mozilla.org/fr/docs/Web/JavaScript/Reference/Statements/async_function)
- [Exports module](https://nodejs.org/api/modules.html#exports-shortcut)
- [bitcoin core rpc](https://bitcoin.org/en/developer-reference#remote-procedure-calls-rpcs)
- [bitcoin-core js](https://www.npmjs.com/package/bitcoin-core)
