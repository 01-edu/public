## raid3

### Instructions

Ce raid est basé sur les fonctions du `raid1`.

Créer un programme `raid3` qui prend une `string` comme argument et qui affiche le nom du `raid1` correspondant et ses dimensions.

-   Si l'argument n'est pas un des `raid1` le programme affiche `Not a Raid function`.

-   Toutes les réponses doivent se terminer avec un retour à la ligne (`'\n'`).

-   Si il y a plus d'un `raid1` correspondant, le programme doit les afficher tous en ordre alphabétique et séparé par un `|`.

### Utilisation

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ echo HELLO | ./raid3
Not a Raid function
student@ubuntu:~/piscine-go/test$ ./raid1a 4 4
o--o
|  |
|  |
o--o
student@ubuntu:~/piscine-go/test$ ./raid1a 4 4 | ./raid3
[raid1a] [4] [4]
student@ubuntu:~/piscine-go/test$ ./raid1a 3 4 | ./raid3
[raid1a] [3] [4]
student@ubuntu:~/piscine-go/test$ ./raid1c 1 1
A
student@ubuntu:~/piscine-go/test$ ./raid1d 1 1
A
student@ubuntu:~/piscine-go/test$ ./raid1e 1 1
A
student@ubuntu:~/piscine-go/test$ ./raid1c 1 1 | ./raid3
[raid1c] [1] [1] || [raid1d] [1] [1] || [raid1e] [1] [1]
student@ubuntu:~/piscine-go/test$
```
