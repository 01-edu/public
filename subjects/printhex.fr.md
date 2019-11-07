## printhex

### Instructions

Écrire un programme qui prend un nombre positif (ou zéro) écrit en base 10, et qui l'affiche en base 16 (avec les lettres en minuscule) suivi d'un retour à la ligne (`'\n'`).

-   Si le nombre de paramètres est différent de 1, le programme affiche un retour à la ligne.
-   Les cas d'erreurs doivent être gérés comme montré dans l'exemple ci-dessous.


### Utilisation

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test "10"
a
student@ubuntu:~/piscine-go/test$ ./test "255"
ff
student@ubuntu:~/piscine-go/test$ ./test "5156454"
4eae66
student@ubuntu:~/piscine-go/test$ ./test

student@ubuntu:~/piscine-go/test$ ./test "123 132 1" | cat -e
0$
student@ubuntu:~/piscine-go/test$ 
```
