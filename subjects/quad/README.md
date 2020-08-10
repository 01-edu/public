# quadrangle

## quadA

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

Here are possible programs to test your function :

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

----

## quadB

### Instructions

Write a function `QuadB` that prints a **valid** rectangle of width `x` and of height `y`.

The function must draw the rectangles as in the examples.

`x` and `y` will always be positive numbers. Otherwise, the function should print nothing.

### Expected function

```go
func QuadB(x,y int) {

}
```

### Usage

Here are possible programs to test your function :

Program #1

```go
package main

import piscine ".."

func main() {
	piscine.QuadB(5,3)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
/***\
*   *
\***/
student@ubuntu:~/[[ROOT]]/test$
```

Program #2

```go
package main

import piscine ".."

func main() {
	piscine.QuadB(5,1)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
/***\
student@ubuntu:~/[[ROOT]]/test$
```

Program #3

```go
package main

import piscine ".."

func main() {
	piscine.QuadB(1,1)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
/
student@ubuntu:~/[[ROOT]]/test$
```

Program #4

```go
package main

import piscine ".."

func main() {
	piscine.QuadB(1,5)
}
```

And its output :

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

----

## quadC

### Instructions

Write a function `QuadC` that prints a **valid** rectangle of width `x` and of height `y`.

The function must draw the rectangles as in the examples.

`x` and `y` will always be positive numbers. Otherwise, the function should print nothing.

### Expected function

```go
func QuadC(x,y int) {

}
```

### Usage

Here are possible programs to test your function :

Program #1

```go
package main

import piscine ".."

func main() {
	piscine.QuadC(5,3)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
ABBBA
B   B
CBBBC
student@ubuntu:~/[[ROOT]]/test$
```

Program #2

```go
package main

import piscine ".."

func main() {
	piscine.QuadC(5,1)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
ABBBA
student@ubuntu:~/[[ROOT]]/test$
```

Program #3

```go
package main

import piscine ".."

func main() {
	piscine.QuadC(1,1)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
A
student@ubuntu:~/[[ROOT]]/test$
```

Program #4

```go
package main

import piscine ".."

func main() {
	piscine.QuadC(1,5)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
A
B
B
B
C
student@ubuntu:~/[[ROOT]]/test$
```

----

## quadD

### Instructions

Write a function `QuadD` that prints a **valid** rectangle of width `x` and of height `y`.

The function must draw the rectangles as in the examples.

`x` and `y` will always be positive numbers. Otherwise, the function should print nothing.

### Expected function

```go
func QuadD(x,y int) {

}
```

### Usage

Here are possible programs to test your function :

Program #1

```go
package main

import piscine ".."

func main() {
	piscine.QuadD(5,3)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
ABBBC
B   B
ABBBC
student@ubuntu:~/[[ROOT]]/test$
```

Program #2

```go
package main

import piscine ".."

func main() {
	piscine.QuadD(5,1)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
ABBBC
student@ubuntu:~/[[ROOT]]/test$
```

Program #3

```go
package main

import piscine ".."

func main() {
	piscine.QuadD(1,1)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
A
student@ubuntu:~/[[ROOT]]/test$
```

Program #4

```go
package main

import piscine ".."

func main() {
	piscine.QuadD(1,5)
}
```

And its output :

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

----

## quadE

### Instructions

Write a function `QuadE` that prints a **valid** rectangle of width `x` and of height `y`.

The function must draw the rectangles as in the examples.

`x` and `y` will always be positive numbers. Otherwise, the function should print nothing.

### Expected function

```go
func QuadE(x,y int) {

}
```

### Usage

Here are possible programs to test your function :

Program #1

```go
package main

import piscine ".."

func main() {
	piscine.QuadE(5,3)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
ABBBC
B   B
CBBBA
student@ubuntu:~/[[ROOT]]/test$
```

Program #2

```go
package main

import piscine ".."

func main() {
	piscine.QuadE(5,1)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
ABBBC
student@ubuntu:~/[[ROOT]]/test$
```

Program #3

```go
package main

import piscine ".."

func main() {
	piscine.QuadE(1,1)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
A
student@ubuntu:~/[[ROOT]]/test$
```

Program #4

```go
package main

import piscine ".."

func main() {
	piscine.QuadE(1,5)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
A
B
B
B
C
student@ubuntu:~/[[ROOT]]/test$
```
