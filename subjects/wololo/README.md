## 🌟 Wololo

### Converting types

As you now know, the number `42` is different than the string `'42'`, but if we
write `` `${42}` `` we converted the number to a string !

We see it because of the delimiters, but it is also true for the memory in your
computer.

For example we can not multiply strings, if you try to do `'hello' * 2` or
`'hello' * 'hello'` you will have an unexpected result.

> Well what were you expecting really ? `'hellohello'` maybe ?

So sometimes it is usefull to go from strings to number to boolean _and back !_

- `Number` is a function to convert to a number.
- `Boolean` is a function to convert to a boolean.
- `String` is a function to convert to, you guessed it, a string.

So to convert a `boolean` to a `string` we would write:

```js
String(true)
```

One other way we can use `placeholders` for, is to convert from any values to a
string, but using functions is more clear than abusing placeholder syntax:

```js
let str42Placeholder = `${42}` // was this a mistake ?
let str42Function = String(42) // ah okey we want a string !
```

> so there you have it, calling, arguments and return values, let's see you
> apply all of that now

### Instructions

For this exercise, we provide 3 variables `num`, `bool` and `str` of a matching
type.

Using the magical power of functions, execute the following conversions:

- a `stringFromNumber` variable of the converted value of `num` to a `string`
- a `stringFromBoolean` variable of the converted value of `bool` to a `string`
- a `numberFromString` variable of the converted value of `str` to a `number`
- a `numberFromBoolean` variable of the converted value of `bool` to a `number`
- a `booleanFromString` variable of the converted value of `str` to a `boolean`
- a `booleanFromNumber` variable of the converted value of `num` to a `boolean`
