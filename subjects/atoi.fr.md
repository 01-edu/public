## atoi

### Instructions

-   Écrire une fonction qui reproduit le comportement de la fonction `Atoi` en Go. `Atoi` transforme un nombre représenté en `string` (chaîne de caractères) en `int` (entier).

-   `Atoi` retourne `0` si la `string` n'est pas considéré un nombre valide. Pour cet exercice des **`string` non valides seront testées!**. Certaines contiendront d'autres caractères que des chiffres.

-   Pour cet exercice la gestion des signes + ou - **doit être** prise en compte.

-   Cette fonction aura **seulement** à retourner l'`int` (entier). Pour cet exercice le retour d'erreur d'atoi de go n'est pas demandé.

### Fonction attendue

```go
func Atoi(s string) int {

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
	s5 := "+1234"
	s6 := "-1234"
	s7 := "++1234"
	s8 := "--1234"

	n := piscine.Atoi(s)
	n2 := piscine.Atoi(s2)
	n3 := piscine.Atoi(s3)
	n4 := piscine.Atoi(s4)
	n5 := piscine.Atoi(s5)
	n6 := piscine.Atoi(s6)
	n7 := piscine.Atoi(s7)
	n8 := piscine.Atoi(s8)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
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
1234
-1234
0
0
student@ubuntu:~/piscine-go/test$
```
