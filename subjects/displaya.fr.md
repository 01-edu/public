## displaya

### Instructions

Écrire un programme qui prend une `string`, et qui affiche le premier caractère `a` qu'il trouve dedans, suivi par un retour à la ligne (`'\n'`). Si il n'y a pas de caractère `a` dans la `string`, le programme affiche juste un `a` suivi d'un retour à la ligne (`'\n'`). Si le nombre de paramètres n'est pas 1, le programme affiche un `a` suivi d'un retour à la ligne (`'\n'`).

### Utilisation

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test "abc"
a
student@ubuntu:~/piscine-go/test$ ./test "bcvbvA"
a
student@ubuntu:~/piscine-go/test$ ./test "nbv"
a
student@ubuntu:~/piscine-go/test$
```
