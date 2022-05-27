## brackets_matching

### Instructions

Write a `program` that takes an undefined number of `string` in arguments. For each argument, if the expression is correctly bracketed, the program prints on the standard output `OK` followed by a newline (`'\n'`), otherwise it prints `Error` followed by a newline.

Symbols considered as brackets are parentheses `(` and `)`, square brackets `[` and `]` and curly braces `{` and `}`. Every other symbols are simply ignored.

An opening bracket must always be closed by the good closing bracket in the correct order. A `string` which does not contain any bracket is considered as a correctly bracketed.

If there is no argument, the program must print nothing.

For receiving arguments from the command line you should use something like:

```rust
fn main() {
    let args: Vec<String> = std::env::args().collect();

    //...
}

```

### Usage

```console
$ cargo run '(johndoe)' | cat -e
OK
$ cargo run '([)]' | cat -e
Error
$ cargo run

$ cargo run '' '{[(0 + 0)(1 + 1)](3*(-1)){()}}' | cat -e
OK
OK
```
