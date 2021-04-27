## Get

### Instructions

Create the `get` function.
It takes 2 arguments:

- `src` an object
- `path` a string

And returns the value at the given string path.

### Example:

```js
const src = { nested: { key: 'peekaboo' } }
const path = 'nested.key'
get(src, path) // -> 'peekaboo'
```
