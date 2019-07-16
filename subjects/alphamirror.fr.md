## alphamirror

### Instructions

Écrire un programme nommé `alphamirror` qui prend une `string` come argument et qui affiche cette `string` après remplacement de chaque caractère alphabétique avec son opposé.

Les majuscules restent des majuscules, de même pour le minuscules, par exemple :

'a' devient 'z', 'Z' devient 'A'
'd' devient 'w', 'M' devient 'N'

Le résultat final sera suivi d'un newline(`'\n'`).

Si le nombre d'arguments est différent de 1, le programme affiche seulement un newline(`'\n'`).

### Utilisation

```console
student@ubuntu:~/student/alphamirror$ go build
student@ubuntu:~/student/alphamirror$ ./alphamirror "abc"
zyx
student@ubuntu:~/student/alphamirror$ ./alphamirror "My horse is Amazing." | cat -e
Nb slihv rh Znzarmt.$
student@ubuntu:~/student/alphamirror$ ./alphamirror | cat -e
$
student@ubuntu:~/student/alphamirror$
```
