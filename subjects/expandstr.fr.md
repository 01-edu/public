## expandstr

### Instructions

Écrire un programme qui prend une `string` et qui l'affiche avec exactement 3 espaces entre chaque mot, sans espace ou tabulation ni au début ni à la fin.

La `string` sera suivie d'un newline(`'\n'`).

Un mot est une suite de caractères alphanumériques.

Si le nombre de paramètres n'est pas 1, ou si il n'y a pas de mot, le programme affiche un newline.

### Utilisation

```console
student@ubuntu:~/student/expandstr$ go build
student@ubuntu:~/student/expandstr$ ./expandstr "you   see   it's   easy   to   display   the   same   thing" | cat -e
you   see   it's   easy   to   display   the   same   thing$
student@ubuntu:~/student/expandstr$ ./expandstr "   only  it's harder   " | cat -e
only   it's   harder$
student@ubuntu:~/student/expandstr$ ./expandstr " how funny it is" "did you  hear, Mathilde ?" | cat -e
$
student@ubuntu:~/student/expandstr$ ./expandstr | cat -e
$
student@ubuntu:~/student/expandstr$
```
