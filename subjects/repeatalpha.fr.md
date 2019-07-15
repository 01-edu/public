## repeatalpha

### Instructions

Écrire un programme `repeat_alpha` qui prend une `string` et qui affiche la répétition de chaque caractère de l'alphabet autant de fois que l'index de cette lettre.

Le résultat doit être suivi d'un newline(`'\n'`).

`'a'` devient `'a'`, `'b'` devient `'bb'`, `'e'` devient `'eeeee'`, etc...

Si une lettre est en majuscule, elle reste en majuscule, de même si elle est en minuscule.

Si le nombre d'arguments est différent de 1, le programme affiche un newline(`'\n'`).

### Utilisation

```console
student@ubuntu:~/student/repeatalpha$ go build
student@ubuntu:~/student/repeatalpha$ ./repeatalpha "abc" | cat -e
abbccc
student@ubuntu:~/student/repeatalpha$ ./repeatalpha "Alex." | cat -e
Alllllllllllleeeeexxxxxxxxxxxxxxxxxxxxxxxx.$
student@ubuntu:~/student/repeatalpha$ ./repeatalpha "abacadaba 42!" | cat -e
abbacccaddddabba 42!$
student@ubuntu:~/student/repeatalpha$ ./repeatalpha | cat -e
$
student@ubuntu:~/student/repeatalpha$ ./repeatalpha "" | cat -e
$
student@ubuntu:~/student/repeatalpha$
```
