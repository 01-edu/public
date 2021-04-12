## 🌟 My House 🏠

One should be aware that copying an objet in Js is not a straightfoward process.
Let's try with an object `mainHouse`. Given what we know, the first instinct
would be to do something like this:

```js
let mainHouse = {
  door: 'blue',
  rooms: {
    bedrooms: 2,
    bathrooms: 1,
  },
}

let sameHouse = mainHouse
```

If we `console.log` both `mainHouse` and `sameHouse` we will obtain exactly the
same output:

```js
console.log(mainHouse) // { door: 'blue', rooms: 3 }
console.log(sameHouse) // { door: 'blue', rooms: 3 }
```

So far, so good. Now, let's get into the issue. Let's paint the door of
`mainHouse` in `red`.

```js
mainHouse.door = 'red'
```

If we `console.log` both `mainHouse` and `sameHouse` we will now obtain:

```js
console.log(mainHouse) // { door: 'red', rooms: 3}
console.log(sameHouse) // { door: 'red', rooms: 3}
```

But wait a second... we just wanted to paint the door of `mainHouse`. What
happened here? Why are both doors painted in red?

Well, unlike primitives that can not changes, _(the number `1` will always be
the number `1`)_ objects can mutate.

You can add, delete and update properties, so if two different variables are
assigned to the same object value, they will both mutate the same object.

We call that a reference, not a copy.

There are multiple ways to copies objects. Understanding the limitations of each
of these ways is very important. Which is why this time, you will have to do a
bit of research...

### Instructions

We will declare the same `mainHouse` seen in the lesson above.

- Declare `acidHouse` which will be a `shallow copy` of `mainHouse`.
- Declare `deepHouse` which will be a `deep copy` of `mainHouse`.

> Not in my house! \
> ― Dikembe Mutombo
