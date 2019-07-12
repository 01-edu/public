## wdmatch

### Instructions

Écrire un programme qui prend deux `strings` et vérifie si il est possible d'écrire la première `string` avec des caractères de la deuxième `string`, tout en respectant l'ordre dans lequel ces caractères apparaissent dans la deuxième `string`.

- Si cela est possible, le programme affiche la `string` suivie par un newline(`'\n'`), autrement le programme affiche un newline.

- Si le nombre d'arguments est différent de 2, le programme affiche un newline.

### Utilisation

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test "faya" "fgvvfdxcacpolhyghbreda"
faya
student@ubuntu:~/piscine/test$ ./test "faya" "fgvvfdxcacpolhyghbred"

student@ubuntu:~/piscine/test$ ./test "error" rrerrrfiiljdfxjyuifrrvcoojh

student@ubuntu:~/piscine/test$ ./test "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
quarante deux

student@ubuntu:~/piscine/test$ ./test

student@ubuntu:~/piscine/test$
```
