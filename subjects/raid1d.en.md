## raid1d

### Instructions

Write a function `Raid1d` that prints a **valid** square of width `x` and of height `y`.

The function must draw the squares as in the examples.

### Expected function

```go
func Raid1d(x,y int) {

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
	student.Raid1d(5,3)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
ABBBC
B   B
ABBBC
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
	student.Raid1d(5,1)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
ABBBC
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
	student.Raid1d(1,1)
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
	student.Raid1d(1,5)
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
A
student@ubuntu:~/piscine-go/test$
```
