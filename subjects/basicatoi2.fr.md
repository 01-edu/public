## basicatoi2

### Instructions

-   Écrire une fonction qui reproduit le comportement de la fonction atoi en Go. Atoi transforme un nombre représenté en `string` (chaîne de caractères) en `int` (entier).

-   Atoi retourne `0` si la `string` n'est pas considérée comme un nombre valide. Pour cet exercice des **`string` non valides seront testées!**. Certaines contiendront d'autres caractères que des chiffres.

-   Pour cet exercice la gestion des signes `+` ou `-` ne doit pas être prise en compte.

-   Cette fonction aura **seulement** à retourner l'`int` (entier). Pour cet exercice le retour d'erreur d'atoi de go n'est pas demandé.

### Fonction attendue

```go
func BasicAtoi2(s string) int {

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
	s3 := "012 345"
	s4 := "Hello World!"

	n := piscine.BasicAtoi2(s)
	n2 := piscine.BasicAtoi2(s2)
	n3 := piscine.BasicAtoi2(s3)
	n4 := piscine.BasicAtoi2(s4)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
12345
12345
0
0
student@ubuntu:~/piscine-go/test$
```
