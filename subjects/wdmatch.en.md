## wdmatch

### Instructions

Write a program that takes two `string` and checks whether it is possible to write the first `string` with characters from the second `string`, while respecting the order in which these characters appear in the second `string`.

If it is possible, the program displays the `string` followed by a newline (`'\n'`), otherwise it simply displays a newline.

If the number of arguments is different from 2, the program displays a newline.

### Usage

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test "faya" "fgvvfdxcacpolhyghbreda"
faya
student@ubuntu:~/piscine-go/test$ ./test "faya" "fgvvfdxcacpolhyghbred"

student@ubuntu:~/piscine-go/test$ ./test "error" rrerrrfiiljdfxjyuifrrvcoojh

student@ubuntu:~/piscine-go/test$ ./test "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
quarante deux

student@ubuntu:~/piscine-go/test$ ./test

student@ubuntu:~/piscine-go/test$
```
