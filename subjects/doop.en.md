## doop

### Instructions

Write a program that is called `doop`.

The program has to be used with three arguments:

- A value
- An operator
- Another value

You should use `Int64`.

The following operators are considered valid: "+", "-", "/", "*", "%".

In case of an invalid operator the programs prints `0`.

In case of an invalid number of arguments the program prints nothing.

In case of overflow the program prints **Overflow**.

The program has to handle the modulo and division operations by 0 as shown on the output examples below.

`fmt.Print` is authorized.

### Usage

```console
student@ubuntu:~/piscine-go/test$ go build doop.go
student@ubuntu:~/piscine-go/test$ ./doop
student@ubuntu:~/piscine-go/test$ ./doop 1 + 1 | cat -e
2$
student@ubuntu:~/piscine-go/test$ ./doop hello + 1 | cat -e
0$
student@ubuntu:~/piscine-go/test$ ./doop 1 p 1 | cat -e
0$
student@ubuntu:~/piscine-go/test$ ./doop 1 / 0 | cat -e
No division by 0$
student@ubuntu:~/piscine-go/test$ ./doop 1 % 0 | cat -e
No Modulo by 0$
student@ubuntu:~/piscine-go/test$ ./doop 9223372036854775807 + 1
Overflow
student@ubuntu:~/piscine-go/test$ ./doop -9223372036854775809 "*" 3
Overflow
student@ubuntu:~/piscine-go/test$ ./doop 9223372036854775807 "*" 3
Overflow
student@ubuntu:~/piscine-go/test$ ./doop 1 "*" 1
1
student@ubuntu:~/piscine-go/test$ ./doop 1 "*" -1
-1
```
