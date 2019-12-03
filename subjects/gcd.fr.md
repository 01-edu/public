## gcd

### Instructions

Écrire un programme qui prend deux `string` représentant deux entiers strictement positifs qui rentrent dans un `int`.

Afficher leur plus grand diviseur commun suivi d'un retour à la ligne (`'\n'`).

Si le nombre de paramètres est différent de 2, le programme affiche un retour à la ligne.

Tous les arguments testés seront des `int` positifs.

### Utilisation

```console
student@ubuntu:~/[[ROOT]]/gcd$ go build
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 42 10 | cat -e
2$
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 42 12 | cat -e
6$
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 14 77 | cat -e
7$
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 17 3 | cat -e
1$
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd | cat -e
$
student@ubuntu:~/[[ROOT]]/gcd$ ./gcd 50 12 4 | cat -e
$
student@ubuntu:~/[[ROOT]]/gcd$
```
