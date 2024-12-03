## looping

### Instructions

Write a **program** that prints a riddle, receives input from the user and checks that the answer is correct.

The program must allow an indefinite number of trials and only quit after the correct answer is given.

Every time the user introduces an incorrect answer the program must print the riddle again and after the user gives the correct answer the program must print the number of tries that took to get the correct answer.

Riddle: I am the beginning of the end, and the end of time and space. I am essential to creation, and I surround every place. What am I?

Answer: The letter e

### Notions

- [Module std::io](https://doc.rust-lang.org/std/io/index.html)
- [loop](https://doc.rust-lang.org/std/keyword.loop.html)

### Usage

```console
$ cargo run
I am the beginning of the end, and the end of time and space. I am essential to creation, and I surround every place. What am I?
I don't know
I am the beginning of the end, and the end of time and space. I am essential to creation, and I surround every place. What am I?
The letter e
Number of trials: 2
$
```
