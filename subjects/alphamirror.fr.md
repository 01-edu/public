## alphamirror

### Instructions

Écrire un programme nommé `alphamirror` qui prend une `string` comme argument et qui affiche cette `string` après remplacement de chaque caractère alphabétique avec son opposé.

Les majuscules restent des majuscules, de même pour le minuscules, par exemple :

'a' devient 'z', 'Z' devient 'A'
'd' devient 'w', 'M' devient 'N'

Le résultat final sera suivi d'un retour à la ligne (`'\n'`).

Si le nombre d'arguments est différent de 1, le programme affiche seulement un retour à la ligne (`'\n'`).

### Utilisation

```console
student@ubuntu:~/piscine-go/alphamirror$ go build
student@ubuntu:~/piscine-go/alphamirror$ ./alphamirror "abc"
zyx
student@ubuntu:~/piscine-go/alphamirror$ ./alphamirror "My horse is Amazing." | cat -e
Nb slihv rh Znzarmt.$
student@ubuntu:~/piscine-go/alphamirror$ ./alphamirror | cat -e
$
student@ubuntu:~/piscine-go/alphamirror$
```
