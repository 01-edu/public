## raid1b

### Instructions

Écrire une fonction `Raid1b` qui affiche un rectangle **valide** de largeur `x` et de hauteur `y`.

Cette fonction doit dessiner les rectangles comme sur les examples.

### Fonction attendue

```go
func Raid1b(x,y int) {

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
	student.Raid1b(5,3)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
/***\
*   *
\***/
student@ubuntu:~/[[ROOT]]/test$
```

Programme #2

```go
package main

import (
	student "./student"
)

func main() {
	student.Raid1b(5,1)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
/***\
student@ubuntu:~/[[ROOT]]/test$
```

Programme #3

```go
package main

import (
	student "./student"
)

func main() {
	student.Raid1b(1,1)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
/
student@ubuntu:~/[[ROOT]]/test$
```

Programme #4

```go
package main

import (
	student "./student"
)

func main() {
	student.Raid1b(1,5)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
/
*
*
*
\
student@ubuntu:~/[[ROOT]]/test$
```
