## fprime

### Instructions

Écrire un programme qui prend un `int` positif et qui affiche ses facteurs premiers sur la sortie standard, suivi d'un retour à la ligne (`'\n'`).

-   Les facteurs doivent être affichés en ordre croissant et séparés par `*`.

-   Si le nombre de paramètres est différent de 1, le programme affiche un retour à la ligne.

-   L'input (l'entrée), quand il y en a un, sera toujours valide.

### Utilisation

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test 225225
3*3*5*5*7*11*13
student@ubuntu:~/piscine-go/test$ ./test 8333325
3*3*5*5*7*11*13*37
student@ubuntu:~/piscine-go/test$ ./test 9539
9539
student@ubuntu:~/piscine-go/test$ ./test 804577
804577
student@ubuntu:~/piscine-go/test$ ./test 42
2*3*7
student@ubuntu:~/piscine-go/test$ ./test

student@ubuntu:~/piscine-go/test$
```
