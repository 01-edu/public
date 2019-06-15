## rot13

### Instructions

Écrire un programme qui prend une `string` et qui l'affiche, en remplaçant chacune de ses lettres par la lettre 13 positions après dans l'ordre alphabétique.

- 'z' devient 'm' et 'Z' devient 'M'. Les majuscules restent des majuscules, de même pour les minuscules.

- l'output sera suivi d'un newline(`'\n'`).

- Si le nombre d'arguments n'est pas 1, le programme affiche un newline(`'\n'`).

### Utilisation

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test "abc"
nop
student@ubuntu:~/piscine/test$ ./test "hello there"
uryyb gurer
student@ubuntu:~/piscine/test$ ./test

student@ubuntu:~/piscine/test$
```
