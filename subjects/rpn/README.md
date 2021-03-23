## rpn

### Instructions

Write a program that takes a `string` which contains an equation written in `Reverse Polish Notation` (RPN) as its first argument,
that evaluates the equation, and that prints the result on the standard output followed by a newline (`'\n'`).

`Reverse Polish Notation` is a mathematical notation in which every operator follows all of its operands. In RPN,
every operator encountered evaluates the previous 2 operands, and the result of this operation then becomes the first of
the two operands for the subsequent operator. Operands and operators must be spaced by at least one space.

The following operators must be implemented : `+`, `-`, `*`, `/`, and `%`.

If the `string` is not valid or if there is not exactly one argument, `Error` must be printed on the standard output followed by a newline.
If the `string` has extra spaces it is still considered valid.

All the given operands must fit in a `int`.

Examples of formulas converted in RPN:

3 + 4 >> 3 4 +

((1 \* 2) \* 3) - 4 >> 1 2 \* 3 \* 4 - or 3 1 2 \* \* 4 -

50 \* (5 - (10 / 9)) >> 5 10 9 / - 50 \*

Here is how to evaluate a formula in RPN:

```console
1 2 * 3 * 4 -
2 3 * 4 -
6 4 -
2
```

Or:

```console
3 1 2 * * 4 -
3 2 * 4 -
6 4 -
2
```

For receiving arguments from the command line you should use something like:

```rust
    fn main() {
        let args: Vec<String> = std::env::args().collect();

        rpn(&args[1]);
    }

```

### Usage

````console
$ cargo run "1 2 * 3 * 4 +"
10
$ cargo run "1 2 3 4 +"
Error
$ cargo run
Error
$ cargo run "     1      3 * 2 -"
1
$ cargo run "     1      3 * ksd 2 -"
Error
```
