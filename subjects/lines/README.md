## Lines

### The new line character `'\n'`

One other special characters in strings is the `\n` (new line), we use it to
represent a new line without having to have a new line in our string.

Also the single and double quote delimited strings can't actually have literal
new lines:

```js
let fewlines = '\nHello👏\nThere👏\n'
// Same string but with litteral new lines:
let usingLiteral = `
Hello👏
There👏
`
```

You can use the literal `\n` character to split text on each lines:

```js
let splited = `
Hello👏
There👏
`.split('\n')

console.log(splited) // ['','Hello👏','There👏', '']
// Note that empty lines becomes empty strings !
```

### Instructions

You have been recruted to analyse a bunch of poems, your first task is to count
the number of lines.

- Declare a variable `linesCount` of the number of lines from the provided
  `poem` variable.

But you must ignore empty lines in the begining and the end using the `trim`
string method.

### Notions

- [devdocs.io/javascript/global_objects/string/trim](https://devdocs.io/javascript/global_objects/string/trim)
