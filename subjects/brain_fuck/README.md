## brain_fuck

### Instructions

Write a `Brainfuck` interpreter program.
The source code will be given as the first argument.
The code will always be valid.
`Brainfuck` is a minimalist language. It consists of an array of bytes (in this exercise as many as 2048 bytes) all initialized with zero, and with a pointer to its first byte.

Every operator consists of a single character :

- '>' increment the pointer
- '<' decrement the pointer
- '+' increment the pointed byte
- '-' decrement the pointed byte
- '.' print the pointed byte on standard output
- '[' go to the matching ']' if the pointed byte is 0 (loop start)
- ']' go to the matching '[' if the pointed byte is not 0 (loop end)

Any other character is a comment.

> Use `std::env::args()` to get the program's arguments.

### Usage

```console
$ cargo run "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
Hello World!
$ cargo run "+++++[>++++[>++++H>+++++i<<-]>>>++\n<<<<-]>>--------.>+++++.>."
Hi
$ cargo run "++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++."
abc
```
