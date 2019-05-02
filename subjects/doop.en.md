## doop

### Instructions

Write a program that takes three strings:

- The first and the third one are representations of base-10 signed integers that fit in an int.

- The second one is an arithmetic operator chosen from: `+ - * / %`

- The program must display the result of the requested arithmetic operation, followed by a newline. If the number of parameters is not 3, the program just displays a newline.

- During tests the strings will have no mistakes or extraneous characters. Negative numbers, in input or output, will have one and only one leading `-`. The result of the operation will fit in an int.

Examples of outputs :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test "123" "*" 456
56088
student@ubuntu:~/piscine/test$ ./test "9828" "/" 234
42
student@ubuntu:~/piscine/test$ ./test "10" "+" "-43"
33
student@ubuntu:~/piscine/test$ ./test

student@ubuntu:~/piscine/test$
```
