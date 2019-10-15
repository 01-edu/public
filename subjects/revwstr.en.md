## revWstr

### Instructions

Write a program that takes a `string` as a parameter, and prints its words in reverse.

-   A word is a sequence of **alphanumerical** characters.

-   If the number of parameters is different from 1, the program will display newline (`'\n'`).

-   In the parameters that are going to be tested, there will not be any extra spaces. (meaning that there will not be additionnal spaces at the beginning or at the end of the `string`and that words will always be separated by exactly one space).

### Usage

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test "the time of contempt precedes that of indifference"
indifference of that precedes contempt of time the
student@ubuntu:~/piscine-go/test$ ./test "abcdefghijklm"
abcdefghijklm
student@ubuntu:~/piscine-go/test$ ./test "he stared at the mountain"
mountain the at stared he
student@ubuntu:~/piscine-go/test$ ./test

student@ubuntu:~/piscine-go/test$
```
