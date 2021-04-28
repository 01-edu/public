## brainfuck

### Instructions

Write a `Brainfuck` interpreter program.
The source code will be given as first parameter.
The code will always be valid, with less than 4096 operations.
`Brainfuck` is a minimalist language. It consists of an array of bytes (in this exercice 2048 bytes) all initialized with zero, and with a pointer to its first byte.

Every operator consists of a single character :

- '>' increment the pointer
- '<' decrement the pointer
- '+' increment the pointed byte
- '-' decrement the pointed byte
- '.' print the pointed byte on standard output
- '[' go to the matching ']' if the pointed byte is 0 (loop start)
- ']' go to the matching '[' if the pointed byte is not 0 (loop end)

Any other character is a comment.

### Usage

```console
$ go run . "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>." | cat -e
Hello World!$
$ go run . "+++++[>++++[>++++H>+++++i<<-]>>>++\n<<<<-]>>--------.>+++++.>." | cat -e
Hi$
$ go run . "++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++." | cat -e
abc$
$ go run .
$
```
