## enigma

### Instructions

Write a function called `Enigma` that receives pointers to as arguments and move its values around to hide them.

This function will put :

-   `a` into `c`.
-   `c` into `d`.
-   `d` into `b`.
-   `b` into `a`.

### Expected function

```go
func Enigma(a ***int, b *int, c *******int, d ****int) {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

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

	piscine.Enigma(a, b, c, d)

	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

}
```

And its output :

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
