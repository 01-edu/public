## paramcount

### Instructions

Écrire un programme qui affiche le nombre d'arguments passés à ce programme. Ce nombre sera suivi d'un retour à la ligne (`'\n'`).

Si il n'y a pas d'argument, le programme affiche `0` suivi d'un retour à la ligne.

### Utilisation

```console
student@ubuntu:~/piscine-go/paramcount$ go build
student@ubuntu:~/piscine-go/paramcount$ ./paramcount 1 2 3 5 7 24
6
student@ubuntu:~/piscine-go/paramcount$ ./paramcount 6 12 24 | cat -e
3$
student@ubuntu:~/piscine-go/paramcount$ ./paramcount | cat -e
0$
student@ubuntu:~/piscine-go/paramcount$
```
