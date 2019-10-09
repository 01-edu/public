## compareprog

### Instructions

Écrire un programme qui se comporte comme la fonction `Compare`.

Ce programme affiche un nombre après avoir comparé deux `string` lexicalement.

### Utilisation :

```console
student@ubuntu:~/compareprog$ go build
student@ubuntu:~/compareprog$ ./compareprog a b | cat -e
-1$
student@ubuntu:~/compareprog$ ./compareprog a a | cat -e
0$
student@ubuntu:~/compareprog$ ./compareprog b a | cat -e
1$
student@ubuntu:~/compareprog$ ./compareprog b a d | cat -e
$
student@ubuntu:~/compareprog$ ./compareprog | cat -e
$
student@ubuntu:~/compareprog$
```
