## rostring

### Instructions

Écrire un programme qui prend une `string` et affiche cette `string` après avoir déplacé un mot vers la gauche.

Le premier mot devient donc le dernier, et le reste des mots reste dans le même ordre.

Un mot est une suite de caractères **alphanumériques.**

Les mots devront être séparés par un seul espace dans l'output.

Si le nombre d'arguments est différent de 1, le programme affiche un retour à la ligne.

### Utilisation

```console
student@ubuntu:~/piscine-go/rostring$ go build
student@ubuntu:~/piscine-go/rostring$ ./rostring "abc   " | cat -e
abc$
student@ubuntu:~/piscine-go/rostring$ ./rostring "Let there     be light"
there be light Let
student@ubuntu:~/piscine-go/rostring$ ./rostring "     AkjhZ zLKIJz , 23y"
zLKIJz , 23y AkjhZ
student@ubuntu:~/piscine-go/rostring$ ./rostring | cat -e
$
student@ubuntu:~/piscine-go/rostring$
```
