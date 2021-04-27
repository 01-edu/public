## Flagger

### Instructions

Create a function called `flags` that receives an object and returns
the specific aliases and descriptions from the properties of that object.

The `help` flag:

- Should be present in the output by default.
- When not present in the input, it should return the description of all flags.
  But when present it specifies the descriptions of the flags that are passed to help. (ex: `help: ['divide']`)

#### Example:

```js
{
  multiply: 'multiply the values',
  divide: 'divides the values',
  help: ['divide']
}
```

and outputs :

```js
{
  alias: { h: 'help', m: 'multiply', d: 'divide'}
  description: '-d, --divide: divides the values',
}
```
