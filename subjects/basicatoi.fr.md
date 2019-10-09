## basicatoi

### Instructions

-   Écrire une fonction qui reproduit le comportement de la fonction atoi en Go. Atoi transforme un nombre représenté en `string` (chaîne de caractères) en `int` (entier).

-   Atoi retourne `0` si la `string` n'est pas considéré un nombre valide. Pour cet exercice **seulement des** `string` **valides** seront testé. Elles ne contiendront que un ou plusieurs chiffres comme charact.

-   Pour cet exercice la gestion des signes `+` ou `-` ne doit pas être prise en compte.

-   Cette fonction aura **seulement** à retourner l'`int` (entier). Pour cet exercice le retour d'erreur d'atoi de go n'est pas demandé.

### Fonction attendue

```go
func BasicAtoi(s string) int {

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "000000"

	n := piscine.BasicAtoi(s)
	n2 := piscine.BasicAtoi(s2)
	n3 := piscine.BasicAtoi(s3)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
12345
12345
0
student@ubuntu:~/piscine-go/test$
```
