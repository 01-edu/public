## hiddenp

### Instructions

Écrire un programme nommé `hiddenp` qui prend deux `string` et qui, si la première `string` est cachée dans la deuxième, affiche `1` suivi d'un retour à la ligne (`'\n'`), autrement il affiche `0` suivi d'un retour à la ligne.

Considérons s1 et s2 comme étant des `string`. Il est considéré que s1 est cachée dans s2 si il est possbile de trouver chaque caractère de s1 dans s2, **dans le même ordre d'apparence que s1.**

Si s1 est une `string` vide elle est considérée cachée dans n'importe quelle `string`.

Si le nombre de paramètres est différent de 2, le programme affiche un retour à la ligne.

### Utilisation

```console
student@ubuntu:~/piscine-go/hiddenp$ go build
student@ubuntu:~/piscine-go/hiddenp$ ./hiddenp "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
1$
student@ubuntu:~/piscine-go/hiddenp$ ./hiddenp "abc" "2altrb53c.sse" | cat -e
1$
student@ubuntu:~/piscine-go/hiddenp$ ./hiddenp "abc" "btarc" | cat -e
0$
student@ubuntu:~/piscine-go/hiddenp$ ./hiddenp | cat -e
$
student@ubuntu:~/piscine-go/hiddenp$
```
