## cl-camp5

### Instructions

"continues à chercher..."

Créer un fichier `lookagain.sh`, qui cherchera et montrera, dans le répertoire courant et ses sous-répertoires, tous les fichiers qui:

- qui finissent avec `.sh`.

Cette commande montrera le nom des fichiers sans le`.sh`.

### Utilisation

```console
student@ubuntu:~/[[ROOT]]/test$ ./lookagain.sh | cat -e
file1$
file2$
file3$
student@ubuntu:~/[[ROOT]]/test$
```

### Indice

Un petit `cut`ter pourrait être utile...
