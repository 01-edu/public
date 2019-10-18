## expandstr

### Instructions

Écrire un programme qui prend une `string` et qui l'affiche avec exactement 3 espaces entre chaque mot, sans espace ou tabulation ni au début ni à la fin.

La `string` sera suivie d'un retour à la ligne (`'\n'`).

Un mot est une suite de caractères alphanumériques.

Si le nombre de paramètres n'est pas 1, ou si il n'y a pas de mot, le programme affiche un retour à la ligne.

### Utilisation

```console
student@ubuntu:~/piscine-go/expandstr$ go build
student@ubuntu:~/piscine-go/expandstr$ ./expandstr "you   see   it's   easy   to   display   the   same   thing" | cat -e
you   see   it's   easy   to   display   the   same   thing$
student@ubuntu:~/piscine-go/expandstr$ ./expandstr "   only  it's harder   " | cat -e
only   it's   harder$
student@ubuntu:~/piscine-go/expandstr$ ./expandstr " how funny it is" "did you  hear, Mathilde ?" | cat -e
$
student@ubuntu:~/piscine-go/expandstr$ ./expandstr | cat -e
$
student@ubuntu:~/piscine-go/expandstr$
```
