## revwstr

### Instructions

Write a program that takes a `string` as a parameter, and prints its words in reverse, followed by a newline.

- A word is a sequence of **alphanumerical** characters.

- If the number of arguments is different from 1, the program will display nothing.

- In the parameters that are going to be tested, there will not be any extra spaces. (meaning that there will not be additional spaces at the beginning or at the end of the `string` and that words will always be separated by exactly one space).

### Usage

```console
$ go run . "the time of contempt precedes that of indifference"
indifference of that precedes contempt of time the
$ go run . "abcdefghijklm"
abcdefghijklm
$ go run . "he stared at the mountain"
mountain the at stared he
$ go run . "" | cat-e
$
$
```
