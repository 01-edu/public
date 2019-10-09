## raid1a

### Instructions

Write a function `Raid1a` that prints a **valid** square of width `x` and of height `y`.

The function must draw the squares as in the examples.

### Expected function

```go
func Raid1a(x,y int) {

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
	student.Raid1a(5,3)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
o---o
|   |
o---o
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
	student.Raid1a(5,1)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
o---o
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
	student.Raid1a(1,1)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
o
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
	student.Raid1a(1,5)
}
```

And its output :

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
