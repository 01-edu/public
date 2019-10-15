## splitprog

### Instructions

Write a program which separates the words of a `string`, which puts them in a `string` array and which then prints it to standard output.

The program receives two parameters:

-   The first is the `string`
-   The second is the separator

### Usage :

```console
student@ubuntu:~/piscine-go/splitprog$ go build
student@ubuntu:~/piscine-go/splitprog$ ./splitprog "HelloHAhowHAareHAyou?" HA | cat -e
[Hello how are you?]$
student@ubuntu:~/piscine-go/splitprog$ ./splitprog "Hello,how,are,you?" ","
[Hello how are you?]
student@ubuntu:~/piscine-go/splitprog$ ./splitprog "HelloHAhowHAareHAyou?"
student@ubuntu:~/piscine-go/splitprog$
student@ubuntu:~/piscine-go/splitprog$ ./splitprog
student@ubuntu:~/piscine-go/splitprog$
```
