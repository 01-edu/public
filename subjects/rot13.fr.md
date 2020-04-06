## rot13

### Instructions

Écrire un programme qui prend une `string` et qui l'affiche, en remplaçant chacune de ses lettres par la lettre qui est 13 positions plus loin dans l'ordre alphabétique.

- 'z' devient 'm' et 'Z' devient 'M'. Les majuscules restent des majuscules, de même pour les minuscules.

- l'output sera suivi d'un retour à la ligne (`'\n'`).

- Si le nombre d'arguments est différent de 1, le programme affiche un retour à la ligne (`'\n'`).

### Utilisation

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test "abc"
nop
student@ubuntu:~/[[ROOT]]/test$ ./test "hello there"
uryyb gurer
student@ubuntu:~/[[ROOT]]/test$ ./test

student@ubuntu:~/[[ROOT]]/test$
```
