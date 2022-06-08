## brainfuck

> Esoteric programming languages (esolang) are designed to test the boundaries of computer programming design, as proofs of concept, as software art, as a hacking interface or simply as a joke. One such esoteric language is `Brainfuck`. It was created by Urban Müller in 1993. It is a minimalist language consisting of 8 simple commands. It is Turing complete, but is not intended for practical use. It exists to amuse and challenge programmers.

### Instructions

Write a `Brainfuck` interpreter program.

The source code will be given as the first parameter, and will always be valid with fewer than 4096 operations.

Your `Brainfuck` interpreter will consist of an array of 2048 bytes, all initialized to 0, with a pointer to the first byte.

Every operator consists of a single character:
- `>`: increment the pointer.
- `<`: decrement the pointer.
- `+`: increment the pointed byte.
- `-`: decrement the pointed byte.
- `.`: print the pointed byte to the standard output.
- `[`: if the pointed byte is 0, then instead of moving onto the next command, skip to the command after the matching `]`.
- `]`: if the pointed byte is **not** 0, then instead of moving onto the next command, move back to the command after the matching `[`.

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
