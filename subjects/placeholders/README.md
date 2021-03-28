## Placeholders

Enough about numbers, let's check out what we can do with strings !

### Strings `` `${placeholders}` ``

The first one are `placeholders` for using values inside our strings, they have
there own syntax: `${}`.

#### Example

```js
console.log(`5 + 10 = ${5 + 10} = 15`) // -> 5 + 10 = 15 = 15
```

**Note that it only works using:** the `` ` `` backtick, not the `"` or `'`
quotes.

```js
// here, with the use of quotes instead of backticks, the placeholder is not evaluated and we see it as text:
console.log('5 + 10 = ${5 + 10} = 15') // -> 5 + 10 = ${5 + 10} = 15
```

### Instructions

We will provide a variable `name` and `age`. They will be pre-declared by us.

Declare a `presentation` variable of the string:

> `Hello, my name is` **name** `and I'm` **age** `years old`

But use placeholders to build the string you will put inside `presentation`.  
Put the values of the provided variables `age` and `name` inside those
placeholders.
