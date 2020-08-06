## quadrangle

### Instructions

Write a function `QuadA` that prints a **valid** rectangle of width `x` and of height `y`.

The function must draw the rectangles as in the examples.

`x` and `y` will always be positive numbers. Otherwise, the function should print nothing.

### Expected function

```go
func QuadA(x,y int) {

}
```

### Usage

Here is are possible programs to test your function :

Program #1

```go
package main

import piscine ".."

func main() {
	piscine.QuadA(5,3)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
o---o
|   |
o---o
student@ubuntu:~/[[ROOT]]/test$
```

Program #2

```go
package main

import piscine ".."

func main() {
	piscine.QuadA(5,1)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
o---o
student@ubuntu:~/[[ROOT]]/test$
```

Program #3

```go
package main

import piscine ".."

func main() {
	piscine.QuadA(1,1)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
o
student@ubuntu:~/[[ROOT]]/test$
```

Program #4

```go
package main

import piscine ".."

func main() {
	piscine.QuadA(1,5)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
o
|
|
|
o
student@ubuntu:~/[[ROOT]]/test$
```
