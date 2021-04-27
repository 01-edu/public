## Fusion

### Instructions

The objective of this exercise is to merge objects into a new object depending on the values type

With this create a function called `fusion` that:

- If the type is an array you must concat it

```js
fusion([1, 2], [3, 4]) // -> [1,2,3,4]
```

- If it is a string you must concatenate with a space

```js
fusion('Salem', 'alem') // -> 'Salem alem'
```

- If it is numbers you must added them

```js
fusion(4, 11) // -> 15
```

- In case of type mismatch you must replace it with the value of the second object

```js
fusion({ a: 'hello', b: [] }, { a: 4 })
// -> { a: 4, b: [] }
```

- If it is an object you must fusion them recursively

```js
fusion({ a: 1, b: { c: 'Salem' } }, { a: 10, x: [], b: { c: 'alem' } })
// -> { a: 11, x: [], b: { c: 'Salem alem' } }
```
