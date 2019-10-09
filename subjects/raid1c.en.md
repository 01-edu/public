## raid1c

### Instructions

Write a function `Raid1c` that prints a **valid** square of width `x` and of height `y`.

The function must draw the squares as in the examples.

### Expected function

```go
func Raid1c(x,y int) {

}
```

### Usage

Here is are possible programs to test your function :

Program #1

```go
package main

import (
	"fmt"
	student "./student"
)

func main() {
	student.Raid1c(5,3)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
ABBBA
B   B
CBBBC
student@ubuntu:~/piscine-go/test$
```

Program #2

```go
package main

import (
	"fmt"
	student "./student"
)

func main() {
	student.Raid1c(5,1)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
ABBBA
student@ubuntu:~/piscine-go/test$
```

Program #3

```go
package main

import (
	"fmt"
	student "./student"
)

func main() {
	student.Raid1c(1,1)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
A
student@ubuntu:~/piscine-go/test$
```

Program #4

```go
package main

import (
	"fmt"
	student "./student"
)

func main() {
	student.Raid1c(1,5)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
A
B
B
B
C
student@ubuntu:~/piscine-go/test$
```
