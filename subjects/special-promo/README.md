## Special Promo 🎟️

### `&&` (the AND operator)

This operator is used to group conditions:

```js
if (user.role === 'admin' && user.available === 'now') {
  console.log('Admin', user.name, 'is available and will help you now !')
}
```

Here, the code will only show the message if **both** conditions are true.

### Instructions

Your traveling company have a new special promo for members between 18 and 25
years old.

Write the `if` condition that will check if the user can benefit the promotion:

- `user.age` must be at least `18`
- `user.age` must be lesser or equal to `25`
- `user.activeMembership` must be `true`

If all of those conditions are true, `console.log` the message
`You can benefit from our special promotion`
