## reverserange

### Instructions

Écrire la fonction `ReverseRange` qui doit:

- allouer (avec make()) une slice d'entiers.
- le remplir avec des valeurs consécutives qui commencent à `end` et qui finissent à `start` (En incluant `start` et `end` !)
- et qui retourne cette slice.

### Fonction attendue

```go
func ReverseRange(start, end int) []int {

}
```

### Utilisation :

- Avec (1, 3) la fonction devra retourner une slice contenant 3, 2 et 1.
- Avec (-1, 2) la fonction devra retourner une slice contenant 2, 1, 0 et -1.
- Avec (0, 0) la fonction devra retourner une slice contenant 0.
- Avec (0, -3) la fonction devra retourner une slice contenant -3, -2, -1 et 0.
