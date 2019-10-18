## rpncalc

### Instructions

Écrire un programme qui prend une `string` qui contient une équation écrite en `Reverse Polish Notation` (RPN) comme premier argument, qui évalue l'équation, et qui affiche le résultat sur la sortie standard suivi d'un retour à la ligne (`'\n'`).

La `Reverse Polish Notation` est une notation mathématique dans laquelle chaque opérateur devance les valeurs qu'il va opérer.
En RPN, chaque opérateur évalue les deux précédentes valeurs, et le résultat de cette opération devient ensuite la première des valeurs de l'opérateur. Les valeurs et les opérateurs doivent être espacés d'au moins un espace.

Les opérateurs suivants doivent être implémentés : `+`, `-`, `*`, `/`, et `%`.

Si la `string` n'est pas valide ou si il n'y pas exactement un argument, le mot `Error` doit être affiché sur la sortie standard suivi d'un retour à la ligne.
Si la `string` a des espaces extra elle est toujours considérée valide.

Toutes les valeurs données doivent rentrer dans un `int`.

Exemples de formules converties en RPN:

3 + 4 >> 3 4 +
((1 _ 2) _ 3) - 4 >> 1 2 _ 3 _ 4 - ou 3 1 2 \* _ 4 -
50 _ (5 - (10 / 9)) >> 5 10 9 / - 50 \*

Voici comment évaluer une formule en RPN:

```
1 2 * 3 * 4 -
2 3 * 4 -
6 4 -
2
```

ou:

```
3 1 2 * * 4 -
3 2 * 4 -
6 4 -
2
```

### Utilisation

```console
student@ubuntu:~/piscine-go/rpncalc$ go build
student@ubuntu:~/piscine-go/rpncalc$ ./rpncalc "1 2 * 3 * 4 +" | cat -e
10$
student@ubuntu:~/piscine-go/rpncalc$ ./rpncalc 1 2 3 4 +" | cat -e
Error$
student@ubuntu:~/piscine-go/rpncalc$ ./rpncalc | cat -e
Error$
student@ubuntu:~/piscine-go/rpncalc$ ./rpncalc "     1      3 * 2 -" | cat -e
1
student@ubuntu:~/piscine-go/rpncalc$ ./rpncalc "     1      3 * ksd 2 -" | cat -e
Error$
student@ubuntu:~/piscine-go/rpncalc$
```
