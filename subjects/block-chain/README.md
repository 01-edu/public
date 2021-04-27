## Block Chain

### Instructions

Create a `blockChain` that create a block in your very own block chain.

the function takes 2 arguments:

- `data` any valid JSON data
- `prev` the previous block, if no block are given it should use the
  genesis block: `{ index: 0, hash: '0' }`

A block must have the following properties:

- `index`
- `hash` a computed hash using the concatenation of the `index`, `hash`
  and stringified `data` and hashing all of it using the provided `hashCode`.
- `data` the data (not encoded in JSON)
- `prev` the previous block
- `chain` a function that takes a new `data` and create the next block with it.

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

- [devdocs.io/javascript/global_objects/json/stringify](https://devdocs.io/javascript/global_objects/json/stringify)

### Code provided

> all code provided will be added to your solution and doesn't need to be submited.

```js
const hashCode = str =>
  (
    [...str].reduce((h, c) => (h = (h << 5) - h + c.charCodeAt(0)) & h, 0) >>> 0
  ).toString(36)
```
