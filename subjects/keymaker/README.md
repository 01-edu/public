## Keymaker

![keymaker](https://user-images.githubusercontent.com/231748/112028342-f8461b80-8b2f-11eb-81da-f959cd146770.jpg)

### String as keys

So far we only used `identifiers` as keys for our `Objects`, but they are more
flexible than variables as you can use **any** strings you want !

#### Declaring objects with `strings` as `keys`

When we want to use `Strings` that is not a valid `identifier` we must add
string delimiters:

```js
let usingIdentifier = { valid: true }
let usingString = { valid: true } // they are the same
```

But this syntax allow us more flexibility, if we need it:

```js
let usingSpace = {
  'Still valid !': true,
  "double quotes works too": 1337,

  // but be careful:
  `backtick quote DO NOT WORK !`: false,
  // because of placeholders, they can not be used as easly :(
  // we will see a way to make that work later on :)
}
```

Has we can see, here in usingSpace, I have 2 spaces and a special character, but
that still works, handy !

#### Accessing properties with `strings` as `keys`

And the most usefull one is actually to access properties values using strings
as keys, for that, much like when we wanted to use a number to access a property
in our arrays, we use `[value]` _(square brackets)_:

```js
let usingSpace = { 'Still valid !': true }

console.log(usingSpace['Still valid !']) // true !
```

#### Assigning a properties value with `strings` as `keys`

Much like arrays too, no surprise here:

```js
usingSpace['New key form string'] = 55
```

It's the same old recipe, we access + use the `=` _(assign operator)_ to set the
value.

### Instructions

Now that your are a true **Keymaker** you have a few tasks to do to assert your
power.

- declare an object `whiteRabbit` with a key that is the rabbit emoji 🐰 with
  the value of your choosing
- declare a variable `traitor` that contains the value from the property
  `'Mr. Reagan'` of the provided `secretData` object

> Another way. Always another way. \
> ― The Keymaker
