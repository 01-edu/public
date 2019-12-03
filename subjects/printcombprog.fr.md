## printcombprog

### Instructions

Écrire un programme qui affiche sur une seule ligne dans l'ordre croissant toutes les combinaisons possibles de trois chiffres différents tels que le premier est inférieur au second et le second est inférieur au troisième.

Les combinaisons sont séparées par une virgule et un espace.

### Utilisation

Voici une sortie **tronquée** :

```console
student@ubuntu:~/[[ROOT]]/printcombprog$ go build
student@ubuntu:~/[[ROOT]]/printcombprog$ ./printcombprog | cat -e
012, 013, 014, 015, 016, 017, 018, 019, 023, ..., 689, 789$
student@ubuntu:~/[[ROOT]]/printcombprog$
```

`000` et `999` ne sont pas des combinations valides parce que les chiffres ne sont pas différents.

`987` ne doit pas être affiché parce que le premier chiffre n'est pas inférieur au second.
