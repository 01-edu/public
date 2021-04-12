## Time to Pay 💸

### `||` (the OR operator)

The other way to group condtions:

```js
if (user.role === 'admin' || user.role === 'moderator') {
  console.log(user.name, 'has access to moderation pages.')
}
```

Here, the code will only show the message if **any** conditions are true.

### Instructions

You are selling planes tickets, they all cost `9.99$` you must confirm that the
`customer` has the means to buy this ticket. The `customer` may have enough
`cash` or a `voucher`

Check if the provided variable `customer` can afford the ticket:

- whether he has enough cash (`customer.cash` property)
- or, if he has a voucher (`customer.hasVoucher` property is `true`)

If so, you must increment the provided variable `ticketSold` value by `1`.
