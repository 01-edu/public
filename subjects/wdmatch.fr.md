## wdmatch

### Instructions

Écrire un programme qui prend deux `string` et vérifie si il est possible d'écrire la première `string` avec des caractères de la deuxième `string`, tout en respectant l'ordre dans lequel ces caractères apparaissent dans la deuxième `string`.

Si cela est possible, le programme affiche la `string` suivie par un retour à la ligne (`'\n'`), autrement le programme affiche un retour à la ligne.

Si le nombre d'arguments est différent de 2, le programme affiche un retour à la ligne.

### Utilisation

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test "faya" "fgvvfdxcacpolhyghbreda"
faya
student@ubuntu:~/piscine-go/test$ ./test "faya" "fgvvfdxcacpolhyghbred"

student@ubuntu:~/piscine-go/test$ ./test "error" rrerrrfiiljdfxjyuifrrvcoojh

student@ubuntu:~/piscine-go/test$ ./test "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
quarante deux

student@ubuntu:~/piscine-go/test$ ./test

student@ubuntu:~/piscine-go/test$
```
