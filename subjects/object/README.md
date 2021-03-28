## Object

This exercise is about structuring multiple values together.

### Data Structures: Objects

In JS, `Object` are the most basic data structures, they are a way to group
values together.

They are like a bag of values.

#### Example

Remember that they are different types of variables:

```js
let currency = 'EURO'
let amount = 77.5
let cashPayment = false
```

Now we can group them all in an **object**, as objects are values too, let's
assign one to a `transaction` variable:

```js
let transaction = {
  currency: 'EURO',
  amount: 77.5,
  cashPayment: false,
}
console.log(transaction) // will show the object transaction
```

The variable `transaction` is declared and its value type is an object.

Let's explain each parts:

#### Object litteral syntax: `{}`

Starting with curly brackets `{}`, they are the delimiters of our object.

```js
let empty = {} // an empty object
```

#### Properties

They define what we want inside our objects. They are composed of two elements:

- a `key`
- and a `value`

```js
//                 ↙ begining of the declaration
let transaction = {
  //  ↙ property key
  currency: 'EURO',
  //           ↖ property value
}
// ↖ end of the declaration
```

We separate them with a `:`, to simplify, we will only use valid identifiers as
keys at the moment.

Each properties must be separated with a `,`

> Note that it's easier to always add a trailing `,` on every properties, but it
> is not required for the last property.

### Instructions

Declare a variable `human` which has a value **an object** with 3 properties:

- a `name` property of your name as a `String`
- an `age` property of your age as a `Number`
- a `secureLuggage` of a `Boolean` saying if your luggage contain dangerous
  things or not. _(still, for obvious security reasons)_

> “I paint objects as I think them, not as I see them.” \
> ― Pablo Picasso
