## Display File

### Instructions

Écrire un programme qui affiche, sur la sortie standard, le contenu d'un fichier donné en argument.

-   Créer un fichier `quest8.txt` et écrire dedans la phrase `Almost there!!`

-   L'argument pour ce programme sera, dans ce cas, `quest8.txt`.

-   En cas d'erreur le programme doit afficher un des deux messages suivants de manière approprié:
    -   `File name missing`.
    -   `Too many arguments`.

### Utilisation:

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
File name missing
student@ubuntu:~/piscine-go/test$ ./test quest8.txt main.go
Too many arguments
student@ubuntu:~/piscine-go/test$ ./test quest8.txt
Almost there!!
```
