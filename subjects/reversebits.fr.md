## reversebits

### Instructions

Écrire une fonction qui prend un `byte`, qui l'inverse `bit` par `bit` (comme montré sur l'exemple) et qui retourne le résultat.

### Fonction attendue

```go
func ReverseBits(by byte) byte {

}
```
### Utilisation

Voici un programme possible pour tester votre fonction :

```go
package main

import (
	"fmt"
)

func main() {
	a := byte(0x26)
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", ReverseBits(a))
}
```

et son résultat :

```console
student@ubuntu:~/piscine-go/reversebits$ go build
student@ubuntu:~/piscine-go/reversebits$ ./reversebits
00100110
01100100
student@ubuntu:~/piscine-go/reversebits$
```
