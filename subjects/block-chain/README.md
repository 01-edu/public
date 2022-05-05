## Block Chain

### Instructions

Create a function named `blockChain` that creates a block in your very own block chain. It takes 2 arguments:

- `data`: any valid JSON data.
- `prev`: the previous block, if no block are given it should use the genesis block: `{ index: 0, hash: '0' }`.

A block must have the following properties:

- `index`
- `hash`: a computed hash using the `hashCode` function provided. You will need to pass it a concatenation of the block's `index`, the previous block's `hash` and the block's stringified `data`.
- `data`: any valid object.
- `prev`: the previous block.
- `chain`: a function that accepts `data` as an argument, and creates the next block with it.

### Examples

```js
const first = blockChain({ a: 1 })
console.log(first.index) //           -> 1
console.log(first.data) //            -> { a: 1 }
console.log(first.prev) //            -> { index: 0, hash: "0" }
console.log(first.hash) //            -> '1103f27'
console.log(hashCode('10{"a":1}')) // -> '1103f27'

const second = first.chain({ hello: 'world' })
console.log(second.hash) //                           -> '18drvvc'
console.log(hashCode('21103f27{"hello":"world"}')) // -> '18drvvc'

const chain = second
  .chain({ value: 4455 })
  .chain({ some: 'data' })
  .chain({ cool: 'stuff' })

const fork = second
  .chain({ value: 335 })
  .chain({ some: 'data' })
  .chain({ cool: 'stuff' })

console.log(chain.hash) //  -> '1qr3qfs'
console.log(fork.hash) //   -> '1x9gsc1'
console.log(chain.index) // -> 5
console.log(fork.index) //  -> 5
```

### Notions

- [JSON.stringify](https://devdocs.io/javascript/global_objects/json/stringify)

### Code provided

> The provided code will be added to your solution, and does not need to be submitted.

```js
const hashCode = str =>
  (
    [...str].reduce((h, c) => (h = (h << 5) - h + c.charCodeAt(0)) & h, 0) >>> 0
  ).toString(36)
```
