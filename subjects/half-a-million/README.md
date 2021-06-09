## Half a million

### Using variables as keys

One possible and very useful way to get something from an object is using a
variable value. Sometimes you don't know in advance what the key will be.

Using variables allows your code to be flexible, let's see how to do it

```js
let greeting = 'hello'
let say = { hello: 'my friend', hoy: 'mate' }

// Here when we access it will use
console.log(say[greeting]) // -> 'my friend'

// It is exactly the same as doing
console.log(say['hello']) // -> 'my friend'

// But if we change the value of our variable:
greeting = 'hoy'

// The same code now returns another value
console.log(say[greeting]) // -> 'mate'
```

We are using the `value` of the variable, not the `identifier` to access the
property, so:

```js
console.log(say[greeting]) // -> 'mate'
console.log(say['greeting']) // -> undefined
console.log(say.greeting) // -> undefined
// greeting is the identifier, but we use it's value, here `'hoy'`
```

### Instructions

In this exercise, we will do a heist.

Our intel already worked the plan out and we are able to provide you with 2
variables:

- a `vault` object
- a `secret` string

Your goal is to steal the `secret` matching value from the `vault`.

Declare a variable `loot` and assign it's value using the variable `secret`
value as a key to access the matching value from the `vault`.

You must then replace the value from the vault's secret by a message of your
choice.

**Example**: if `vault` is `{ '53CR37-k0D3': '1/2 MM $' }` and `secret` is
`'53CR37-k0D3'`:

- `loot` should be `1/2 MM $`
- `vault` should be `{ '53CR37-k0D3': 'I stole your $$' }`

> I sure hope you didn't do anything stupid, Jackie. \
> ― Ray Nicolette
