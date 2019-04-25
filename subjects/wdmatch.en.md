
# switchcase

## Instructions

Write a program that takes two strings and checks whether it is possible to write the first string with characters from the second string, while respecting the order in which these characters appear in the second string.

- If it is possible, the program displays the string followed by a `\n`, otherwise it simply displays a `\n`.

- If the number of arguments is not 2, the program displays `\n`.

Example of output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test "faya" "fgvvfdxcacpolhyghbreda"
faya
student@ubuntu:~/piscine/test$ ./test "faya" "fgvvfdxcacpolhyghbred"

student@ubuntu:~/piscine/test$ ./test "error" rrerrrfiiljdfxjyuifrrvcoojh

student@ubuntu:~/piscine/test$ ./test "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
quarante deux

student@ubuntu:~/piscine/test$ ./test

student@ubuntu:~/piscine/test$
```
