## hiddenp

### Instructions

Écrire un programme nommé `hiddenp` qui prend deux `strings` et qui, si la première `string` est cachée dans la deuxième, affiche `1`
suivi d'un newline(`'\n'`), autrement il affiche `0` suivi d'un newline.

Considérons s1 et s2 comme étant des `strings`. Il est considéré que s1 est cachée dans s2 si il est possbile de trouver chaque caractère de s1 dans s2, **dans le même ordre d'apparence que s1.**

Si s1 est une `string` vide elle est considérée cachée dans n'importe quelle `string`.

Si le nombre de paramètres est différent de 2, le programme affiche un newline.

### Utilisation

```console
student@ubuntu:~/student/hiddenp$ go build
student@ubuntu:~/student/hiddenp$ ./hiddenp "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
1$
student@ubuntu:~/student/hiddenp$ ./hiddenp "abc" "2altrb53c.sse" | cat -e
1$
student@ubuntu:~/student/hiddenp$ ./hiddenp "abc" "btarc" | cat -e
0$
student@ubuntu:~/student/hiddenp$ ./hiddenp | cat -e
$
student@ubuntu:~/student/hiddenp$
```
