## splitprog

### Instructions

Écrire une fonction qui sépare les mots d'une `string`, qui les met dans un tableau de `string` et qui les affichent sur la sortie standard.

Le programme reçoit deux paramètres:

-   Le premier est la `string`
-   Le deuxième est le séparateur

### Utilisation :

```console
student@ubuntu:~/piscine-go/splitprog$ go build
student@ubuntu:~/piscine-go/splitprog$ ./splitprog "HelloHAhowHAareHAyou?" HA | cat -e
[Hello how are you?]$
student@ubuntu:~/piscine-go/splitprog$ ./splitprog "Hello,how,are,you?" ","
[Hello how are you?]
student@ubuntu:~/piscine-go/splitprog$ ./splitprog "HelloHAhowHAareHAyou?"
student@ubuntu:~/piscine-go/splitprog$
student@ubuntu:~/piscine-go/splitprog$ ./splitprog
student@ubuntu:~/piscine-go/splitprog$
```
