## raid3

### Instructions

Ce raid est basé sur les fonctions du `raid1`.

Créer un programme `raid3` qui prend une `string` comme argument et qui affiche le nom du `raid1` correspondant et ses dimensions.

- Si l'argument n'est pas un des `raid1` le programme affiche `Not a Raid function`.

- Toutes les réponses doivent se terminer avec un retour à la ligne (`'\n'`).

- Si il y a plus d'un `raid1` correspondant, le programme doit les afficher tous en ordre alphabétique et séparé par un `||`.

### Utilisation

- Si c'est `radi1a`

```console
student@ubuntu:~/[[ROOT]]/raid3$ ls -l
-rw-r--r-- 1 student student  nov 23 14:30 raid3.go
-rwxr-xr-x 1 student student  nov 23 19:18 raid3
-rwxr-xr-x 1 student student  nov 23 19:50 raid1a
drwxr-xr-x 2 student student  nov 23 19:02 raid1aProg
-rwxr-xr-x 1 student student  nov 23 19:50 raid1b
drwxr-xr-x 2 student student  nov 22 23:36 raid1bProg
-rwxr-xr-x 1 student student  nov 23 19:50 raid1c
drwxr-xr-x 2 student student  nov 23 19:02 raid1cProg
-rwxr-xr-x 1 student student  nov 23 19:50 raid1d
drwxr-xr-x 2 student student  nov 23 00:40 raid1dProg
-rwxr-xr-x 1 student student  nov 23 19:50 raid1e
drwxr-xr-x 2 student student  nov 23 19:02 raid1eProg
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1a 3 3 | ./raid3
[raid1a] [3] [3]
student@ubuntu:~/[[ROOT]]/raid3$
student@ubuntu:~/[[ROOT]]/raid3$
student@ubuntu:~/[[ROOT]]/raid3$
student@ubuntu:~/[[ROOT]]/raid3$
```

- Si c'est `raidc 1 1` :

```console
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1c 1 1
A
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1d 1 1
A
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1e 1 1
A
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1c 1 1 | ./raid3
[raid1c] [2] [1] || [raid1d] [2] [1] || [raid1e] [2] [1]
student@ubuntu:~/[[ROOT]]/raid3$
```

- Si c'est `raidc 1 2` :

```console
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1e 1 2
A
C
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1c 1 2
A
C
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1e 1 2 | ./raid3
[raid1c] [1] [2] || [raid1e] [1] [2]
student@ubuntu:~/[[ROOT]]/raid3$
```
