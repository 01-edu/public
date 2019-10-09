## enigma

### Instructions

Écrire une fonction nommé `Enigma` qui prends des pointeurs comme arguments et qui interchanges leurs valeurs pour les cacher.
Cette fonction déplacera :

-   `a` dans `c`.
-   `c` dans `d`.
-   `d` dans `b`.
-   `b` dans `a`.

### Fonction attendue

```go
func Enigma(a ***int, b *int, c *******int, d ****int) {

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
	x := 5
	y := &x
	z := &y
	a := &z

	w := 2
	b := &w

	u := 7
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j

	k := 6
	l := &k
	m := &l
	n := &m
	d := &n

	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

	student.Enigma(a, b, c, d)

	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
5
2
7
6
After using Enigma
2
6
5
7
student@ubuntu:~/piscine-go/test$
```
