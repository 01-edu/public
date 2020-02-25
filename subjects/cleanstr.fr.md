## cleanstr

### Instructions

Écrire un programme qui prend une `string`, et qui affiche cette `string` avec exactement:

- un espace entre les mots.
- aucun espace ni de tabulation ni au début ni à la fin.
- le résultat avecsuivi d'un saut de ligne("`\n`").

Un "mot" est défini comme une partie de `string` délimité par soit des espaces/tabulations soit par le début/fin d'une `string`.

Si le nombre d'arguments est différent de 1, ou si il n'y a pas de mots à afficher, alors le programme affiche un saut de ligne("`\n`").

### Utilisation :

```console
student@ubuntu:~/[[ROOT]]/cleanstr$ go build
student@ubuntu:~/[[ROOT]]/cleanstr$ ./cleanstr "you see it's easy to display the same thing" | cat -e
you see it's easy to display the same thing$
student@ubuntu:~/[[ROOT]]/cleanstr$ ./cleanstr " only    it's  harder   "
only it's harder$
student@ubuntu:~/[[ROOT]]/cleanstr$ ./cleanstr " how funny" "Did you   hear Mathilde ?"
$
student@ubuntu:~/[[ROOT]]/cleanstr$ ./cleanstr "" | cat -e
$
student@ubuntu:~/[[ROOT]]/cleanstr$
```
