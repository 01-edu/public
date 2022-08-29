## Flow

### Instructions

Create the function named `flow` that will act like the `_.flow([funcs])` from lodash.

### Example

```js
const square = (nbr) => nbr * nbr
const add2Numbers = (nbr1, nbr2) => nbr1 + nbr2

const flowedFunctions = flow([add2Numbers, square])
flowedFunctions(2, 3) // -> 25
```

### Notions

- [lodash.com/docs/4.17.15#flow](https://lodash.com/docs/4.17.15#flow)
