## printrevcombprog

### Instructions

Écrire un programme qui affiche sur une seule ligne dans l'ordre décroissant toutes les combinaisons possibles de trois chiffres différents tels que le premier est supérieur au second et le second est supérieur au troisième.

Les combinaisons sont séparées par une virgule et un espace.

### Utilisation

Voici une sortie **tronquée** :

```console
student@ubuntu:~/piscine-go/printcombprog$ go build
student@ubuntu:~/piscine-go/printcombprog$ ./printcombprog | cat -e
987, 986, 985, 984, 983, 982, 981, 980, 976, ..., 310, 210$
student@ubuntu:~/piscine-go/printcombprog$
```

`999` et `000` ne sont pas des combinations valides parce que les chiffres ne sont pas différents.

`789` ne doit pas être affiché parce que le premier chiffre n'est pas inférieur au second.
