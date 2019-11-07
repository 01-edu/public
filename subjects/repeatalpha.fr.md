## repeatalpha

### Instructions

Écrire un programme `repeat_alpha` qui prend une `string` et qui affiche la répétition de chaque caractère de l'alphabet autant de fois que l'index de cette lettre.

Le résultat doit être suivi d'un retour à la ligne (`'\n'`).

`'a'` devient `'a'`, `'b'` devient `'bb'`, `'e'` devient `'eeeee'`, etc...

Si une lettre est en majuscule, elle reste en majuscule, de même si elle est en minuscule.

Si le nombre d'arguments est différent de 1, le programme affiche un retour à la ligne (`'\n'`).

### Utilisation

```console
student@ubuntu:~/piscine-go/repeatalpha$ go build
student@ubuntu:~/piscine-go/repeatalpha$ ./repeatalpha "abc" | cat -e
abbccc
student@ubuntu:~/piscine-go/repeatalpha$ ./repeatalpha "Choumi." | cat -e
CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
student@ubuntu:~/piscine-go/repeatalpha$ ./repeatalpha "abacadaba 01!" | cat -e
abbacccaddddabba 01!$
student@ubuntu:~/piscine-go/repeatalpha$ ./repeatalpha | cat -e
$
student@ubuntu:~/piscine-go/repeatalpha$ ./repeatalpha "" | cat -e
$
student@ubuntu:~/piscine-go/repeatalpha$
```
