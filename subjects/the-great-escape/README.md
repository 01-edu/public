## The Great Escape

### Escaping

**Quote delimiters** can be one of the tricky things to deal with.

Since they are used for delimiting text, they need a trick to include them in
our text.

For example, we want a `'` _(single quote)_ in or text, but use them as
delimiters:

```js
console.log('Houdini once said:')
console.log('Magic is the sole science not accepted by scientists,')
//                    Uh oh... ↙ JS thinks your string ends here
console.log(' because they can't understand it.')
//                       ...and new starts here ↖ that never finish !
// too bad ! a quote, ruined by quotes, ironic and very sad.
```

The `\` _(backslash)_ is used for that:

Everytime there is an _extra special_ character into your string, putting a `\`
in front of it will **escape** it and doing so will let JS understand you meant
the **litteral** following character and not the delimiter, _or whatever else
the character normaly means for a string_

### Instructions

Nothing can stop you now with that new knowledge. Like Houdini, master of
escapes, in this exercise you are going to **escape some strings**:

- Create a `escapeFromDelimiters` that includes all 3 quotes _(`` ` ``, `"` and
  `'`)_.

- Create a `escapeTheEscape` that includes a backslash _(`\`)_.

> “How did I escape? With difficulty. How did I plan this moment? With
> pleasure.” \
> ― Alexandre Dumas, The Count of Monte Cristo
