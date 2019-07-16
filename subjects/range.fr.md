## range

### Instructions

Écrire la fonction `Range` qui doit:

- allouer (avec make()) une slice d'entiers.
- le remplir avec des valeurs consécutives qui commencent à `start` et qui finissent à `end` (En incluant `start` et `end` !)
- et qui retourne cette slice.

### Fonction attendue

```go
func Range(start, end int) []int {

}
```

### Utilisation

- Avec (1, 3) la fonction devra retourner une slice contenant 1, 2 et 3.
- Avec (-1, 2) la fonction devra retourner une slice contenant -1, 0, 1 et 2.
- Avec (0, 0) la fonction devra retourner une slice contenant 0.
- Avec (0, -3) la fonction devra retourner une slice contenant 0, -1, -2 et -3.
