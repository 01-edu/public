## wdmatch

### Instructions

Write a program that takes two `string` and checks whether it is possible to write the first `string` with characters from the second `string`, while respecting the order in which these characters appear in the second `string`.

If it is possible, the program displays the `string` followed by a newline (`'\n'`), otherwise it simply displays nothing.

If the number of arguments is different from 2, the program displays nothing.

### Usage

```console
student@ubuntu:~/wdmatch/test$ go build
student@ubuntu:~/wdmatch/test$ ./test faya fgvvfdxcacpolhyghbreda
faya
student@ubuntu:~/wdmatch/test$ ./test faya fgvvfdxcacpolhyghbred
student@ubuntu:~/wdmatch/test$ ./test error rrerrrfiiljdfxjyuifrrvcoojh
student@ubuntu:~/wdmatch/test$ ./test "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
quarante deux
student@ubuntu:~/wdmatch/test$ ./test
student@ubuntu:~/wdmatch/test$
```
