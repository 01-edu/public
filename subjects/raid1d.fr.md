## raid1d

### Instructions

Écrire une fonction `Raid1d` qui affiche un rectangle **valide** de largeur `x` et de hauteur `y`.

Cette fonction doit dessiner les rectangles comme sur les examples.

### Fonction attendue

```go
func Raid1d(x,y int) {

}
```

### Utilisation

Voici d'éventuels programmes pour tester votre fonction :

Programme #1

```go
package main

import (
	student "./student"
)

func main() {
	student.Raid1d(5,3)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
ABBBC
B   B
ABBBC
student@ubuntu:~/[[ROOT]]/test$
```

Programme #2

```go
package main

import (
	student "./student"
)

func main() {
	student.Raid1d(5,1)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
ABBBC
student@ubuntu:~/[[ROOT]]/test$
```

Programme #3

```go
package main

import (
	student "./student"
)

func main() {
	student.Raid1d(1,1)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
A
student@ubuntu:~/[[ROOT]]/test$
```

Programme #4

```go
package main

import (
	student "./student"
)

func main() {
	student.Raid1d(1,5)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
A
B
B
B
A
student@ubuntu:~/[[ROOT]]/test$
```
