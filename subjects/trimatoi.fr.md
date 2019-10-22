## trimatoi

### Instructions

-   'Écrire une [function](TODO-LINK) qui transforme un nombre dans une `string` en un nombre représenté en `int`(entier).

-   Pour cet exercice la gestion des signes + ou - **doit être** prise en compte. Si un des signes se situe avant un chiffre alors il déterminera le signe de l'`int` retourné.

-   Cette fonction aura **seulement** à retourner l'`int` (entier). En cas d'input invalide, la fonction devra retourné `0`.

### Fonction attendue

```go
func TrimAtoi(s string) int {

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
	s2 := "str123ing45"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "sd+x1fa2W3s4"
        s6 := "sd-x1fa2W3s4"
        s7 := "sdx1-fa2W3s4"

	n := piscine.TrimAtoi(s)
	n2 := piscine.TrimAtoi(s2)
	n3 := piscine.TrimAtoi(s3)
	n4 := piscine.TrimAtoi(s4)
	n5 := piscine.TrimAtoi(s5)
	n6 := piscine.TrimAtoi(s6)
	n7 := piscine.TrimAtoi(s7)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
12345
12345
12345
0
1234
-1234
1234
student@ubuntu:~/piscine-go/test$
```
