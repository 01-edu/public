## Alter Ego 🎭

You can change your objects in multiple ways:

- modify values of properties
- add new properties
- delete properties

### Modifying objects

Let's declare an object

```js
// we create our object with 2 properties
const user = {
  points: 0,
  code: '75lai78wn',
}
```

#### Adding a new property

```js
user.name = 'Martin' // we add a name to our user
```

The syntax is very similar to modifying a variable, the only difference is we
start **from** our user and use the `.` _(property accessor operator)_

#### Changing a property value

```js
user.points = 10 // set points to 10 points
```

The syntax is the same, if the property already exist, its value will be changed
!

#### Removing a property

```js
user.code = undefined // remove the value from a property
```

The trick here is to set its value to `undefined` as this will clear the value
of our property

### Instructions

Modify the provided `alterEgo` variable:

- add a `self` property with the string value `'altered'`.
- add a `fullName` property that is the joined value of the `firstName` and the
  `lastName` with a space in between.
- add `10` to it's `points` property
