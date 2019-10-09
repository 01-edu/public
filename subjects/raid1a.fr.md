## raid1a

### Instructions

Écrire une fonction `Raid1a` qui affiche un carré **valide** de largeur `x` et de hauteur `y`.

Cette fonction doit dessiner les carrés comme sur les examples.

### Fonction attendue

```go
func Raid1a(x,y int) {

}
```

### Utilisation

Voici d'éventuels programmes pour tester votre fonction :

Programme #1

```go
package main

import (
	"fmt"
	student "./student"
)

func main() {
	student.Raid1a(5,3)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
o---o
|   |
o---o
student@ubuntu:~/piscine-go/test$
```

Programme #2

```go
package main

import (
	"fmt"
	student "./student"
)

func main() {
	student.Raid1a(5,1)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
o---o
student@ubuntu:~/piscine-go/test$
```

Programme #3

```go
package main

import (
	"fmt"
	student "./student"
)

func main() {
	student.Raid1a(1,1)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
o
student@ubuntu:~/piscine-go/test$
```

Programme #4

```go
package main

import (
	"fmt"
	student "./student"
)

func main() {
	student.Raid1a(1,5)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
o
|
|
|
o
student@ubuntu:~/piscine-go/test$
```
