## Access 🔑

Now that we know how to define objects, let's start to use them

### The dot opperator `.` (property accessor)

Let's start with getting values from them. Well, turns out you already have been
doing it a lot, remember `console.log` ?

The `.` here is _accessing_ the property `log` from the `console` object.

Taking the previous example `transaction` object, we can access properties from
it using `.` same way we got the `log` function:

```js
let transaction = {
  currency: 'EURO',
  amount: 77.5,
  cashPayment: false,
}

console.log(transaction) // Will log the whole transaction
console.log(transaction.amount) // will only log the amount of the transaction
```

Accessing a property with `.` only give you it's value, and because it is a
value you can use it like any other values:

```js
let taxes = 1.2 // let's define 20% taxes
let transaction = {
  currency: 'EURO',
  amount: 77.5,
  cashPayment: false,
}

const totalWithTaxes = transaction.amount * taxes

console.log(totalWithTaxes) // will log 93 (77.5 * 1.2)
```

### Instructions

We will provide a `human` variable of type object just like the one you did in
the previous exercise

Your job will be to decompose each property in its own variable:

- define a `name` variable with the value of the `name` property of the `human`
  variable
- same for `age`
- and same for `secureLuggage`
