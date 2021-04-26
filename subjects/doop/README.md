## doop

### Instructions

Write a program that is called `doop`.

The program has to be used with three arguments:

- A value
- An operator, one of : `+`, `-`, `/`, `*`, `%`
- Another value

In case of an invalid operator, value, number of arguments or an overflow the programs prints nothing.

The program has to handle the modulo and division operations by 0 as shown on the output examples below.

### Usage

```console
student@ubuntu:~/doop/test$ go build doop.go
student@ubuntu:~/doop/test$ ./doop
student@ubuntu:~/doop/test$ ./doop 1 + 1 | cat -e
2$
student@ubuntu:~/doop/test$ ./doop hello + 1
student@ubuntu:~/doop/test$ ./doop 1 p 1
student@ubuntu:~/doop/test$ ./doop 1 / 0 | cat -e
No division by 0$
student@ubuntu:~/doop/test$ ./doop 1 % 0 | cat -e
No modulo by 0$
student@ubuntu:~/doop/test$ ./doop 9223372036854775807 + 1
0
student@ubuntu:~/doop/test$ ./doop -9223372036854775809 - 3
0
student@ubuntu:~/doop/test$ ./doop 9223372036854775807 "*" 3
0
student@ubuntu:~/doop/test$ ./doop 1 "*" 1
1
student@ubuntu:~/doop/test$ ./doop 1 "*" -1
-1
student@ubuntu:~/doop/test$
```
