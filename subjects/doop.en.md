# doop

## Instructions

Write a [program](TODO-LINK) that is called `doop`.

The program has to be used with three arguments:

- A value
- An operator
- Another value

In case of an invalid entry the programs prints `0`.

In case of an invalid number of arguments the program prints nothing.

`fmt.Print` is authorized.

## Usage

```console
student@ubuntu:~/piscine/test$ go build doop.go
student@ubuntu:~/piscine/test$ ./doop
student@ubuntu:~/piscine/test$ ./doop 1 + 1
2
student@ubuntu:~/piscine/test$ ./doop hello + 1
0
student@ubuntu:~/piscine/test$ ./doop 1 p 1
0
student@ubuntu:~/piscine/test$ ./doop 1 + 1
2
student@ubuntu:~/piscine/test$ ./doop 1 / 0
No division by 0
student@ubuntu:~/piscine/test$ ./doop 1 % 0
No modulo by 0
student@ubuntu:~/piscine/test$ ./doop 1 * 1
1

```
