## Anything to Declare ? 🛂

### Variables

Values, need a way to be identified, that's why we use variables.\
They add meaning to a value by _pointing_ to it.

It's like a **label**, a way to name things.

If we say `20`, it doesn't carry much meaning, _`20` what ?_

Imagine we are talking about what's in your backpack, you have 20 pairs of
socks.

_Now that's a lot of socks !_

> We defined _what_ we have (pair of socks) and it's _value_ (20)

#### Example

An `identifier` is used to _define_ what it is, using this syntax:

```js
let pairOfSocks = 20
```

> 😱 Woa what's all this ?!

Let's explain each parts:

#### Keyword: `let`

Firstly, a keyword, `let`.

> A `keyword` is a special word that JS knows about, it is used to tell the
> computer to perform a specific action.

`let` that indicate the script that it's defining a new variable.

#### Identifier

After that, it need a **valid** identifier.

In this case it's `pairOfSocks`, we chose what ever we want here that will be
meaningful, _(it's often hard to name things correctly)_.

A few rules to apply to make sure an identifier is valid:

- No space allowed _(`pair of socks` would be 3 distinct identifiers)_
- Not **starting** with a number _(that's reserved for number values)_
- Not being a reserved keyword _(for example using `let`)_
- No special characters

As such we use what's called `camelCase`.

> Note that in JS, it is a convention to not uppercase the first letter as this
> is reserved for special declarations, we won't go into details now.

```js
let pair of socks = 20 // invalid because of spaces
let 'pair of socks' = 20 // invalid because identifiers are not strings
let pair-of-socks = 20 // invalid because of special character -
let pair.of.socks = 20 // invalid because of special character /
let 20PairOfSocks = 20 // invalid because beginning with a number
let PairOfSocks = 20 // valid but incorrect because of the uppercase
let pairOfSocks = 20 // Just right
let let = true // invalid because `let` is a JS keyword
```

#### Operator: `=`

The special character `=` is an **operator**, like in math, they are used to
define specific operations.

In this case, `=` define the `assignation` operation.

It means assigning a value to our variable.

This is what **links** the choosen `identifier` with our `value`.

#### Value

Lastly, a value, like the one you already know: `string`, `number` and
`boolean`.

Full example with descriptive comments:

```js
// ↙ keyword        ↙ assignation operator
let comicBookTitle = 'Tintin in Tibet'
//       ↖ identifier       ↖ the value (here a string)
```

Using multiple variables to define something more complex:

```js
// Example of variables that could represent a transaction:
let currency = 'EURO'
let amount = 77.5
let cashPayment = false

// Use them with console.log, like a normal value:
console.log('You have to pay:')
console.log(amount)
console.log('in')
console.log(currency)
console.log('using cash:')
console.log(cashPayment)
```

### Instructions

All right, before we can embark on this adventure, you are going to tell us
more about yourself using **variables**.

Declare three variables:

- `age`: your age as a `number`

- `name`:  your name as a `string`

- `secureLuggage`: which will be a `boolean` stating whether or not your
  luggage contains dangerous materials. _(for obvious security reasons)_

> PS: Remember you are trying to board a plane, so use reasonable values.
