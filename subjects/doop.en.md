## doop

### Instructions

Write a [program](TODO-LINK) that is called `doop`.

The program has to be used with three arguments:

- A value
- An operator
- Another value

In case of an invalid operator the programs prints `0`.

In case of an invalid number of arguments the program prints nothing.

The program has to handle the modulo and division operations by 0 as shown on the output examples below. 

`fmt.Print` is authorized.

Examples of outputs :

```console
student@ubuntu:~/piscine/test$ go build doop.go
student@ubuntu:~/piscine/test$ ./doop
student@ubuntu:~/piscine/test$ ./doop 1 + 1
2
student@ubuntu:~/piscine/test$ ./doop hello + 1 | cat -e
0$
student@ubuntu:~/piscine/test$ ./doop 1 p 1
0
student@ubuntu:~/piscine/test$ ./doop 1 + 1
2
student@ubuntu:~/piscine/test$ ./doop 1 / 0 | cat -e
No division by 0$
student@ubuntu:~/piscine/test$ ./doop 1 % 0 | cat -e
No modulo by 0$
student@ubuntu:~/piscine/test$ ./doop 1 "*" 1
1
```
