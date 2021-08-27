## cleanstr

### Instructions

Write a **program** that takes a `string`, and displays this `string` with exactly:

- one space between words.
- without spaces nor tabs at the beginning nor at the end.
- with the result followed by a newline ("`\n`").

A "word" is defined as a part of a `string` delimited either by spaces/tabs, or
by the start/end of the `string`.

If the number of arguments is not 1, or if there are no words to display, the
program displays a newline("`\n`").

### Usage :

```console
$ go run . "you see it's easy to display the same thing" | cat -e
you see it's easy to display the same thing$
$ go run . " only    it's  harder   "
only it's harder$
$ go run . " how funny" "Did you   hear Mathilde ?"

$ go run . ""

$
```
