## union

### Instructions

Écrire un programme qui prend deux `string` et affiche, sans doublons, les caractères qui apparaissent dans l'une ou l'autre des `string`.

L'affichage se fera dans l'ordre d'apparition des caractères de la ligne de commande, et il sera suivi d'un retour à la ligne (`'\n'`).

Si le nombre d'arguments est différent de 2, le programme affiche un retour à la ligne (`'\n'`).

### Utilisation

```console
student@ubuntu:~/[[ROOT]]/union$ go build
student@ubuntu:~/[[ROOT]]/union$ ./union zpadinton "paqefwtdjetyiytjneytjoeyjnejeyj" | cat -e
zpadintoqefwjy$
student@ubuntu:~/[[ROOT]]/union$ ./union ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
df6vewg4thras$
student@ubuntu:~/[[ROOT]]/union$ ./union "rien" "cette phrase ne cache rien" | cat -e
rienct phas$
student@ubuntu:~/[[ROOT]]/union$ ./union | cat -e
$
student@ubuntu:~/[[ROOT]]/union$ ./union "rien" | cat -e
$
student@ubuntu:~/[[ROOT]]/union$
```
