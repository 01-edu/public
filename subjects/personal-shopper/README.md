## personal-shopper

### Instructions

You know your guests, but don't forget you have to feed them if you want to be
considered as a good host.

Create a `personal-shopper.mjs` script that:

- Takes a file as first argument, for example `shopping-list.json`
- Takes one of these keywords as second argument:
  - `create`: create the file
  - `delete`: delete the file
  - `add`: add a new element to the list in the file.
    - This command line must take a third argument which is the name of the new
      entry in your list. If no third argument is passed, console must print
      this error: `No elem specified.`.
    - This command line could take a fourth argument which is the number of
      elements you want for this new entry.
      - If there is no 4rth argument or the 4rth argument is `NaN`, 1 would be
        the value by default.
      - If the entry already exists, it would add 1 or more to the original
        value.
      - Using a negative number must behave as `rm` command.
  - `rm`: remove an element from the list in the file
    - This command line must take a third argument which is the name of the
      entry to remove from the list.
      - If no third argument is passed, console must print this error:
        `No elem specified.`.
      - If the entry does not exists, it does nothing.
    - This command line could take a fourth argument which is the number of
      elements you want to delete from this entry.
      - If there is no 4rth argument: it remove the entry.
      - If the 4rth argument is `NaN`, nothing is removed and console must print
        this error `Unexpected request: nothing has been removed`.
      - If the 4rth argument is a number, it will substract this number from the
        original value (if the new value is <= 0, it will remove the entry).
      - Using a negative number must behave as `add` command.
  - `help`: print all the command lines available, with a description of it
    (specifications in the examples)
  - `ls` or no more arguments: print the list in the console.
    - Each line is formated like this: `- element (number)`
    - If the list is empty, this message should appear in console:
      `Empty list.`.

If no keyword is passed as second argument, the helper should be printed in the console.

#### Examples

- `node personal-shopper.mjs shopping-list.json create` would create the file
- `node personal-shopper.mjs shopping-list.json delete` would remove the file

- `node personal-shopper.mjs shopping-list.json add "tzatziki pot"` would update
  the content of the file like:

```json
'{
  "tzatziki pot": 1
}'
```

- `node personal-shopper.mjs shopping-list.json add carrots 5` would update the
  content of the file like:

```json
'{
  "tzatziki pot": 1,
  "carrots": 5
}'
```

- `node personal-shopper.mjs shopping-list.json add carrots 2` would update the
  content of the file like:

```json
'{
  "tzatziki pot": 1,
  "carrots": 7
}'
```

- `node personal-shopper.mjs shopping-list.json rm carrots 4` would update the
  content of the file like:

```json
'{
  "tzatziki pot": 1,
  "carrots": 3
}'
```

- `node personal-shopper.mjs shopping-list.json rm carrots` would update the
  content of the file like:

```json
'{
  "tzatziki pot": 1
}'
```

- `node personal-shopper.mjs shopping-list.json rm carrots -3` would update the
  content of the file like:

```json
'{
  "tzatziki pot": 1,
  "carrots": 3
}'
```

- `node personal-shopper.mjs shopping-list.json add carrots -3` would update the
  content of the file like:

```json
'{
  "tzatziki pot": 1
}'
```

- `node personal-shopper.mjs shopping-list.json ls` would print the list in your
  console like this:

```
- tzatziki pot (1)
```

- `node personal-shopper.mjs help` would print a list of all available commands
  in your console :

```
Commands:
- create: takes a filename as argument and create it (should have `.json` extension specified)
- delete: takes a filename as argument and delete it
<!-- etc. -->
```

### Notions

- [Node file system: `rm`](https://nodejs.org/docs/latest/api/fs.html#fs_fspromises_rm_path_options)
- [Node file system: `writeFile`](https://nodejs.org/docs/latest/api/fs.html#fs_fspromises_writefile_file_data_options)
- [`JSON.parse()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON/parse)
- [`JSON.stringify()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON/stringify)
- [`isNaN()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/isNaN)
- [`Number()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number)
- [`console.error()`](https://developer.mozilla.org/en-US/docs/Web/API/Console/error)
