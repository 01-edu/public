## revWstr

### Instructions

Write a program that takes a string as a parameter, and prints its words in reverse.

- A "word" is a part of the string bounded by spaces and/or tabs, or the begin/end of the string.

- If the number of parameters is different from 1, the program will display `\n`.

- In the parameters that are going to be tested, there won't be any additional" spaces (meaning that there won't be additionnal spaces at the beginning or at the end of the string, and words will always be separated by exactly one space).

Example of output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test "the time of contempt precedes that of indifference"
indifference of that precedes contempt of time the
student@ubuntu:~/piscine/test$ ./test "abcdefghijklm"
abcdefghijklm
student@ubuntu:~/piscine/test$ ./test "he stared at the mountain"
mountain the at stared he
student@ubuntu:~/piscine/test$ ./test

student@ubuntu:~/piscine/test$
```