## nruneprog

### Instructions

Écrire un programme qui reçoit une `string` en paramètre et qui retourne la énième `rune` de la `string`.

### Utilisation :

```console
student@ubuntu:~/[[ROOT]]/nruneprog$ ./nruneprog "hello" 2
e
student@ubuntu:~/[[ROOT]]/nruneprog$ ./nruneprog "hello" 4 | cat -e
l$
student@ubuntu:~/[[ROOT]]/nruneprog$ ./nruneprog "hello" 5
o
student@ubuntu:~/[[ROOT]]/nruneprog$ ./nruneprog "hello"
student@ubuntu:~/[[ROOT]]/nruneprog$ ./nruneprog "hello" f | cat -e
"f" is not an integer value$
student@ubuntu:~/[[ROOT]]/nruneprog$ ./nruneprog "hello" 9
Invalid position: "9" in "hello"
student@ubuntu:~/[[ROOT]]/nruneprog$
```
