## romannumbers

### Instructions

Write a program called `rn`. The objective is to converte a number, given has argument, into a roman number and print it with roman number calculation.

The program should have a limit of `4000`. In case of an invalid number, for example `"hello"` or  `0` the program should print `ERROR: can not convert to roman digit`.

## Usage

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./rn hello
ERROR: can not convert to roman digit
student@ubuntu:~/[[ROOT]]/test$ ./rn 123
C+X+X+I+I+I
CXXIII
student@ubuntu:~/[[ROOT]]/test$ ./rn 999
(M-C)+(C-X)+(X-I)
CMXCIX
student@ubuntu:~/[[ROOT]]/test$ ./rn 3999
M+M+M+(M-C)+(C-X)+(X-I)
MMMCMXCIX
student@ubuntu:~/[[ROOT]]/test$ ./rn 4000
ERROR: can not convert to roman digit
student@ubuntu:~/[[ROOT]]/test$
```
