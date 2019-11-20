## raid1b

### Instructions

Write a function `Raid1b` that prints a **valid** rectangle of width `x` and of height `y`.

The function must draw the rectangles as in the examples.

`x` and `y` will always be positive numbers.

### Expected function

```go
func Raid1b(x,y int) {

}
```

### Usage

Here is are possible programs to test your function :

Program #1

```go
package main

import (
	piscine ".."
)

func main() {
	piscine.Raid1b(5,3)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
/***\
*   *
\***/
student@ubuntu:~/piscine-go/test$
```

Program #2

```go
package main

import (
	piscine ".."
)

func main() {
	piscine.Raid1b(5,1)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
/***\
student@ubuntu:~/piscine-go/test$
```

Program #3

```go
package main

import (
	piscine ".."
)

func main() {
	piscine.Raid1b(1,1)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
/
student@ubuntu:~/piscine-go/test$
```

Program #4

```go
package main

import (
	piscine ".."
)

func main() {
	piscine.Raid1b(1,5)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
/
*
*
*
\
student@ubuntu:~/piscine-go/test$
```
